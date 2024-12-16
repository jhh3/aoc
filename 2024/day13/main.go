package main

import (
	_ "embed"
	"fmt"
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
	cost := problemInput.TotalCostToPrizes(false)
	return fmt.Sprintf("%v", cost)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	cost := problemInput.TotalCostToPrizes(true)
	return fmt.Sprintf("%v", cost)
}

type Point struct {
	X, Y int64
}

type Vector struct {
	X, Y int64
}

const (
	COST_OF_A_BUTTON = 3
	COST_OF_B_BUTTON = 1
)

type ClawGame struct {
	ButtonA Vector
	ButtonB Vector

	Prize Point
}

// CostToPrize returns the lowest cost to get the prize
// returns -1 if the prize is unreachable
func (cg *ClawGame) CostToPrize(startingPoint Point, seen map[Point]int) int {
	// are we past the prize?
	if startingPoint.X > cg.Prize.X || startingPoint.Y > cg.Prize.Y {
		return -1
	}

	// are we at the prize?
	if startingPoint.X == cg.Prize.X && startingPoint.Y == cg.Prize.Y {
		return 0
	}

	nextAfterA := Point{startingPoint.X + cg.ButtonA.X, startingPoint.Y + cg.ButtonA.Y}
	if _, ok := seen[nextAfterA]; !ok {
		seen[nextAfterA] = cg.CostToPrize(nextAfterA, seen)
	}
	costOfARoute := seen[nextAfterA]

	nextAfterB := Point{startingPoint.X + cg.ButtonB.X, startingPoint.Y + cg.ButtonB.Y}
	if _, ok := seen[nextAfterB]; !ok {
		seen[nextAfterB] = cg.CostToPrize(nextAfterB, seen)
	}
	costOfBRoute := seen[nextAfterB]

	if costOfARoute == -1 && costOfBRoute == -1 {
		return -1
	}

	if costOfARoute == -1 {
		return costOfBRoute + COST_OF_B_BUTTON
	}
	if costOfBRoute == -1 {
		return costOfARoute + COST_OF_A_BUTTON
	}

	if costOfARoute < costOfBRoute {
		return costOfARoute + COST_OF_A_BUTTON
	}
	return costOfBRoute + COST_OF_B_BUTTON
}

type ProblemInput struct {
	ClawGames []ClawGame
}

func (pi *ProblemInput) TotalCostToPrizes(withConversionError bool) int64 {
	totalCost := int64(0)
	for _, cg := range pi.ClawGames {
		// use system of equations to solve for aPresses and bPresses
		// should have started here .... lol
		if withConversionError {
			err := int64(10000000000000)
			cg.Prize.X += err
			cg.Prize.Y += err

			aPresses := (cg.Prize.Y*cg.ButtonB.X - cg.Prize.X*cg.ButtonB.Y) / (cg.ButtonA.Y*cg.ButtonB.X - cg.ButtonA.X*cg.ButtonB.Y)
			bPresses := (cg.Prize.X - aPresses*cg.ButtonA.X) / cg.ButtonB.X

			tx := aPresses*cg.ButtonA.X + bPresses*cg.ButtonB.X
			ty := aPresses*cg.ButtonA.Y + bPresses*cg.ButtonB.Y

			if tx == cg.Prize.X && ty == cg.Prize.Y {
				totalCost += aPresses*int64(COST_OF_A_BUTTON) + bPresses*int64(COST_OF_B_BUTTON)
			}

		} else {
			cost := cg.CostToPrize(Point{0, 0}, map[Point]int{})
			if cost == -1 { // prize is unreachable
				continue
			}
			totalCost += int64(cost)
		}
	}
	return totalCost
}

func parseInput(input string) *ProblemInput {
	pi := &ProblemInput{
		ClawGames: make([]ClawGame, 0),
	}

	lines := strings.Split(input, "\n")
	currentClawGame := ClawGame{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Button A: ") {
			secondHalf := strings.TrimPrefix(line, "Button A: ")
			parts := strings.Split(secondHalf, ", ")
			// should be 2 parts
			if len(parts) != 2 {
				panic("unexpected number of parts")
			}
			x := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[0]), "X+"))
			y := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[1]), "Y+"))
			currentClawGame.ButtonA = Vector{int64(x), int64(y)}
		}

		if strings.HasPrefix(line, "Button B: ") {
			secondHalf := strings.TrimPrefix(line, "Button B: ")
			parts := strings.Split(secondHalf, ", ")
			// should be 2 parts
			if len(parts) != 2 {
				panic("unexpected number of parts")
			}
			x := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[0]), "X+"))
			y := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[1]), "Y+"))
			currentClawGame.ButtonB = Vector{int64(x), int64(y)}
		}

		if strings.HasPrefix(line, "Prize: ") {
			secondHalf := strings.TrimPrefix(line, "Prize: ")
			parts := strings.Split(secondHalf, ", ")
			// should be 2 parts
			if len(parts) != 2 {
				panic("unexpected number of parts")
			}
			x := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[0]), "X="))
			y := common.MustAtoi(strings.TrimPrefix(strings.TrimSpace(parts[1]), "Y="))
			currentClawGame.Prize = Point{int64(x), int64(y)}

			// copy the current claw game to the list and reset the current claw game
			pi.ClawGames = append(pi.ClawGames, currentClawGame)
			currentClawGame = ClawGame{}
		}
	}

	return pi
}
