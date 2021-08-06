BIN_D = $(PWD)/bin

install:
	@mkdir -p $(BIN_D)
	go mod vendor && go build -o bin/slacker slacker.go
	@echo "Add bin folder location to PATH env for use 'slacker' command by default"

version:
	@echo "slacker-cli 0.1"