EXECUTABLE=gc
WINDOWS=bin/$(EXECUTABLE).exe
LINUX=bin/$(EXECUTABLE)_linux_amd64
DARWIN=bin/$(EXECUTABLE)

.PHONY: all test clean

all: test build ## Build and run tests

test:
	cd gc && go test -cover -v ./...

build: clean windows linux darwin ## Build binaries

windows: clean $(WINDOWS) ## Build for Windows

linux: clean $(LINUX) ## Build for Linux

darwin: clean $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	cd gc && env GOOS=windows GOARCH=amd64 go build -v -o ../$(WINDOWS)

$(LINUX):
	cd gc && env GOOS=linux GOARCH=amd64 go build -v -o ../$(LINUX)

$(DARWIN):
	cd gc && env GOOS=darwin GOARCH=amd64 go build -v -o ../$(DARWIN)

clean:
	rm -rf bin/