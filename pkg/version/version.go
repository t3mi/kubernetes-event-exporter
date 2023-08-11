package version

import (
	"runtime"
)

// Build information. Populated at build-time.
var (
	Version   string
	Revision  string
	Branch    string
	GoVersion = runtime.Version()
	GoOS      = runtime.GOOS
	GoArch    = runtime.GOARCH
)
