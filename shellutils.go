// "unix shell" like utilities

package shellutils

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(os.Stderr, "runtime.GOOS == \"windows\"\n")
		os.Exit(1)
	}
}
