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
	result := problemInput.SumMiddleOfValidPageOrders()
	return strconv.Itoa(result)
}

func (s *solver) SolvePart2(input string) string {
	// TODO
	return ""
}

type PageOrder struct {
	pages []int
}

func (po *PageOrder) IsValid(rules map[int][]int) bool {
	seen := make(map[int]bool)
	ifSeenInvalid := make(map[int]bool)

	for _, page := range po.pages {
		// ensure we shouldn't have seen this page already
		if _, ok := ifSeenInvalid[page]; ok {
			return false
		}

		// mark this page as seen
		seen[page] = true

		// mark pages we can't see in the future
		if pagesThatIfAlsoProducedMustBeProducedFirst, ok := rules[page]; ok {
			for _, mustBeProducedFirst := range pagesThatIfAlsoProducedMustBeProducedFirst {
				if _, ok := seen[mustBeProducedFirst]; !ok {
					// if we've not seen this page yet, we should never see it
					// if we see it, this page order is invalid
					ifSeenInvalid[mustBeProducedFirst] = true
				}
			}
		}
		// no rules for this page, continue
	}

	return true
}

type ProblemInput struct {
	// X|Y in input
	// => if both X and Y are to be produced, X must be produced before Y
	// store Y -> list of X
	rules map[int][]int

	pageOrders []PageOrder
}

func (pi *ProblemInput) SumMiddleOfValidPageOrders() int {
	result := 0

	validPageOrders := pi.ValidPageOrders()
	for _, pageOrder := range validPageOrders {
		middlePage := pageOrder.pages[len(pageOrder.pages)/2]
		result += middlePage
	}

	return result
}

func (pi *ProblemInput) ValidPageOrders() []PageOrder {
	validPageOrders := make([]PageOrder, 0)
	for _, pageOrder := range pi.pageOrders {
		if pageOrder.IsValid(pi.rules) {
			validPageOrders = append(validPageOrders, pageOrder)
		}
	}

	return validPageOrders
}

func parseInput(input string) *ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	problemInput := ProblemInput{
		rules:      make(map[int][]int),
		pageOrders: make([]PageOrder, 0),
	}

	for _, line := range lines {
		// handle empty line between rules and page orders
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}

		// rules
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				panic("Invalid rule")
			}
			left, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err != nil {
				panic(err)
			}

			// add to rules, creating if necessary
			if _, ok := problemInput.rules[right]; !ok {
				problemInput.rules[right] = make([]int, 0)
			}
			problemInput.rules[right] = append(problemInput.rules[right], left)
		} else {
			// page order
			pages := strings.Split(line, ",")
			pageOrder := PageOrder{
				pages: make([]int, len(pages)),
			}
			for i, page := range pages {
				pageInt, err := strconv.Atoi(page)
				if err != nil {
					panic(err)
				}
				pageOrder.pages[i] = pageInt
			}
			problemInput.pageOrders = append(problemInput.pageOrders, pageOrder)

		}
	}

	return &problemInput
}
