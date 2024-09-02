package version

type VersionInfo struct {
	Version string
}

func Get() *VersionInfo {
	return &VersionInfo{
		Version: "0.0.1",
	}
}
