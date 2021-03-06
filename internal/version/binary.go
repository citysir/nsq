package version

import (
	"fmt"
	"runtime"
)

const Binary = "0.3.6-alpha"

func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Binary, runtime.Version())
}
