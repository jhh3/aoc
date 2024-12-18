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
	result := 0
	for _, pageOrder := range problemInput.ValidPageOrders() {
		middlePage := pageOrder.pages[len(pageOrder.pages)/2]
		result += middlePage
	}
	return strconv.Itoa(result)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	result := 0
	for _, invalidPageOrder := range problemInput.InvalidPageOrders() {
		correctedPageOrder := invalidPageOrder.Correct(problemInput.rules)
		result += correctedPageOrder.pages[len(correctedPageOrder.pages)/2]
	}
	return strconv.Itoa(result)
}

type PageOrder struct {
	pages []int
}

func (po *PageOrder) Correct(rules map[int][]int) *PageOrder {
	// If valid, return
	validity := po.IsValid(rules)
	if validity.IsValid {
		return po
	}

	// If invalid, make correction and continue
	badPage := po.pages[validity.InvalidPageIndex]
	pagesCopy := make([]int, len(po.pages))
	copy(pagesCopy, po.pages)

	// Remove bad page
	newPages := common.RemoveIndex(pagesCopy, validity.InvalidPageIndex)
	// Insert page before the page it should come before
	newPages = common.Insert(newPages, validity.ShouldComeBeforePageIndex, badPage)
	newPageOrder := &PageOrder{pages: newPages}

	return newPageOrder.Correct(rules)
}

type Validity struct {
	IsValid                   bool
	InvalidPageIndex          int // -1 if valid
	ShouldComeBeforePageIndex int // -1 if valid
}

func (po *PageOrder) IsValid(rules map[int][]int) Validity {
	seen := make(map[int]bool)
	ifSeenInvalid := make(map[int]int)

	for i, page := range po.pages {
		// ensure we shouldn't have seen this page already
		if shouldComeBeforePageIndex, ok := ifSeenInvalid[page]; ok {
			return Validity{IsValid: false, InvalidPageIndex: i, ShouldComeBeforePageIndex: shouldComeBeforePageIndex}
		}

		// mark this page as seen
		seen[page] = true

		// mark pages we can't see in the future
		if pagesThatIfAlsoProducedMustBeProducedFirst, ok := rules[page]; ok {
			for _, mustBeProducedFirst := range pagesThatIfAlsoProducedMustBeProducedFirst {
				if _, ok := seen[mustBeProducedFirst]; !ok {
					// if we've not seen this page yet, we should never see it
					// if we see it, this page order is invalid
					ifSeenInvalid[mustBeProducedFirst] = i
				}
			}
		}
		// no rules for this page, continue
	}

	return Validity{IsValid: true, InvalidPageIndex: -1, ShouldComeBeforePageIndex: -1}
}

type ProblemInput struct {
	// X|Y in input
	// => if both X and Y are to be produced, X must be produced before Y
	// store Y -> list of X
	rules map[int][]int

	pageOrders []PageOrder
}

func (pi *ProblemInput) ValidPageOrders() []PageOrder {
	validPageOrders := make([]PageOrder, 0)
	for _, pageOrder := range pi.pageOrders {
		if pageOrder.IsValid(pi.rules).IsValid {
			validPageOrders = append(validPageOrders, pageOrder)
		}
	}

	return validPageOrders
}

func (pi *ProblemInput) InvalidPageOrders() []PageOrder {
	invalidPageOrders := make([]PageOrder, 0)
	for _, pageOrder := range pi.pageOrders {
		if !pageOrder.IsValid(pi.rules).IsValid {
			invalidPageOrders = append(invalidPageOrders, pageOrder)
		}
	}

	return invalidPageOrders
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
