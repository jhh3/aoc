package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed test_example.txt
var exampleInput string

func Test_solution(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{
			{
				Name:  "day03-example",
				Input: exampleInput,
				Part:  1,
				Want:  "4361",
			},
			{
				Name:  "day03-example",
				Input: exampleInput,
				Part:  2,
				Want:  "467835",
			},
		},
	)
}
