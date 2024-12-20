package common

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

type ProblemRunner interface {
	Run()
}

type baseProblemRunnerImpl struct {
	flags  *ProblemSolverFlags
	Solver ProblemSolver
	input  string
}

func NewProblemRunner(flags *ProblemSolverFlags, solver ProblemSolver, input string) ProblemRunner {
	return &baseProblemRunnerImpl{
		flags:  flags,
		Solver: solver,
		input:  input,
	}
}

func RunFromSolver(solver ProblemSolver, input string) {
	flags := MustParseSolverFlags(os.Args[1:], true)
	runner := NewProblemRunner(
		flags,
		solver,
		input,
	)
	runner.Run()
}

func (pr *baseProblemRunnerImpl) Run() {
	fmt.Println("Solving...")
	var result string
	if pr.flags.Part == 1 {
		result = pr.Solver.SolvePart1(pr.input)
	} else {
		result = pr.Solver.SolvePart2(pr.input)
	}
	fmt.Println("Result:", result)
}

type ProblemSolver interface {
	SolvePart1(input string) string
	SolvePart2(input string) string
}

type ProblemReader interface {
	GetInput() (string, error)
	MustGetInput() string
}

type baseProblemReaderImpl struct {
	flags      *InputGetterFlags
	httpClient *http.Client
}

func NewProblemReader(flags *InputGetterFlags) ProblemReader {
	pr := baseProblemReaderImpl{flags: flags}
	pr.initHttpClient()
	return &pr
}

func (pr *baseProblemReaderImpl) MustGetInput() string {
	result, err := pr.GetInput()
	CheckErr(err, "Failed to get input")
	return result
}

func (pr *baseProblemReaderImpl) GetInput() (string, error) {
	// Check if cached data exists.
	if data, err := os.ReadFile(pr.flags.CacheKey()); err == nil {
		fmt.Println("\tUsing cached input")
		return string(data), nil
	}

	// If not, get data from the website.
	result, err := pr.GetInputFromWebsite()
	if err != nil {
		return "", err
	}

	// Cache the data.
	os.WriteFile(pr.flags.CacheKey(), []byte(result), 0644)

	return result, nil
}

func (pr *baseProblemReaderImpl) GetInputFromWebsite() (string, error) {
	resp, err := pr.httpClient.Get(pr.flags.InputUrl())
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

func (pr *baseProblemReaderImpl) initHttpClient() {
	// Create an HTTP client with a cookie jar
	jar, err := cookiejar.New(nil)
	CheckErr(err, "Failed to create cookie jar")

	// Read cookies from a file
	cookies, err := readCookies(pr.flags.CookieFilePath)
	CheckErr(err, "Failed to read cookies")

	// Set cookies for the client
	baseUrl, err := url.Parse(pr.flags.BaseUrl)
	CheckErr(err, "Failed to parse base url")
	jar.SetCookies(baseUrl, cookies)

	pr.httpClient = &http.Client{
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
