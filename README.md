#### Go utilities similar to "Unix Shell" utilities.

- ls — List directory contents.

---

```go
package main

import(
	"os"
	"fmt"
	"github.com/aspieunixgeek/shellutils"
)

func main() {
	rootDir, err := shellutils.Ls("/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "shellutils.Ls: %v\n", err)
		os.Exit(1)
	}

	for _, i := range rootDir {
		fmt.Println(i)
	}

}
```

Output:
```bash

L bin
d boot/
d cdrom/
d dev/
d etc/
d home/
L lib
L lib32
L lib64
L libx32
d lost+found/
d media/
d mnt/
d opt/
d proc/
d root/
d run/
L sbin
d snap/
d srv/
d sys/
d tmp/
d usr/
d var/
```


Unit tests output:
```bash
$ go test -v
=== RUN   TestLs
    shellutils_test.go:42: L bin
    shellutils_test.go:42: d boot/
    shellutils_test.go:42: d cdrom/
    shellutils_test.go:42: d dev/
    shellutils_test.go:42: d etc/
    shellutils_test.go:42: d home/
    shellutils_test.go:42: L lib
    shellutils_test.go:42: L lib32
    shellutils_test.go:42: L lib64
    shellutils_test.go:42: L libx32
    shellutils_test.go:42: d lost+found/
    shellutils_test.go:42: d media/
    shellutils_test.go:42: d mnt/
    shellutils_test.go:42: d opt/
    shellutils_test.go:42: d proc/
    shellutils_test.go:42: d root/
    shellutils_test.go:42: d run/
    shellutils_test.go:42: L sbin
    shellutils_test.go:42: d snap/
    shellutils_test.go:42: d srv/
    shellutils_test.go:42: d sys/
    shellutils_test.go:42: d tmp/
    shellutils_test.go:42: d usr/
    shellutils_test.go:42: d var/
--- PASS: TestLs (0.00s)
PASS
ok      github.com/aspieunixgeek/shellutils     0.002s

```

---