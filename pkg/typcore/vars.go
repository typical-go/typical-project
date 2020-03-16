package typcore

const (
	// Version of Typical-Go
	Version = "0.9.41"
)

var (
	// DefaultProjectPackage is default value for ProjectPackage
	DefaultProjectPackage = "" // NOTE: supply by ldflags

	// DefaultTempFolder is default value for temp folder location
	DefaultTempFolder = ".typical-tmp"

	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"
)
