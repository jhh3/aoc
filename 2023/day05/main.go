package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"slices"
	"strings"
	"sync"

	"github.com/jhh3/aoc/common"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	// use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	common.RunFromSolver(&solver{}, input)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	seedInput := ParseSeedInput(input)

	var locations []int
	for _, seed := range seedInput.Seeds {
		locations = append(locations, seedInput.GetLocation(seed))
	}

	minLocation := slices.Min(locations)

	return fmt.Sprintf("%d", minLocation)
}

func (s *solver) SolvePart2(input string) string {
	seedInput := ParseSeedInput(input)
	var wg sync.WaitGroup

	processSeed := func(seed int, rngLen int, ch chan<- int) {
		defer wg.Done()
		var locations []int
		for i := 0; i < rngLen; i++ {
			locations = append(locations, seedInput.GetLocation(seed+i))
		}
		minLocation := slices.Min(locations)
		ch <- minLocation
	}

	numSeeds := len(seedInput.Seeds) / 2
	wg.Add(numSeeds)
	valueChan := make(chan int, numSeeds)

	for idx, seed := range seedInput.Seeds {
		if idx%2 == 0 {
			go processSeed(seed, seedInput.Seeds[idx+1], valueChan)
		}
	}

	wg.Wait()
	close(valueChan)

	locations := []int{}
	for result := range valueChan {
		locations = append(locations, result)
	}

	minLocation := slices.Min(locations)

	return fmt.Sprintf("%d", minLocation)
}

//--------------------------------------------------------------------
// Parser
//--------------------------------------------------------------------

type SeedMap struct {
	DestRng int
	SrcRng  int
	RngLen  int
}

type SeedInput struct {
	Seeds []int

	SeedToSoil            []SeedMap
	SoilToFertilizer      []SeedMap
	FetilizerToWater      []SeedMap
	WaterToLight          []SeedMap
	LightToTemperature    []SeedMap
	TemperatureToHumidity []SeedMap
	HumidityToLocation    []SeedMap
}

func GetNextFromMap(seed int, seedMap []SeedMap) int {
	for _, seedMap := range seedMap {
		if seedMap.SrcRng <= seed && seed < seedMap.SrcRng+seedMap.RngLen {
			return seedMap.DestRng + (seed - seedMap.SrcRng)
		}
	}
	return seed
}

func (s *SeedInput) GetSoil(seed int) int {
	return GetNextFromMap(seed, s.SeedToSoil)
}

func (s *SeedInput) GetFertilizer(seedId int) int {
	soil := s.GetSoil(seedId)
	return GetNextFromMap(soil, s.SoilToFertilizer)
}

func (s *SeedInput) GetWater(seedId int) int {
	fertilizer := s.GetFertilizer(seedId)
	return GetNextFromMap(fertilizer, s.FetilizerToWater)
}

func (s *SeedInput) GetLight(seedId int) int {
	water := s.GetWater(seedId)
	return GetNextFromMap(water, s.WaterToLight)
}

func (s *SeedInput) GetTemperature(seedId int) int {
	light := s.GetLight(seedId)
	return GetNextFromMap(light, s.LightToTemperature)
}

func (s *SeedInput) GetHumidity(seedId int) int {
	temperature := s.GetTemperature(seedId)
	return GetNextFromMap(temperature, s.TemperatureToHumidity)
}

func (s *SeedInput) GetLocation(seedId int) int {
	humidity := s.GetHumidity(seedId)
	return GetNextFromMap(humidity, s.HumidityToLocation)
}

func ParseSeedInput(input string) *SeedInput {
	result := SeedInput{}
	lines := strings.Split(string(input), "\n")

	mode := ""
	justActivatedMap := false
	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)

		if cleanLine == "" {
			continue
		}

		mode, justActivatedMap = getMode(cleanLine, mode)
		if justActivatedMap {
			continue
		}

		lineMap := SeedMap{}
		if mode != "seeds" {
			lineMap = lineToMap(cleanLine)
		}

		switch mode {
		case "seeds":
			seedIdsStr := strings.TrimSpace(strings.Split(cleanLine, ":")[1])
			for _, seedId := range strings.Split(seedIdsStr, " ") {
				result.Seeds = append(result.Seeds, common.MustAtoi(seedId))
			}
		case "seed-to-soil":
			result.SeedToSoil = append(result.SeedToSoil, lineMap)
		case "soil-to-fertilizer":
			result.SoilToFertilizer = append(result.SoilToFertilizer, lineMap)
		case "fertilizer-to-water":
			result.FetilizerToWater = append(result.FetilizerToWater, lineMap)
		case "water-to-light":
			result.WaterToLight = append(result.WaterToLight, lineMap)
		case "light-to-temperature":
			result.LightToTemperature = append(result.LightToTemperature, lineMap)
		case "temperature-to-humidity":
			result.TemperatureToHumidity = append(result.TemperatureToHumidity, lineMap)
		case "humidity-to-location":
			result.HumidityToLocation = append(result.HumidityToLocation, lineMap)
		default:
			panic("unknown mode: " + mode)

		}
	}

	return &result
}

func lineToMap(line string) SeedMap {
	linePieces := strings.Split(line, " ")
	destRng := common.MustAtoi(linePieces[0])
	srcRng := common.MustAtoi(linePieces[1])
	rnglen := common.MustAtoi(linePieces[2])

	return SeedMap{
		DestRng: destRng,
		SrcRng:  srcRng,
		RngLen:  rnglen,
	}
}

var lineToModes = map[string]string{
	"seed-to-soil map:":            "seed-to-soil",
	"soil-to-fertilizer map:":      "soil-to-fertilizer",
	"fertilizer-to-water map:":     "fertilizer-to-water",
	"water-to-light map:":          "water-to-light",
	"light-to-temperature map:":    "light-to-temperature",
	"temperature-to-humidity map:": "temperature-to-humidity",
	"humidity-to-location map:":    "humidity-to-location",
}

func getMode(line string, currentMode string) (string, bool) {
	if strings.HasPrefix(line, "seeds:") {
		return "seeds", false
	}

	if maybeMode, ok := lineToModes[line]; ok {
		return maybeMode, true
	}

	return currentMode, false
}
