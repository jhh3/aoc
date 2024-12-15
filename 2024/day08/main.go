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
	antinodeCount := problemInput.AnnotateAntinodesPart1()
	return strconv.Itoa(antinodeCount)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	antinodeCount := problemInput.AnnotateAntinodesPart2()
	return strconv.Itoa(antinodeCount)
}

type Point struct {
	Row, Col int
}

func (p *Point) Equal(other Point) bool {
	return p.Row == other.Row && p.Col == other.Col
}

type ProblemInput struct {
	Grid              [][]rune
	AntinodeLocations [][]rune

	AntennaLocations map[rune][]Point
}

func (pi *ProblemInput) AnnotateAntinodesPart1() int {
	count := 0

	// For each antenna type, mark all antinodes created by each pair of antennas
	for _, antennaLocations := range pi.AntennaLocations {
		for _, firstAntenna := range antennaLocations {
			for _, secondAntenna := range antennaLocations {
				// Skip if the same antenna
				if firstAntenna.Equal(secondAntenna) {
					continue
				}

				// Mark all antinodes created by this pair of antennas
				deltaRow := secondAntenna.Row - firstAntenna.Row
				deltaCol := secondAntenna.Col - firstAntenna.Col

				possibleAntinode1 := Point{secondAntenna.Row + deltaRow, secondAntenna.Col + deltaCol}
				possibleAntinode2 := Point{firstAntenna.Row - deltaRow, firstAntenna.Col - deltaCol}

				for _, possibleAntinode := range []Point{possibleAntinode1, possibleAntinode2} {
					if pi.IsInGrid(possibleAntinode) {
						if pi.AntinodeLocations[possibleAntinode.Row][possibleAntinode.Col] == '.' {
							pi.AntinodeLocations[possibleAntinode.Row][possibleAntinode.Col] = '#'
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func (pi *ProblemInput) AnnotateAntinodesPart2() int {
	count := 0

	// For each antenna type, mark all antinodes created by each pair of antennas
	for _, antennaLocations := range pi.AntennaLocations {
		for _, firstAntenna := range antennaLocations {
			for _, secondAntenna := range antennaLocations {
				// Skip if the same antenna
				if firstAntenna.Equal(secondAntenna) {
					continue
				}

				// Mark all antinodes created by this pair of antennas
				deltaRow := secondAntenna.Row - firstAntenna.Row
				deltaCol := secondAntenna.Col - firstAntenna.Col

				lastAntinode := Point{firstAntenna.Row, firstAntenna.Col}

				// Mark the first antinode
				if pi.AntinodeLocations[lastAntinode.Row][lastAntinode.Col] == '.' {
					pi.AntinodeLocations[lastAntinode.Row][lastAntinode.Col] = '#'
					count++
				}

				directions := []int{-1, 1}
				for _, direction := range directions {
					for {
						nextAntinode := Point{lastAntinode.Row + direction*deltaRow, lastAntinode.Col + direction*deltaCol}
						if !pi.IsInGrid(nextAntinode) {
							break
						}

						if pi.AntinodeLocations[nextAntinode.Row][nextAntinode.Col] == '.' {
							pi.AntinodeLocations[nextAntinode.Row][nextAntinode.Col] = '#'
							count++
						}

						lastAntinode = nextAntinode
					}
				}
			}
		}
	}

	return count
}

func (pi *ProblemInput) IsInGrid(p Point) bool {
	return p.Row >= 0 && p.Row < len(pi.Grid) && p.Col >= 0 && p.Col < len(pi.Grid[p.Row])
}

func (pi *ProblemInput) String() string {
	var sb strings.Builder
	for _, row := range pi.Grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (pi *ProblemInput) StringAntinodeLocations() string {
	var sb strings.Builder
	for _, row := range pi.AntinodeLocations {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func parseInput(input string) ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := ProblemInput{
		Grid:              make([][]rune, len(lines)),
		AntinodeLocations: make([][]rune, len(lines)),
		AntennaLocations:  map[rune][]Point{},
	}
	for i, row := range lines {
		result.Grid[i] = make([]rune, len(row))
		result.AntinodeLocations[i] = make([]rune, len(row))
		for j, element := range row {
			result.Grid[i][j] = element
			result.AntinodeLocations[i][j] = '.'

			// if element is lowercase letter, uppercase letter, or digit
			if (element >= 'a' && element <= 'z') || (element >= 'A' && element <= 'Z') || (element >= '0' && element <= '9') {
				if _, ok := result.AntennaLocations[element]; !ok {
					result.AntennaLocations[element] = []Point{}
				}
				result.AntennaLocations[element] = append(result.AntennaLocations[element], Point{i, j})
			}
		}
	}

	return result
}
