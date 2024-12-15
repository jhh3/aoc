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
	equations := problemInput.EquationsWithSolutions()
	sumOfValues := 0
	for _, equation := range equations {
		sumOfValues += equation.Value
	}
	return strconv.Itoa(sumOfValues)
}

func (s *solver) SolvePart2(input string) string {
	// TODO
	return ""
}

type Equation struct {
	Value   int
	Numbers []int
}

func (e *Equation) HasSolution() bool {
	operations := []string{"+", "*"}
	numOperators := len(e.Numbers) - 1
	numPossibleEquations := 1 << numOperators
	for i := 0; i < numPossibleEquations; i++ {
		result := e.Numbers[0]
		for j, num := range e.Numbers[1:] {
			// shift i to get the next operator
			operatorIndex := i >> j & 1
			operator := operations[operatorIndex]
			if operator == "+" {
				result += num
			} else {
				result *= num
			}
		}

		if result == e.Value {
			return true
		}
	}

	return false
}

type ProblemInput struct {
	Equations []Equation
}

func (pi *ProblemInput) EquationsWithSolutions() []Equation {
	equations := make([]Equation, 0)
	for _, equation := range pi.Equations {
		if equation.HasSolution() {
			equations = append(equations, equation)
		}
	}
	return equations
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := ProblemInput{
		Equations: make([]Equation, 0, len(lines)),
	}

	for _, line := range lines {
		equation := parseEquation(line)
		result.Equations = append(result.Equations, equation)
	}

	return result
}

func parseEquation(line string) Equation {
	// Split on colon to separate value from numbers
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		panic("invalid equation format")
	}

	// Parse the value
	value := common.MustAtoi(strings.TrimSpace(parts[0]))

	// Parse the numbers
	numberStrs := strings.Fields(strings.TrimSpace(parts[1]))
	numbers := make([]int, 0, len(numberStrs))
	for _, numStr := range numberStrs {
		num := common.MustAtoi(numStr)
		numbers = append(numbers, num)
	}

	return Equation{
		Value:   value,
		Numbers: numbers,
	}
}
