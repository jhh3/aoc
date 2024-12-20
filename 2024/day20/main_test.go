package main

import (
	_ "embed"
	"testing"

	"github.com/jhh3/aoc/common"
)

//go:embed example_input.txt
var exampleInput string

func Test_y2024d20(t *testing.T) {
	common.RunTests(
		&solver{},
		t,
		[]common.Test{},
	)
}
