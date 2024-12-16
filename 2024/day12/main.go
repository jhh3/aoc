package main

import (
	_ "embed"
	"strconv"
	"strings"

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
}

func main() {
	common.RunFromSolver(&solver{}, input)
}

//--------------------------------------------------------------------
// Solution
//--------------------------------------------------------------------

type solver struct{}

func (s *solver) SolvePart1(input string) string {
	problemInput := parseInput(input)
	cost := problemInput.ComputeFenceCost()
	return strconv.Itoa(cost)
}

func (s *solver) SolvePart2(input string) string {
	// TODO: Implement part 2
	return ""
}

type Point struct {
	Row, Col int
}

type ProblemInput struct {
	Garden             [][]rune
	Visited            map[Point]bool
	PlantProcessedMap  map[rune]bool
	PlantToPositionMap map[rune][]Point
}

// Regions are defined as a contiguous group of plants of the same type
type Region struct {
	PlantType rune
	Points    []Point
}

func (pi *ProblemInput) ComputeFenceCost() int {
	cost := 0

	for plantType := range pi.PlantProcessedMap {
		regions := pi.Regions(plantType)
		for _, region := range regions {
			cost += pi.ComputeCostForRegion(region)
		}
	}

	return cost
}

func (pi *ProblemInput) ComputeCostForRegion(region Region) int {
	// Cost defined as area * perimeter
	area := len(region.Points)
	perimeter := 0

	for _, point := range region.Points {
		// each point can contribute up to 4 to the perimeter
		contribution := 0
		for _, neighbor := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextPoint := Point{Row: point.Row + neighbor.Row, Col: point.Col + neighbor.Col}
			// if the neighbor is out of bounds or not the same plant type, it contributes to the perimeter
			if !pi.IsInBounds(nextPoint) || pi.Garden[nextPoint.Row][nextPoint.Col] != region.PlantType {
				contribution++
			}
		}
		perimeter += contribution
	}

	return perimeter * area
}

func (pi *ProblemInput) Regions(plantType rune) []Region {
	regions := make([]Region, 0)

	points := pi.PlantToPositionMap[plantType]
	seen := make(map[Point]bool)
	for _, point := range points {
		// Have we seen this point before?
		if _, ok := seen[point]; ok {
			continue
		}

		// Start a new region
		region := Region{PlantType: plantType, Points: make([]Point, 0)}
		pointsToProcess := []Point{point}
		pointsProcessed := make(map[Point]bool)
		for len(pointsToProcess) > 0 {
			// pop the first point
			point := pointsToProcess[0]
			pointsToProcess = pointsToProcess[1:]

			// Have we seen this point before?
			if _, ok := pointsProcessed[point]; ok {
				continue
			}

			// Add this point to the region
			region.Points = append(region.Points, point)
			pointsProcessed[point] = true

			// Add the neighbors to the points to process
			for _, neighbor := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextPoint := Point{Row: point.Row + neighbor.Row, Col: point.Col + neighbor.Col}
				if pi.IsInBounds(nextPoint) && pi.Garden[nextPoint.Row][nextPoint.Col] == plantType {
					if _, ok := pointsProcessed[nextPoint]; !ok {
						pointsToProcess = append(pointsToProcess, nextPoint)
					}
				}
			}
		}

		// Add the region to the list of regions
		regions = append(regions, region)

		// Mark all the points in the region as seen
		for _, point := range region.Points {
			seen[point] = true
		}
	}

	return regions
}

func (pi *ProblemInput) IsInBounds(point Point) bool {
	return point.Row >= 0 && point.Row < len(pi.Garden) && point.Col >= 0 && point.Col < len(pi.Garden[0])
}

func parseInput(input string) *ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	pi := &ProblemInput{
		Garden:             make([][]rune, len(lines)),
		Visited:            make(map[Point]bool),
		PlantProcessedMap:  make(map[rune]bool),
		PlantToPositionMap: make(map[rune][]Point),
	}

	for i, row := range lines {
		pi.Garden[i] = []rune(row)
		for j, element := range row {
			pi.PlantProcessedMap[element] = false
			point := Point{Row: i, Col: j}
			if _, ok := pi.PlantToPositionMap[element]; !ok {
				pi.PlantToPositionMap[element] = make([]Point, 0)
			}
			pi.PlantToPositionMap[element] = append(pi.PlantToPositionMap[element], point)
		}
	}

	return pi
}
