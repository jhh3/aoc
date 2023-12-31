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
	go run $(YEAR)/day$(DAY)/main.go --cookie ~/Downloads/cookies.txt --part $(PART)


.PHONY: test run help
