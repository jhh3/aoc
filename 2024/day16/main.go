package main

import (
	_ "embed"
	"sort"
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
	score := problemInput.Solve()
	return strconv.Itoa(score)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	score := problemInput.Solve2()
	return strconv.Itoa(score)
}

const (
	MOVE_COST    = 1
	TURN_90_COST = 1000
	PENALTY      = 100000000000
)

type Point struct {
	Row, Col int
}

func (p *Point) Equal(other Point) bool {
	return p.Row == other.Row && p.Col == other.Col
}

type Position struct {
	Point     Point
	Direction int // 0 = up, 1 = right, 2 = down, 3 = left
}

type ProblemInput struct {
	Maze [][]rune

	Start Position
	Goal  Point
}

type QueueItem struct {
	Position Position
	Score    int
	Path     []Point
}

func (pi *ProblemInput) Solve() int {
	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{Position: pi.Start, Score: 0})
	visited := make(map[Position]bool)

	for len(queue) > 0 {
		// pop lowest score
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Score < queue[j].Score
		})
		// pop
		current := queue[0]
		queue = queue[1:]

		if pi.Goal.Equal(current.Position.Point) {
			return current.Score
		}
		if _, ok := visited[current.Position]; ok {
			continue
		}
		visited[current.Position] = true

		possibleMoves := pi.PossibleMoves(current.Position)
		for next, cost := range possibleMoves {
			queue = append(queue, QueueItem{Position: next, Score: current.Score + cost})
		}
	}

	return -1
}

func (pi *ProblemInput) Solve2() int {
	queue := make([]QueueItem, 0)
	queue = append(queue, QueueItem{Position: pi.Start, Score: 0, Path: []Point{pi.Start.Point}})
	visited := make(map[Position]int)
	targetScore := pi.Solve()
	optimalPoints := make(map[Point]bool)
	pointCount := 0

	for len(queue) > 0 {
		// pop lowest score
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Score < queue[j].Score
		})
		// pop
		current := queue[0]
		queue = queue[1:]

		if current.Score > targetScore {
			continue
		}

		if score, ok := visited[current.Position]; ok && score < current.Score {
			continue
		}
		visited[current.Position] = current.Score

		if pi.Goal.Equal(current.Position.Point) && current.Score == targetScore {
			for _, point := range current.Path {
				if _, ok := optimalPoints[point]; !ok {
					optimalPoints[point] = true
					pointCount++
				}
			}
			continue
		}

		possibleMoves := pi.PossibleMoves(current.Position)
		for next, cost := range possibleMoves {
			if next.Direction != current.Position.Direction {
				queue = append(queue, QueueItem{Position: next, Score: current.Score + cost, Path: current.Path})
			} else {
				newPath := make([]Point, len(current.Path))
				copy(newPath, current.Path)
				newPath = append(newPath, next.Point)
				queue = append(queue, QueueItem{Position: next, Score: current.Score + cost, Path: newPath})
			}

		}
	}

	return pointCount
}

func (pi *ProblemInput) PossibleMoves(from Position) map[Position]int {
	nextPositionToCost := make(map[Position]int)
	// Move forward along current direction
	// Direction int // 0 = up, 1 = right, 2 = down, 3 = left
	newPoint := from.Point
	switch from.Direction {
	case 0:
		newPoint.Row--
	case 1:
		newPoint.Col++
	case 2:
		newPoint.Row++
	case 3:
		newPoint.Col--
	}

	// If we remain in the maze and don't hit a wall
	if pi.IsInMaze(newPoint) && pi.Maze[newPoint.Row][newPoint.Col] != '#' {
		nextPositionToCost[Position{Point: newPoint, Direction: from.Direction}] = MOVE_COST
	}

	// Turn 90 degrees in either direction
	for _, turn := range []int{1, -1} {
		newDirection := (from.Direction + turn) % 4
		if newDirection < 0 {
			newDirection += 4
		}

		nextPositionToCost[Position{Point: from.Point, Direction: newDirection}] = TURN_90_COST
	}

	return nextPositionToCost
}

func (pi *ProblemInput) IsInMaze(p Point) bool {
	return p.Row >= 0 && p.Row < len(pi.Maze) && p.Col >= 0 && p.Col < len(pi.Maze[0])
}

func parseInput(input string) *ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	pi := &ProblemInput{
		Maze: make([][]rune, len(lines)),
		Start: Position{
			Direction: 1, // always start facing right
		},
	}

	for i, line := range lines {
		pi.Maze[i] = []rune(line)
		for j, element := range line {
			if element == 'S' {
				pi.Start.Point = Point{Row: i, Col: j}
			}

			if element == 'E' {
				pi.Goal = Point{Row: i, Col: j}
			}
		}
	}

	return pi
}
