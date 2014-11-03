export GOPATH := ${PWD}/.gopath

all: build

build: clean
		@echo "Building binary..."
		scripts/build.sh build

clean:
		@echo "Cleaning up..."
		@rm -rf bin

deps:
		@echo "Adding dependencies..."
		@scripts/deps.sh
