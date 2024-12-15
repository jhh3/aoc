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
	problemInput.Blink(25, false)
	numberOfStones := problemInput.CountStones(false)
	return strconv.Itoa(numberOfStones)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	problemInput.Blink(75, false)
	numberOfStones := problemInput.CountStones(false)
	return strconv.Itoa(numberOfStones)
}

type ProblemInput struct {
	Stones    []int
	StonesMap map[int]int // key is stone, value is count
}

func (pi *ProblemInput) CountStones(part1 bool) int {
	if part1 {
		return len(pi.Stones)
	} else {
		sum := 0
		for _, count := range pi.StonesMap {
			sum += count
		}
		return sum
	}
}

func (pi *ProblemInput) Blink(times int, part1 bool) {
	for i := 0; i < times; i++ {
		if part1 {
			pi.SingleBlinkPart1()
		} else {
			pi.SingleBlinkPart2()
		}
	}
}

func (pi *ProblemInput) SingleBlinkPart2() {
	newStonesMap := make(map[int]int)
	for stone, count := range pi.StonesMap {
		if stone == 0 {
			if _, ok := newStonesMap[1]; !ok {
				newStonesMap[1] = 0
			}
			newStonesMap[1] += count
		} else {
			numDigits := common.NumDigits(stone)
			if numDigits%2 == 0 {
				strNum := strconv.Itoa(stone)

				leftHalf := strNum[:numDigits/2]
				rightHalf := strNum[numDigits/2:]

				leftStone := common.MustAtoi(leftHalf)
				rightStone := common.MustAtoi(rightHalf)

				if _, ok := newStonesMap[leftStone]; !ok {
					newStonesMap[leftStone] = 0
				}
				newStonesMap[leftStone] += count
				if _, ok := newStonesMap[rightStone]; !ok {
					newStonesMap[rightStone] = 0
				}
				newStonesMap[rightStone] += count
			} else {
				newStone := stone * 2024
				if _, ok := newStonesMap[newStone]; !ok {
					newStonesMap[newStone] = 0
				}
				newStonesMap[newStone] += count
			}
		}
	}
	pi.StonesMap = newStonesMap
}

// SingleBlinkPart1 applies the follwing rules:
// - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
// - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
func (pi *ProblemInput) SingleBlinkPart1() {
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
		Stones:    make([]int, len(strStones)),
		StonesMap: make(map[int]int),
	}
	for i, s := range strStones {
		val := common.MustAtoi(s)
		pi.Stones[i] = val
		if _, ok := pi.StonesMap[val]; !ok {
			pi.StonesMap[val] = 0
		}
		pi.StonesMap[val]++
	}

	return pi
}
