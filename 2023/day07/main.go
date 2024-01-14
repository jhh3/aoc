package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	camelCardInput := parseCamelCardInput(input)

	// Sort hands
	sort.Slice(camelCardInput.Hands, func(i, j int) bool {
		return camelCardInput.Hands[i].Compare(camelCardInput.Hands[j]) == -1
	})

	result := 0.0
	for rank, hand := range camelCardInput.Hands {
		result += hand.Bid * float64(rank+1)
	}

	return fmt.Sprintf("%.0f", result)
}

func (s *solver) SolvePart2(input string) string {
	panic("not implemented")
}

//--------------------------------------------------------------------
// Parse Input
//--------------------------------------------------------------------

type HandType int

const (
	FIVE_OF_A_KIND HandType = iota
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIRS
	ONE_PAIR
	HIGH_CARD
)

type Hand struct {
	Cards      string
	CardValues []int
	Bid        float64
	Type       HandType
}

func (h Hand) Compare(other Hand) int {
	if h.Type == other.Type {
		// If types equal then compare card values
		for i := 0; i < len(h.CardValues); i++ {
			if h.CardValues[i] == other.CardValues[i] {
				continue
			}

			if h.CardValues[i] > other.CardValues[i] {
				return 1
			} else {
				return -1
			}
		}

		return 0
	} else {
		if h.Type < other.Type {
			return 1
		} else {
			return -1
		}
	}
}

// func (h Hand) Compare(other Hand) int {
// }

type CamelCardInput struct {
	Hands []Hand
}

func parseCamelCardInput(input string) CamelCardInput {
	result := CamelCardInput{}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		cleanLine := strings.TrimRight(line, "\n")

		result.Hands = append(result.Hands, parseHand(cleanLine))
	}

	return result
}

func parseHand(input string) Hand {
	pieces := strings.Split(input, " ")
	cardStr := strings.TrimSpace(pieces[0])
	cardValues := []int{}
	for _, card := range cardStr {
		switch card {
		case 'A':
			cardValues = append(cardValues, 14)
		case 'K':
			cardValues = append(cardValues, 13)
		case 'Q':
			cardValues = append(cardValues, 12)
		case 'J':
			cardValues = append(cardValues, 11)
		case 'T':
			cardValues = append(cardValues, 10)
		default:
			cardValues = append(cardValues, common.MustAtoi(string(card)))
		}
	}

	// Classify hand
	valCountMap := map[int]int{}
	for _, val := range cardValues {
		valCountMap[val]++
	}
	countMap := map[int]int{}
	for _, count := range valCountMap {
		countMap[count]++
	}

	handType := HIGH_CARD
	if countMap[5] == 1 {
		handType = FIVE_OF_A_KIND
	}
	if countMap[4] == 1 {
		handType = FOUR_OF_A_KIND
	}
	if countMap[3] == 1 {
		if countMap[2] == 1 {
			handType = FULL_HOUSE
		} else {
			handType = THREE_OF_A_KIND
		}
	} else {
		if countMap[2] == 2 {
			handType = TWO_PAIRS
		} else if countMap[2] == 1 {
			handType = ONE_PAIR
		}
	}

	return Hand{
		Cards:      cardStr,
		CardValues: cardValues,
		Type:       handType,
		Bid:        float64(common.MustAtoi(strings.TrimSpace(pieces[1]))),
	}
}
