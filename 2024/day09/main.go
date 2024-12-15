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
	problemInput.ExpandDiskMap()
	problemInput.CompactPart1()
	checksum := problemInput.ComputeChecksum()
	return strconv.Itoa(checksum)
}

func (s *solver) SolvePart2(input string) string {
	problemInput := parseInput(input)
	problemInput.ExpandDiskMap()
	problemInput.CompactPart2()
	checksum := problemInput.ComputeChecksum()
	return strconv.Itoa(checksum)
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
			continue
		}
		sum += idx * fileId
	}
	return sum
}

func (pi *ProblemInput) CompactPart1() {
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

func (pi *ProblemInput) CompactPart2() {
	// copy the expanded disk map
	pi.CompactExpandedDiskMap = make([]int, len(pi.ExpandedDiskMap))
	copy(pi.CompactExpandedDiskMap, pi.ExpandedDiskMap)

	startPtr := 0
	endPtr := len(pi.CompactExpandedDiskMap) - 1
	for startPtr < endPtr {
		endValue := pi.CompactExpandedDiskMap[endPtr]

		// do we want to move what's at endPtr
		if endValue == -1 {
			endPtr--
			continue
		}

		// is startPtr already pointing at a file?
		if pi.CompactExpandedDiskMap[startPtr] != -1 {
			startPtr++
			continue
		}

		// We're going to try our best to move this file
		tmpStartPtr := startPtr

		// how big is the file?
		fileLength := 0
		for i := endPtr; i >= startPtr; i-- {
			if pi.CompactExpandedDiskMap[i] == endValue {
				fileLength++
			} else {
				break
			}
		}

		// look for contiguous free space
		moveIdx := -1
		for tmpStartPtr < endPtr {
			// is startPtr already pointing at a file?
			if pi.CompactExpandedDiskMap[tmpStartPtr] != -1 {
				tmpStartPtr++
				continue
			}

			// how much contiguous free space is there?
			freeSpace := 0
			for i := tmpStartPtr; i < endPtr; i++ {
				if pi.CompactExpandedDiskMap[i] == -1 {
					freeSpace++
				} else {
					break
				}
			}

			if freeSpace >= fileLength {
				moveIdx = tmpStartPtr
				break
			} else {
				tmpStartPtr += freeSpace
			}
		}

		// can we move the file?
		if moveIdx != -1 {
			// move the file
			for i := 0; i < fileLength; i++ {
				pi.CompactExpandedDiskMap[moveIdx+i] = endValue
				pi.CompactExpandedDiskMap[endPtr-i] = -1
			}
		}

		// In either case, move the endPtr
		endPtr -= fileLength
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
