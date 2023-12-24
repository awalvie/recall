NAME=recall
VERSION=0.0.1

.DEFAULT_GOAL := help

.PHONY: build
## build: Compile the packages.
build:
	@CGO_ENABLED=1 go build -o $(NAME) cmd/$(NAME).go

.PHONY: run
## run : Run the program
run: build
	@./$(NAME)

.PHONY: clean
## clean: Clean projects and previous builds
clean:
	@rm -rf $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: watch
## watch: Reload the app whenever the source changes
watch:
	@which reflex > /dev/null || (go install github.com/cespare/reflex@latest)
	reflex -s -r '\.go$$' make run

.PHONY: docker-build
## docker-build: Build docker image
docker-build:
	@docker build -t awalvie/$(NAME):$(VERSION) -f contrib/Dockerfile .

.PHONY: help
all: help
## help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo

