BUILD_DATE="$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')"
GIT_COMMIT="$(shell git rev-parse --short HEAD)"
VERSION="0.0.1"

build:
	go build -o bin/tcomp -ldflags="-X 'github.com/bmehdi777/tcomp/internal/pkg/version.buildDate=$(BUILD_DATE)' -X 'github.com/bmehdi777/tcomp/internal/pkg/version.gitCommit=$(GIT_COMMIT)' -X 'github.com/bmehdi777/tcomp/internal/pkg/version.version=$(VERSION)'" cmd/tcomp/main.go

clean:
	rm bin/*
