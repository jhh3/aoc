package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhh3/aoc/common"
)

func main() {
	common.RunFromSolver(&solver{}, 2)
}

type solver struct{}

// Determine which games would have been possible if the bag had been loaded
// with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum
// of the IDs of those games?
func (ps *solver) SolvePart1(input string) string {
	games := parseInput(input)
	sum := 0

	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			sum += game.id
		}
	}

	return strconv.Itoa(sum)
}

func (ps *solver) SolvePart2(input string) string {
	return ""
}

// Input parser

// Example input
// ...
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

// Global regexes
var gameRegex = regexp.MustCompile(`Game (\d+):`)

type Round struct {
	blue  int
	red   int
	green int
}

func (r *Round) IsPossible(red, green, blue int) bool {
	return blue >= r.blue && red >= r.red && green >= r.green
}

type Game struct {
	id     int
	rounds []Round
}

func (g *Game) Print() {
	fmt.Printf("Game %d:\n", g.id)
	for _, round := range g.rounds {
		fmt.Printf("\t%d blue, %d red, %d green\n", round.blue, round.red, round.green)
	}
}

func (g *Game) IsPossible(red, green, blue int) bool {
	for _, round := range g.rounds {
		if !round.IsPossible(red, green, blue) {
			return false
		}
	}
	return true
}

func parseInput(input string) []Game {
	result := []Game{}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}

		result = append(result, parseGame(cleanLine))
	}

	return result
}

func parseGame(input string) Game {
	// 1. Get the game id
	gameId := gameRegex.FindStringSubmatch(input)[1]

	// 2. Get the rounds

	// split on colon
	roundInfo := strings.TrimSpace(strings.Split(input, ":")[1])
	rounds := []Round{}
	for _, roundStr := range strings.Split(roundInfo, ";") {
		cleanRoundStr := strings.TrimSpace(roundStr)
		rounds = append(rounds, parseRound(cleanRoundStr))
	}

	return Game{
		id:     MustAtoi(gameId),
		rounds: rounds,
	}
}

func parseRound(input string) Round {
	r := Round{}

	blockInfos := strings.Split(input, ",")
	for _, blockInfo := range blockInfos {
		cleanBlockInfo := strings.TrimSpace(blockInfo)
		pieces := strings.Split(cleanBlockInfo, " ")
		quantity := MustAtoi(pieces[0])
		color := pieces[1]
		switch color {
		case "blue":
			r.blue = quantity
		case "red":
			r.red = quantity
		case "green":
			r.green = quantity
		default:
			panic(fmt.Sprintf("Unknown color: %s", color))
		}
	}

	return r
}

func MustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	common.CheckErr(err, "Failed to convert to integer")
	return result
}
