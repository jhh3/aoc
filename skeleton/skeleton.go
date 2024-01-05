package skeleton

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

//go:embed templates/*.tmpl
var fs embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	validateInput(day, year)

	// Load the templates
	ts, err := template.ParseFS(fs, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	// Make the directory
	mainFilename := filepath.Join(dirname(), "../", fmt.Sprintf("%d/day%02d/main.go", year, day))
	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	// Make each file
	makeFile("main.go", day, year, ts)
	makeFile("main_test.go", day, year, ts)
	makeFile("example_input.txt", day, year, ts)

	fmt.Printf("templates made for %d-day%d\n", year, day)
}

func makeFile(filename string, day, year int, tmpl *template.Template) {
	fullFn := filepath.Join(dirname(), "../", fmt.Sprintf("%d/day%02d/%s", year, day, filename))
	ensureNotOverwriting(fullFn)

	f, err := os.Create(fullFn)
	if err != nil {
		log.Fatalf("creating %s file: %v", filename, err)
	}

	tmpl.ExecuteTemplate(f, fmt.Sprintf("%s.tmpl", filename), nil)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}

func dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

func validateInput(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}
}
