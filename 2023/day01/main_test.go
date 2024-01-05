package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_part1_input.txt
var examplePart1Input string

//go:embed example_part2_input.txt
var examplePart2Input string

func Test_y2023d01(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: examplePart1Input,
				Part:  1,
				Want:  "142",
			},
			{
				Name:  "provided-example",
				Input: examplePart2Input,
				Part:  2,
				Want:  "281",
			},
		},
	)
}
