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
	// TODO
	problemInput := parseInput(input)
	problemInput.ExpandDiskMap()
	problemInput.Compact()
	checksum := problemInput.ComputeChecksum()

	// fmt.Printf("DiskMap: %v\n", problemInput.DiskMap)
	// fmt.Printf("DiskMap: %v\n", problemInput.ExpandedDiskMap)
	// fmt.Printf("DiskMap: %v\n", problemInput.CompactExpandedDiskMap)
	// fmt.Printf("Checksum: %v\n", checksum)

	return strconv.Itoa(checksum)
}

func (s *solver) SolvePart2(input string) string {
	// TODO
	return ""
}

type ProblemInput struct {
	DiskMap                []int
	ExpandedDiskMap        []int // -1 is free
	CompactExpandedDiskMap []int
}

func (pi *ProblemInput) ComputeChecksum() int {
	sum := 0
	for idx, fileId := range pi.CompactExpandedDiskMap {
		if fileId == -1 {
			break
		}
		sum += idx * fileId
	}
	return sum
}

func (pi *ProblemInput) Compact() {
	// copy the expanded disk map
	pi.CompactExpandedDiskMap = make([]int, len(pi.ExpandedDiskMap))
	copy(pi.CompactExpandedDiskMap, pi.ExpandedDiskMap)

	startPtr := 0
	endPtr := len(pi.CompactExpandedDiskMap) - 1

	for startPtr < endPtr {
		// do we want to move what's at endPtr
		if pi.CompactExpandedDiskMap[endPtr] == -1 {
			// shift left if looking at free space
			endPtr--
			continue
		}

		// is startPtr already pointing at a file?
		if pi.CompactExpandedDiskMap[startPtr] != -1 {
			startPtr++
			continue
		}

		// swap
		pi.CompactExpandedDiskMap[startPtr], pi.CompactExpandedDiskMap[endPtr] = pi.CompactExpandedDiskMap[endPtr], pi.CompactExpandedDiskMap[startPtr]
	}
}

func (pi *ProblemInput) ExpandDiskMap() {
	currentFileIdx := 0
	for i, diskValue := range pi.DiskMap {
		isFiles := i%2 == 0
		for j := 0; j < diskValue; j++ {
			if isFiles {
				pi.ExpandedDiskMap = append(pi.ExpandedDiskMap, currentFileIdx)
			} else {
				pi.ExpandedDiskMap = append(pi.ExpandedDiskMap, -1)
			}
		}
		if isFiles && diskValue > 0 {
			currentFileIdx++
		}
	}
}

func parseInput(input string) *ProblemInput {
	problemInput := &ProblemInput{
		DiskMap:                []int{},
		ExpandedDiskMap:        []int{},
		CompactExpandedDiskMap: []int{},
	}

	cleanInput := strings.TrimSpace(input)
	for _, char := range cleanInput {
		problemInput.DiskMap = append(problemInput.DiskMap, int(char-'0'))
	}

	return problemInput
}
