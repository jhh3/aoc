package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_input.txt
var exampleInput string

func Test_y2023d04(t *testing.T) {
	cleanExInput := strings.TrimRight(exampleInput, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "provided-example",
				Input: cleanExInput,
				Part:  1,
				Want:  "13",
			},
			{
				Name:  "provided-example",
				Input: cleanExInput,
				Part:  2,
				Want:  "30",
			},
		},
	)
}
