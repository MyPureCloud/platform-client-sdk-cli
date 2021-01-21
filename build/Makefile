EXECUTABLE=gc
WINDOWS=$(EXECUTABLE).exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)

.PHONY: all test clean

all: test build ## Build and run tests

test:
	cd src && go test -cover -v ./...

build: windows linux darwin ## Build binaries

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	cd src && env GOOS=windows GOARCH=amd64 go build -v -o ../bin/$(WINDOWS)

$(LINUX):
	cd src && env GOOS=linux GOARCH=amd64 go build -v -o ../bin/$(LINUX)

$(DARWIN):
	cd src && env GOOS=darwin GOARCH=amd64 go build -v -o ../bin/$(DARWIN)

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)