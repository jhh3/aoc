package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_input.txt
var exampleInput string

func Test_y2024d05(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: exampleInput,
				Part:  1,
				Want:  "143",
			},
			{
				Name:  "provided-example",
				Input: exampleInput,
				Part:  2,
				Want:  "123",
			},
		},
	)
}
