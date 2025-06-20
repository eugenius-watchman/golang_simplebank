package db

import (
	"fmt"
	"runtime"
)

func VerifyGoVersion(minVersion string) {
	if runtime.Version() < minVersion {
		panic(fmt.Sprintf(
			"Requires Go %s or later (using %s)",
			minVersion,
			runtime.Version(),
		))
	}
}