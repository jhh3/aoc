package main

import (
	_ "embed"
	"fmt"
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
	steps := 12
	if len(problemInput.Obstacles) > 30 {
		steps = 1024
	}
	shortestLen, path := problemInput.FindShortestPath(steps)
	problemInput.PrettyPrint(steps, path)
	println()
	problemInput.PrettyPrint(steps, nil)
	return strconv.Itoa(shortestLen)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	p := problemInput.FindUnreachableStep()
	return p.String()
}

//--------------------------------------------------------------------

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.Y, p.X)
}

type ProblemInput struct {
	Start Point
	Exit  Point

	Obstacles   []Point
	ObstacleMap map[Point]int
}

func (pi *ProblemInput) FindUnreachableStep() Point {
	maxStep := len(pi.Obstacles) - 1
	for step := 0; step < maxStep; step++ {
		shortestLen, _ := pi.FindShortestPath(step)
		println(step, shortestLen)
		if shortestLen == -1 {
			return pi.Obstacles[step-1]

		}
	}
	return Point{-1, -1}

}

func (pi *ProblemInput) PrettyPrint(maxObstacleIdx int, path []Point) {
	pathMap := make(map[Point]bool)
	for _, p := range path {
		pathMap[p] = true
	}
	for x := 0; x <= pi.Exit.X; x++ {
		for y := 0; y <= pi.Exit.Y; y++ {
			p := Point{x, y}
			if p == pi.Start {
				print("O")
			} else if p == pi.Exit {
				print("O")
			} else if _, ok := pathMap[p]; ok {
				print("O")
			} else if idx, ok := pi.ObstacleMap[p]; ok {
				if idx < maxObstacleIdx {
					print("#")
				} else {
					print(".")
				}
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func (pi *ProblemInput) FindShortestPath(maxObstacleIdx int) (int, []Point) {
	// BFS with distance tracking
	type QueueItem struct {
		pos  Point
		dist int
		Path []Point
	}

	queue := []QueueItem{{pos: pi.Start, dist: 0, Path: []Point{pi.Start}}}
	costs := make(map[Point]int)
	costs[pi.Start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // pop front

		if current.pos == pi.Exit {
			return current.dist, current.Path
		}

		// right, left, down, up
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{current.pos.X + dir.X, current.pos.Y + dir.Y}

			// Check bounds
			if next.X < 0 || next.Y < 0 || next.X > pi.Exit.X || next.Y > pi.Exit.Y {
				continue
			}

			// Check obstacles
			if idx, ok := pi.ObstacleMap[next]; ok {
				if idx < maxObstacleIdx {
					continue
				}
			}

			// Check cost
			if cost, ok := costs[next]; ok {
				if current.dist+1 >= cost {
					continue
				}
			}
			costs[next] = current.dist + 1

			pathcopy := make([]Point, len(current.Path))
			copy(pathcopy, current.Path)
			queue = append(queue, QueueItem{pos: next, dist: current.dist + 1, Path: append(pathcopy, next)})
		}
	}

	return -1, nil
}

var (
	ExampleExit = Point{6, 6}
	RealExit    = Point{70, 70}
)

func parseInput(input string) ProblemInput {
	obstactleStrs := common.ReadAsLines(input)

	pi := ProblemInput{
		Start:       Point{0, 0},
		Obstacles:   make([]Point, 0),
		ObstacleMap: make(map[Point]int),
		Exit:        ExampleExit,
	}

	if len(obstactleStrs) > 30 {
		pi.Exit = RealExit
	}

	for i, obsStr := range obstactleStrs {
		parts := strings.Split(strings.TrimSpace(obsStr), ",")
		x, y := common.MustAtoi(parts[0]), common.MustAtoi(parts[1])
		row, col := y, x
		obs := Point{row, col}
		pi.Obstacles = append(pi.Obstacles, obs)
		pi.ObstacleMap[obs] = i
	}

	return pi
}
