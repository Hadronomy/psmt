package build

import (
	"runtime/debug"
)

var Version = "DEV"

var Date = "Unknown"

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}
