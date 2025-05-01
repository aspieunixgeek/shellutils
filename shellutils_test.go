package shellutils

import (
	"errors"
	"os"
	"testing"
)

var rootDirs = map[string]bool{
	"etc":  false,
	"boot": false,
	"home": false,
	"var":  false,
	"mnt":  false,
	"opt":  false,
	"usr":  false,
}

func TestLs(t *testing.T) {
	list, err := Ls("/")
	if err != nil && !errors.Is(err, os.ErrPermission) {
		t.Fatalf("ls: %v\n", err)
	}

	for _, file := range list {
		if file.IsDir() {
			if _, ok := rootDirs[file.Name()]; ok {
				rootDirs[file.Name()] = true
			}
		}
	}

	for _, file := range list {
		if file.IsDir() {
			if isExist, ok := rootDirs[file.Name()]; ok && !isExist {
				t.Errorf("dir is not exist: %v\n", file.Name())
			}
		}
	}

	for _, de := range list {
		t.Logf("%v\n", de)
	}
}
