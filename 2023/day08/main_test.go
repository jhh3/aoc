package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_part1_input.txt
var examplePart1Input string

//go:embed example_part2_input.txt
var examplePart2Input string

func Test_y2023d08(t *testing.T) {
	exp1Input := strings.TrimRight(examplePart1Input, "\n")
	exp2Input := strings.TrimRight(examplePart2Input, "\n")

	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: exp1Input,
				Part:  1,
				Want:  "2",
			},
			{
				Name:  "provided-example",
				Input: exp2Input,
				Part:  2,
				Want:  "6",
			},
		},
	)
}
