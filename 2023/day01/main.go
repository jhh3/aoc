package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jhh3/aoc/common"
)

func main() {
	common.RunFromSolver(&solver{})
}

type solver struct{}

func (ps *solver) SolvePart1(input string) string {
	fmt.Println("\tSolving problem")
	lines := strings.Split(string(input), "\n")
	sum := 0

	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}

		// iterate over each character in the line
		var firstDigit rune
		var lastDigit rune
		for _, r := range cleanLine {
			// check if the current character is a digit
			if unicode.IsDigit(r) {
				if firstDigit == 0 {
					firstDigit = r
				}
				lastDigit = r
			}
		}

		// convert to two digit integer
		val, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		common.CheckErr(err, "Failed to convert to integer")

		sum += val
	}

	return strconv.Itoa(sum)
}

func (ps *solver) SolvePart2(input string) string {
	fmt.Println("\tSolving problem")
	lines := strings.Split(string(input), "\n")
	sum := 0

	// written number to number map
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}

		// iterate over each character in the line
		var firstDigit rune
		var lastDigit rune

		// lineLen := len(cleanLine)
		for idx, r := range cleanLine {
			// check if the current character is a digit
			if unicode.IsDigit(r) {
				if firstDigit == 0 {
					firstDigit = r
					lastDigit = firstDigit
				}
				lastDigit = r
			}

			// check if we're rooted at a number word
			for word, val := range numberMap {
				// first letter same
				if r == rune(word[0]) {
					// check bounds
					if idx+len(word) > len(cleanLine) {
						continue
					}
					// slice out word lenght from line
					s := cleanLine[idx : idx+len(word)]
					if s == word {
						if firstDigit == 0 {
							firstDigit = rune(val[0])
							lastDigit = firstDigit
						}
						lastDigit = rune(val[0])
					}
				}
			}
		}

		// convert to two digit integer
		val, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		common.CheckErr(err, "Failed to convert to integer")

		sum += val
	}

	return strconv.Itoa(sum)
}
