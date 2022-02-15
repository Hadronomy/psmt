package build

import (
	"runtime/debug"
	"time"
)

var Version = "DEV"

var Date = "DEV\n" + time.Now().Format("02-01-2006 15:04:05")

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}
