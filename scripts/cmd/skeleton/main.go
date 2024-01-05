package main

import (
	"flag"
	"os"
	"time"

	"github.com/jhh3/aoc/common"
	"github.com/jhh3/aoc/skeleton"
)

func main() {
	args := mustParseArgs()
	skeleton.Run(args.Day, args.Year)
}

type Args struct {
	Year int
	Day  int
}

func parseArgs() (*Args, error) {
	parsedArgs := Args{}

	fs := flag.NewFlagSet("aoc-skeleton", flag.ExitOnError)

	today := time.Now()
	fs.IntVar(&parsedArgs.Year, "year", today.Year(), "The year of the puzzle to solve")
	fs.IntVar(&parsedArgs.Day, "day", today.Day(), "The day of the puzzle to solve")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	return &parsedArgs, nil
}

func mustParseArgs() *Args {
	args, err := parseArgs()
	common.CheckErr(err, "Failed to parse args")
	return args
}
