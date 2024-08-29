package version

import (
	"fmt"
	"runtime"
)

var (
	version   = "v0.0.0"
	gitCommit = ""
	buildDate = "1970-01-01T00:00:00Z"
)

type VersionInfo struct {
	Version   string `json:"version" yaml:"version"`
	GitCommit string `json:"gitCommit" yaml:"gitCommit"`
	BuildDate string `json:"buildDate" yaml:"buildDate"`
	GoVersion string `json:"goVersion" yaml:"goVersion"`
	Compiler  string `json:"compiler" yaml:"compiler"`
	Platform  string `json:"platform" yaml:"platform"`
}

func Get() *VersionInfo {
	return &VersionInfo{
		Version:   version,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
