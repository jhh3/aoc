package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

func main() {
	common.RunFromSolver(&solver{}, 2023, 4)
}

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
	return ""
}

// Parsing code

var cardRegex = regexp.MustCompile(`Card\s+(\d+):`)

type Card struct {
	id             string
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
		id:             id,
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
