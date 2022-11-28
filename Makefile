NAME=recall
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME) cmd/${NAME}/main.go

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
	reflex -s -r '\.go$$' make run

.PHONY: help
all: help
# help: show this help message
help: makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo

