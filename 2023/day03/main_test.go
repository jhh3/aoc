package main

import (
	"testing"

	"github.com/jhh3/aoc/common"
)

const inputStr = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func Test_solution(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "day03-example",
				Input: inputStr,
				Part:  1,
				Want:  "4361",
			},
			{
				Name:  "day03-example",
				Input: inputStr,
				Part:  2,
				Want:  "467835",
			},
		},
	)
}
