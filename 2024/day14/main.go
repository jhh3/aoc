package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/2024/d14"
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
	problemInput.Step(100)
	fmt.Printf("Dimensions: %v\n", problemInput.Dimensions)
	safetyFactor := problemInput.ComputeSafetyfactor()

	dAns, _ := d14.Part1(input)

	fmt.Printf("Day 14 Part 1: %v\n", dAns)

	return strconv.Itoa(safetyFactor)
}

func (s *solver) SolvePart2(input string) string {
	// TODO: Implement part 2
	return ""
}

type RoomDimensions struct {
	Width, Height int
}

var (
	ExampleDimensions = RoomDimensions{11, 7}
	RealDimensions    = RoomDimensions{101, 103}
)

type Point struct {
	X, Y int
}

type Velocity struct {
	Dx, Dy int
}

type Robot struct {
	Position Point
	Velocity Velocity
}

type ProblemInput struct {
	Dimensions RoomDimensions
	Robots     []Robot
}

func (pi *ProblemInput) Step(numSteps int) {
	for i := range pi.Robots {
		// robots wrap around the room, use modulo to handle this
		pi.Robots[i].Position.X += numSteps * pi.Robots[i].Velocity.Dx
		pi.Robots[i].Position.X %= pi.Dimensions.Width

		pi.Robots[i].Position.Y = numSteps * pi.Robots[i].Velocity.Dy
		pi.Robots[i].Position.Y %= pi.Dimensions.Height

		if pi.Robots[i].Position.X < 0 {
			pi.Robots[i].Position.X += pi.Dimensions.Width
		}

		if pi.Robots[i].Position.Y < 0 {
			pi.Robots[i].Position.Y += pi.Dimensions.Height
		}
	}
}

func (pi *ProblemInput) ComputeSafetyfactor() int {
	var counts [5]int

	for _, robot := range pi.Robots {
		quadrant := pi.WhichQuadrant(robot.Position)
		counts[quadrant]++
	}

	safetyFactor := counts[0] * counts[1] * counts[2] * counts[3]

	fmt.Printf("ounts: %v\n", counts)

	return safetyFactor
}

func (pi *ProblemInput) WhichQuadrant(p Point) int {
	// 0 = top left, 1 = top right, 2 = bottom left, 3 = bottom right 4 = on middle
	if p.X < pi.Dimensions.Width/2 && p.Y < pi.Dimensions.Height/2 {
		return 0
	}
	if p.X >= pi.Dimensions.Width/2 && p.Y < pi.Dimensions.Height/2 {
		return 1
	}

	if p.X < pi.Dimensions.Width/2 && p.Y > pi.Dimensions.Height/2 {
		return 2
	}

	if p.X > pi.Dimensions.Width/2 && p.Y > pi.Dimensions.Height/2 {
		return 3
	}

	return 4
}

func (pi *ProblemInput) PrettyPrint() {
	for y := 0; y < pi.Dimensions.Height; y++ {
		for x := 0; x < pi.Dimensions.Width; x++ {
			count := 0
			for _, robot := range pi.Robots {
				if robot.Position.X == x && robot.Position.Y == y {
					count++
				}
			}
			if count > 0 {
				fmt.Printf("%d", count)
			} else {
				print(".")
			}
		}
		println()
	}
}

func parseInput(input string) ProblemInput {
	lines := common.ReadAsLines(input)

	pi := ProblemInput{}
	if len(lines) > 20 {
		pi.Dimensions = RealDimensions
	} else {
		pi.Dimensions = ExampleDimensions
	}

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("invalid line")
		}
		posStr := strings.TrimPrefix(strings.TrimSpace(parts[0]), "p=")
		posParts := strings.Split(posStr, ",")
		pos := Point{common.MustAtoi(posParts[0]), common.MustAtoi(posParts[1])}
		velStr := strings.TrimPrefix(strings.TrimSpace(parts[1]), "v=")
		velParts := strings.Split(velStr, ",")
		vel := Velocity{common.MustAtoi(velParts[0]), common.MustAtoi(velParts[1])}

		pi.Robots = append(pi.Robots, Robot{pos, vel})
	}

	return pi
}
