package ls

import (
	"bytes"
	"testing"
)

// TestList
// ls with -a flag
// print entries starting with "." and ".."
func TestList(t *testing.T) {
	dotCnt := 0

	rootDir, err := List("/tmp", []string{"-a"})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	o := bytes.Split(rootDir, []byte("\n"))
	for _, f := range o {
		if len(f) == 1 && bytes.Equal(f, []byte(".")) {
			dotCnt++
		}
		if len(f) == 2 && bytes.Equal(f, []byte("..")) {
			dotCnt++
		}
	}
	if dotCnt != 2 { // "." + ".." = 2
		t.Fatalf("dotCnt = %d is not equal 2", dotCnt)
	}

	t.Logf("dotCnt: %v\n", dotCnt)
}
