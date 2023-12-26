package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const _baseUrl = "https://adventofcode.com"
const _inputUrl = "https://adventofcode.com/2023/day/1/input"

func main() {
	var part int
	var year int
	var day int
	var cookieFilePath string
	var cacheDir string
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.IntVar(&year, "year", 2023, "year")
	flag.IntVar(&day, "day", 1, "day")
	flag.StringVar(&cookieFilePath, "cookie", "cookie.txt", "path to cookie file")
	flag.StringVar(&cacheDir, "cache", "inputs", "path to cache directory")
	flag.Parse()
	fmt.Printf("Running part %d of day %d of year %d\n", part, day, year)

	fmt.Println("Solving problem")

	ps := NewProblemSolver(cookieFilePath, _baseUrl, cacheDir, year, day, part)
	result := ps.Solve()

	fmt.Println("Result:", result)
}

type IPuzzleSolver interface {
	Solve() string
	GetInput() (string, error)
}

type ProblemSolver struct {
	part           int
	baseUrl        string
	inputUrl       string
	cacheKey       string
	cookieFilePath string
	httpClient     *http.Client
}

func NewProblemSolver(cookieFilePath string, baseUrl string, cacheDir string, year int, day int, part int) IPuzzleSolver {
	inputUrl := fmt.Sprintf("%s/%d/day/%d/input", baseUrl, year, day)
	cacheKey := fmt.Sprintf("%s/input-%d-%d-%d.txt", cacheDir, year, day, part)

	ps := ProblemSolver{cookieFilePath: cookieFilePath, inputUrl: inputUrl, baseUrl: baseUrl, part: part, cacheKey: cacheKey}
	ps.initHttpClient()
	return &ps
}

func (ps *ProblemSolver) Solve() string {
	if ps.part == 1 {
		return ps.SolvePart1()
	} else {
		return ps.SolvePart2()
	}
}

func (ps *ProblemSolver) SolvePart1() string {
	fmt.Println("\tPulling down the input file")
	input, err := ps.GetInput()
	check(err, "Failed to get input")

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
					lastDigit = firstDigit
				} else {
					lastDigit = r
				}
			}
		}

		// convert to two digit integer
		val, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		check(err, "Failed to convert to integer")

		sum += val
	}

	return strconv.Itoa(sum)
}

func (ps *ProblemSolver) SolvePart2() string {
	return fmt.Sprintf("Solving problem")
}

func (ps *ProblemSolver) GetInput() (string, error) {
	// Check if cached data exists.
	if data, err := os.ReadFile(ps.cacheKey); err == nil {
		fmt.Println("\tUsing cached input")
		return string(data), nil
	}

	// If not, get data from the website.
	result, err := ps.GetInputFromWebsite()
	if err != nil {
		return "", err
	}

	// Cache the data.
	os.WriteFile(ps.cacheKey, []byte(result), 0644)

	return result, nil
}

func (ps *ProblemSolver) GetInputFromWebsite() (string, error) {
	resp, err := ps.httpClient.Get(ps.inputUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (ps *ProblemSolver) initHttpClient() {
	// Create an HTTP client with a cookie jar
	jar, err := cookiejar.New(nil)
	check(err, "Failed to create cookie jar")

	// Read cookies from a file
	cookies, err := readCookies(ps.cookieFilePath)
	check(err, "Failed to read cookies")

	// Set cookies for the client
	baseUrl, err := url.Parse(ps.baseUrl)
	check(err, "Failed to parse base url")
	jar.SetCookies(baseUrl, cookies)

	ps.httpClient = &http.Client{
		Jar: jar,
	}
}

// readCookies reads cookies from a file and returns them as a slice of *http.Cookie.
func readCookies(filepath string) ([]*http.Cookie, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var cookies []*http.Cookie

	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) == 7 {
			cookies = append(cookies, &http.Cookie{
				Name:  parts[5],
				Value: parts[6],
				// Add other cookie attributes as needed, e.g., Domain, Path, etc.
			})
		}
	}

	return cookies, nil
}

func check(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		panic(e)
	}
}
