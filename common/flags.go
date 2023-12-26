package common

import (
	"flag"
	"fmt"
)

type ProblemSolverFlags struct {
	Year           int
	Day            int
	Part           int
	CookieFilePath string
	CacheDir       string
	BaseUrl        string
}

func ParseFlags(args []string, debug bool) (*ProblemSolverFlags, error) {
	parsedFlags := ProblemSolverFlags{}
	fs := flag.NewFlagSet("aoc", flag.ContinueOnError)

	fs.IntVar(&parsedFlags.Part, "part", 1, "part 1 or 2")
	fs.StringVar(&parsedFlags.CookieFilePath, "cookie", "cookie.txt", "path to cookie file")
	fs.StringVar(&parsedFlags.CacheDir, "cache", "inputs", "path to cache directory")
	fs.StringVar(&parsedFlags.BaseUrl, "baseurl", AOCBaseURL, "base url")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if debug {
		fmt.Println("Parsed flags:")
		fmt.Println("\tPart:", parsedFlags.Part)
		fmt.Println("\tCookieFilePath:", parsedFlags.CookieFilePath)
		fmt.Println("\tCacheDir:", parsedFlags.CacheDir)
		fmt.Println("\tBaseUrl:", parsedFlags.BaseUrl)
	}

	return &parsedFlags, nil
}

func MustParseFlags(args []string, year, day int, debug bool) *ProblemSolverFlags {
	flags, err := ParseFlags(args, debug)
	CheckErr(err, "Failed to parse flags")
	flags.Day = day
	flags.Year = year
	return flags
}

func (ps *ProblemSolverFlags) InputUrl() string {
	return fmt.Sprintf("%s/%d/day/%d/input", ps.BaseUrl, ps.Year, ps.Day)
}

func (ps *ProblemSolverFlags) CacheKey() string {
	return fmt.Sprintf("%s/input-%d-%d.txt", ps.CacheDir, ps.Year, ps.Day)
}
