SHELL := /bin/bash

# https://gist.github.com/prwhite/8168133
help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run tests, all or a single day using optional YEAR and DAY variables
	@ if [[ -n $$DAY && -n $$YEAR  ]]; then \
		go test -v ./... --run '^Test_y$(YEAR)d$(DAY)' ; \
	else \
		go test -v ./... ; \
	fi

run: ## Run a specific year, day, and part using DAY, YEAR and PART variables
	go run $(YEAR)/day$(DAY)/main.go --part $(PART)


skeleton: ## create solution template files files, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY) -year $(YEAR) ; \
	elif [[ -n $$DAY ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY); \
	else \
		go run scripts/cmd/skeleton/main.go; \
	fi

input: ## get input, requires $AOC_SESSION_COOKIE, optional: $DAY and $YEAR
	@ if [[ -n $$YEAR ]]; then \
		go run scripts/cmd/input/main.go -day $(DAY) -cookie $(COOKIE); \
	else \
		go run scripts/cmd/input/main.go -day $(DAY) -year $(YEAR) -cookie $(COOKIE); \
	fi


init-question: skeleton input ## create question template and get input files, optional: $DAY and $YEAR
	@ echo "Question $(YEAR) day $(DAY) initialized"


.PHONY: test run help skeleton init-question input
