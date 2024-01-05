package main

import (
	_ "embed"
	"math"
	"regexp"
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

func (ps *solver) SolvePart1(input string) string {
	cards := parseInput(input)
	sum := 0

	for _, card := range cards {
		intersection := common.SetIntersection[int](card.numsYouHave, card.winningNumbers)
		if len(intersection) == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(len(intersection)-1)))
	}

	return strconv.Itoa(sum)
}

func (ps *solver) SolvePart2(input string) string {
	cards := parseInput(input)
	numCardssMap := map[int]int{}

	// init map
	for _, card := range cards {
		numCardssMap[card.id] = 1
	}

	// gain card copies
	for _, card := range cards {
		intersection := common.SetIntersection[int](card.numsYouHave, card.winningNumbers)
		if len(intersection) == 0 {
			continue
		}
		mul := numCardssMap[card.id]
		for id := 0; id < len(intersection); id++ {
			numCardssMap[card.id+id+1] += mul
		}
	}

	// count total cards
	totalCards := 0
	for _, quantity := range numCardssMap {
		totalCards += quantity
	}

	return strconv.Itoa(totalCards)
}

// Parsing code

var cardRegex = regexp.MustCompile(`Card\s+(\d+):`)

type Card struct {
	id             int
	numsYouHave    []int
	winningNumbers []int
}

func parseInput(input string) []Card {
	cards := []Card{}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}
		cards = append(cards, parseCard(line))
	}
	return cards
}

func parseCard(line string) Card {
	id := cardRegex.FindStringSubmatch(line)[1]

	numsPiece := strings.Split(line, ":")[1]
	pieces := strings.Split(numsPiece, "|")
	numsYouHaveStr := strings.TrimSpace(pieces[0])
	winningNumbersStr := strings.TrimSpace(pieces[1])

	return Card{
		id:             common.MustAtoi(id),
		numsYouHave:    MustStringToInts(numsYouHaveStr, " "),
		winningNumbers: MustStringToInts(winningNumbersStr, " "),
	}
}

// MustStringToInts converts a string of integers into a slice of ints.
func MustStringToInts(str string, sep string) []int {
	ints := []int{}
	for _, num := range strings.Split(str, sep) {
		cleanNum := strings.TrimSpace(num)
		if cleanNum == "" {
			continue
		}
		ints = append(ints, common.MustAtoi(cleanNum))
	}
	return ints
}
