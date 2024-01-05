package main

import (
	"os"

	"github.com/jhh3/aoc/common"
)

func main() {
	args := common.MustParseInputGetterFlags(os.Args[1:], false)
	reader := common.NewProblemReader(args)
	reader.MustGetInput()
}
