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

---