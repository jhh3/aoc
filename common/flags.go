package common

import (
	"flag"
	"fmt"
)

type ProblemSolverFlags struct {
	Part int
}

func ParseSolverFlags(args []string, debug bool) (*ProblemSolverFlags, error) {
	parsedFlags := ProblemSolverFlags{}
	fs := flag.NewFlagSet("aoc", flag.ContinueOnError)

	fs.IntVar(&parsedFlags.Part, "part", 1, "part 1 or 2")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if debug {
		fmt.Println("Parsed flags:")
		fmt.Println("\tPart:", parsedFlags.Part)
	}

	return &parsedFlags, nil
}

func MustParseSolverFlags(args []string, debug bool) *ProblemSolverFlags {
	flags, err := ParseSolverFlags(args, debug)
	CheckErr(err, "Failed to parse flags")
	return flags
}

type InputGetterFlags struct {
	Year           int
	Day            int
	CookieFilePath string
	BaseUrl        string
}

func ParseInputGetterFlags(args []string, debug bool) (*InputGetterFlags, error) {
	parsedFlags := InputGetterFlags{}
	fs := flag.NewFlagSet("aoc", flag.ContinueOnError)

	fs.IntVar(&parsedFlags.Year, "year", 2023, "year")
	fs.IntVar(&parsedFlags.Day, "day", 1, "day")
	fs.StringVar(&parsedFlags.CookieFilePath, "cookie", "cookie.txt", "path to cookie file")
	fs.StringVar(&parsedFlags.BaseUrl, "baseurl", AOCBaseURL, "base url")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if debug {
		fmt.Println("Parsed flags:")
		fmt.Println("\tYear:", parsedFlags.Year)
		fmt.Println("\tDay:", parsedFlags.Day)
		fmt.Println("\tCookieFilePath:", parsedFlags.CookieFilePath)
		fmt.Println("\tBaseUrl:", parsedFlags.BaseUrl)
	}

	return &parsedFlags, nil
}

func MustParseInputGetterFlags(args []string, debug bool) *InputGetterFlags {
	flags, err := ParseInputGetterFlags(args, debug)
	CheckErr(err, "Failed to parse flags")
	return flags
}

func (ig *InputGetterFlags) InputUrl() string {
	return fmt.Sprintf("%s/%d/day/%d/input", ig.BaseUrl, ig.Year, ig.Day)
}

func (ig *InputGetterFlags) CacheKey() string {
	return fmt.Sprintf("%d/%02d/input.txt", ig.Year, ig.Day)
}
