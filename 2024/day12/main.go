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
	cost := problemInput.ComputeFenceCost(false)
	return strconv.Itoa(cost)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	cost := problemInput.ComputeFenceCost(true)
	return strconv.Itoa(cost)
}

type Point struct {
	Row, Col int
}

type ProblemInput struct {
	Garden             [][]rune
	Visited            map[Point]bool
	PlantProcessedMap  map[rune]bool
	PlantToPositionMap map[rune][]Point
}

// Regions are defined as a contiguous group of plants of the same type
type Region struct {
	PlantType rune
	Points    []Point
}

func (pi *ProblemInput) ComputeFenceCost(bulkDiscount bool) int {
	cost := 0

	for plantType := range pi.PlantProcessedMap {
		regions := pi.Regions(plantType)
		for _, region := range regions {
			if bulkDiscount {
				cost += pi.ComputeCostForRegionBulkDiscount(region)
			} else {
				cost += pi.ComputeCostForRegion(region)
			}
		}
	}

	return cost
}

type Side struct {
	Start, End Point
	Label      string
}

func (pi *ProblemInput) ComputeCostForRegionBulkDiscount(region Region) int {
	// Cost defined as area * number of sides
	area := len(region.Points)
	numSides := 0

	seenSides := make(map[Side]bool)

	// NOTE: I know this could be A LOT DRYer, but was just trying to get it working

	for _, point := range region.Points {
		// a point can contribute up to 4 sides

		// 1. try top horizontal side
		topPoint := Point{Row: point.Row - 1, Col: point.Col}
		if !pi.IsInBounds(topPoint) || pi.Garden[topPoint.Row][topPoint.Col] != region.PlantType {
			// this is part of a side

			// find the start of the side
			start := point
			end := point
			for {
				newStart := Point{Row: start.Row, Col: start.Col - 1}

				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newStart) && pi.Garden[newStart.Row][newStart.Col] == region.PlantType {
					newStartTop := Point{Row: newStart.Row - 1, Col: newStart.Col}
					if !pi.IsInBounds(newStartTop) || pi.Garden[newStartTop.Row][newStartTop.Col] != region.PlantType {
						start = newStart
					} else {
						break
					}
				} else {
					break
				}
			}

			for {
				newEnd := Point{Row: end.Row, Col: end.Col + 1}
				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newEnd) && pi.Garden[newEnd.Row][newEnd.Col] == region.PlantType {
					newStartTop := Point{Row: newEnd.Row - 1, Col: newEnd.Col}
					if !pi.IsInBounds(newStartTop) || pi.Garden[newStartTop.Row][newStartTop.Col] != region.PlantType {
						end = newEnd
					} else {
						break
					}
				} else {
					break
				}
			}

			side := Side{Start: start, End: end, Label: "top"}
			if _, ok := seenSides[side]; !ok {
				numSides++
				seenSides[side] = true
			}
		}

		// 2. try bottom horizontal side
		bottomPoint := Point{Row: point.Row + 1, Col: point.Col}
		if !pi.IsInBounds(bottomPoint) || pi.Garden[bottomPoint.Row][bottomPoint.Col] != region.PlantType {
			// this is part of a side

			// find the start of the side
			start := point
			end := point
			for {
				newStart := Point{Row: start.Row, Col: start.Col - 1}

				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newStart) && pi.Garden[newStart.Row][newStart.Col] == region.PlantType {
					newStartBottom := Point{Row: newStart.Row + 1, Col: newStart.Col}
					if !pi.IsInBounds(newStartBottom) || pi.Garden[newStartBottom.Row][newStartBottom.Col] != region.PlantType {
						start = newStart
					} else {
						break
					}
				} else {
					break
				}
			}

			for {
				newEnd := Point{Row: end.Row, Col: end.Col + 1}
				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newEnd) && pi.Garden[newEnd.Row][newEnd.Col] == region.PlantType {
					newEndBottom := Point{Row: newEnd.Row + 1, Col: newEnd.Col}
					if !pi.IsInBounds(newEndBottom) || pi.Garden[newEndBottom.Row][newEndBottom.Col] != region.PlantType {
						end = newEnd
					} else {
						break
					}
				} else {
					break
				}
			}

			side := Side{Start: start, End: end, Label: "bottom"}
			if _, ok := seenSides[side]; !ok {
				numSides++
				seenSides[side] = true
			}
		}

		// 3. try left vertical side
		leftPoint := Point{Row: point.Row, Col: point.Col - 1}
		if !pi.IsInBounds(leftPoint) || pi.Garden[leftPoint.Row][leftPoint.Col] != region.PlantType {
			// this is part of a side

			// find the start of the side
			start := point
			end := point
			for {
				newStart := Point{Row: start.Row - 1, Col: start.Col}

				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newStart) && pi.Garden[newStart.Row][newStart.Col] == region.PlantType {
					newStartLeft := Point{Row: newStart.Row, Col: newStart.Col - 1}
					if !pi.IsInBounds(newStartLeft) || pi.Garden[newStartLeft.Row][newStartLeft.Col] != region.PlantType {
						start = newStart
					} else {
						break
					}
				} else {
					break
				}
			}

			for {
				newEnd := Point{Row: end.Row + 1, Col: end.Col}
				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newEnd) && pi.Garden[newEnd.Row][newEnd.Col] == region.PlantType {
					newEndLeft := Point{Row: newEnd.Row, Col: newEnd.Col - 1}
					if !pi.IsInBounds(newEndLeft) || pi.Garden[newEndLeft.Row][newEndLeft.Col] != region.PlantType {
						end = newEnd
					} else {
						break
					}
				} else {
					break
				}
			}

			side := Side{Start: start, End: end, Label: "left"}
			if _, ok := seenSides[side]; !ok {
				numSides++
				seenSides[side] = true
			}
		}

		// 4. try right vertical side
		rightPoint := Point{Row: point.Row, Col: point.Col + 1}
		if !pi.IsInBounds(rightPoint) || pi.Garden[rightPoint.Row][rightPoint.Col] != region.PlantType {
			// this is part of a side

			// find the start of the side
			start := point
			end := point
			for {
				newStart := Point{Row: start.Row - 1, Col: start.Col}

				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newStart) && pi.Garden[newStart.Row][newStart.Col] == region.PlantType {
					newStartRight := Point{Row: newStart.Row, Col: newStart.Col + 1}
					if !pi.IsInBounds(newStartRight) || pi.Garden[newStartRight.Row][newStartRight.Col] != region.PlantType {
						start = newStart
					} else {
						break
					}
				} else {
					break
				}
			}

			for {
				newEnd := Point{Row: end.Row + 1, Col: end.Col}
				// must be the same plant type and inbounds
				// add the top of it is out of bounds or not the same plant type
				if pi.IsInBounds(newEnd) && pi.Garden[newEnd.Row][newEnd.Col] == region.PlantType {
					newEndRight := Point{Row: newEnd.Row, Col: newEnd.Col + 1}
					if !pi.IsInBounds(newEndRight) || pi.Garden[newEndRight.Row][newEndRight.Col] != region.PlantType {
						end = newEnd
					} else {
						break
					}
				} else {
					break
				}
			}

			side := Side{Start: start, End: end, Label: "right"}
			if _, ok := seenSides[side]; !ok {
				numSides++
				seenSides[side] = true
			}
		}

	}

	return numSides * area
}

func (pi *ProblemInput) ComputeCostForRegion(region Region) int {
	// Cost defined as area * perimeter
	area := len(region.Points)
	perimeter := 0

	for _, point := range region.Points {
		// each point can contribute up to 4 to the perimeter
		contribution := 0
		for _, neighbor := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextPoint := Point{Row: point.Row + neighbor.Row, Col: point.Col + neighbor.Col}
			// if the neighbor is out of bounds or not the same plant type, it contributes to the perimeter
			if !pi.IsInBounds(nextPoint) || pi.Garden[nextPoint.Row][nextPoint.Col] != region.PlantType {
				contribution++
			}
		}
		perimeter += contribution
	}

	return perimeter * area
}

func (pi *ProblemInput) Regions(plantType rune) []Region {
	regions := make([]Region, 0)

	points := pi.PlantToPositionMap[plantType]
	seen := make(map[Point]bool)
	for _, point := range points {
		// Have we seen this point before?
		if _, ok := seen[point]; ok {
			continue
		}

		// Start a new region
		region := Region{PlantType: plantType, Points: make([]Point, 0)}
		pointsToProcess := []Point{point}
		pointsProcessed := make(map[Point]bool)
		for len(pointsToProcess) > 0 {
			// pop the first point
			point := pointsToProcess[0]
			pointsToProcess = pointsToProcess[1:]

			// Have we seen this point before?
			if _, ok := pointsProcessed[point]; ok {
				continue
			}

			// Add this point to the region
			region.Points = append(region.Points, point)
			pointsProcessed[point] = true

			// Add the neighbors to the points to process
			for _, neighbor := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nextPoint := Point{Row: point.Row + neighbor.Row, Col: point.Col + neighbor.Col}
				if pi.IsInBounds(nextPoint) && pi.Garden[nextPoint.Row][nextPoint.Col] == plantType {
					if _, ok := pointsProcessed[nextPoint]; !ok {
						pointsToProcess = append(pointsToProcess, nextPoint)
					}
				}
			}
		}

		// Add the region to the list of regions
		regions = append(regions, region)

		// Mark all the points in the region as seen
		for _, point := range region.Points {
			seen[point] = true
		}
	}

	return regions
}

func (pi *ProblemInput) IsInBounds(point Point) bool {
	return point.Row >= 0 && point.Row < len(pi.Garden) && point.Col >= 0 && point.Col < len(pi.Garden[0])
}

func parseInput(input string) *ProblemInput {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	pi := &ProblemInput{
		Garden:             make([][]rune, len(lines)),
		Visited:            make(map[Point]bool),
		PlantProcessedMap:  make(map[rune]bool),
		PlantToPositionMap: make(map[rune][]Point),
	}

	for i, row := range lines {
		pi.Garden[i] = []rune(row)
		for j, element := range row {
			pi.PlantProcessedMap[element] = false
			point := Point{Row: i, Col: j}
			if _, ok := pi.PlantToPositionMap[element]; !ok {
				pi.PlantToPositionMap[element] = make([]Point, 0)
			}
			pi.PlantToPositionMap[element] = append(pi.PlantToPositionMap[element], point)
		}
	}

	return pi
}
