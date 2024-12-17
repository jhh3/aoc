package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_input.txt
var exampleInput string

//go:embed example_input2.txt
var exampleInput2 string

func Test_y2024d17(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: exampleInput,
				Part:  1,
				Want:  "4,6,3,5,6,3,5,2,1,0",
			},
			{
				Name:  "provided-example",
				Input: exampleInput2,
				Part:  2,
				Want:  "117440",
			},
		},
	)
}
