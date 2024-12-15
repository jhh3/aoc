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
	problemInput.Blink(25)
	numberOfStones := len(problemInput.Stones)
	return strconv.Itoa(numberOfStones)
}

func (s *solver) SolvePart2(input string) string {
	// TODO
	return ""
}

type ProblemInput struct {
	Stones []int
}

func (pi *ProblemInput) Blink(times int) {
	for i := 0; i < times; i++ {
		pi.SingleBlink()
	}
}

// SingleBlink applies the follwing rules:
// - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func (pi *ProblemInput) SingleBlink() {
	newStones := make([]int, 0)

	for _, stone := range pi.Stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else {
			numDigits := common.NumDigits(stone)
			if numDigits%2 == 0 {
				strNum := strconv.Itoa(stone)

				leftHalf := strNum[:numDigits/2]
				rightHalf := strNum[numDigits/2:]

				leftStone := common.MustAtoi(leftHalf)
				rightStone := common.MustAtoi(rightHalf)

				newStones = append(newStones, leftStone)
				newStones = append(newStones, rightStone)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
	}

	pi.Stones = newStones
}

func parseInput(input string) *ProblemInput {
	cleanInput := strings.TrimSpace(input)
	strStones := strings.Fields(cleanInput)

	pi := &ProblemInput{
		Stones: make([]int, len(strStones)),
	}
	for i, s := range strStones {
		pi.Stones[i] = common.MustAtoi(s)
	}

	return pi
}
