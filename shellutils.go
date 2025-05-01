// "unix shell" like utilities

package shellutils

import (
	"fmt"
	"io/fs"
	"os"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(os.Stderr, "runtime.GOOS == \"windows\"\n")
		os.Exit(1)
	}
}

// Ls
// List information about the FILEs.
func Ls(path string) (dirEns []fs.DirEntry, err error) {
	dirFS := os.DirFS(path)              // fs.FS
	dirEns, err = fs.ReadDir(dirFS, ".") // []fs.DirEntry
	if err != nil {
		return nil, fmt.Errorf("fs.ReadDir: %v\n", err)
	}

	return dirEns, nil
}
