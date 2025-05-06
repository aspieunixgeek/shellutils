// List information about the FILEs

package ls

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(os.Stderr, "runtime.GOOS == \"windows\"\n")
		os.Exit(1)
	}
}

// List
// Get information about the FILEs (the path directory),
func List(path string, flags []string) ([]byte, error) {
	lsPath, err := exec.LookPath("ls")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}

	if err != nil {
		return nil, fmt.Errorf("exec.LookPath: %v\n", err)
	}

	err = os.Chdir(path)
	if err != nil {
		return nil, fmt.Errorf("os.Chdir: %v\n", err)
	}

	out, err := exec.Command(lsPath, flags...).Output()
	if err != nil {
		return nil, fmt.Errorf("exec.Command: %v\n", err)
	}

	return out, nil
}
