package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_input.txt
var exampleInput string

//go:embed example_input_part2.txt
var exampleInputPart2 string

func Test_y2024d03(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: exampleInput,
				Part:  1,
				Want:  "161",
			},
			{
				Name:  "provided-example",
				Input: exampleInputPart2,
				Part:  2,
				Want:  "48",
			},
		},
	)
}
