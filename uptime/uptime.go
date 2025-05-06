// Go-uptime returns a named structure, containing the current time, system uptime, and load average

package uptime

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	CurTime int = 0
	User        = 1
	Hours       = 2
	Min         = 3
)

type Uptime struct {
	CurTime  string    // current time
	Up       string    // system’s uptime
	User     string    // number of logged-in users
	Label    string    // load average label
	LoadAver []float64 // current load average
}

func (u Uptime) String() string {
	s1 := strconv.FormatFloat(u.LoadAver[0], 'f', -1, 64)
	s2 := strconv.FormatFloat(u.LoadAver[1], 'f', -1, 64)
	s3 := strconv.FormatFloat(u.LoadAver[2], 'f', -1, 64)
	return " " + u.CurTime + " " + u.Up + u.User + " " + u.Label + " " + s1 + ", " + s2 + ", " + s3
}

func New() (u *Uptime, err error) {
	path, err := exec.LookPath("uptime")
	if err != nil {
		err = fmt.Errorf("exec.LookPath: %v", err)
		return
	}

	out, err := exec.Command(path).Output()
	if err != nil {
		err = fmt.Errorf("exec.Command: %v", err)
		return
	}

	ss := strings.Split(string(out), ",")
	timeUptime := strings.TrimSpace(ss[0])

	x := strings.Split(timeUptime, " ")

	curTime, err := time.Parse("15:04:05", x[CurTime])
	if err != nil {
		fmt.Printf("time.Parse: %v\n", err)
		os.Exit(1)
	}

	trd := strings.TrimSpace(ss[2])
	v := strings.Split(trd, " ")

	la1 := v[len(v)-1]              // get load average 1
	la2 := strings.TrimSpace(ss[3]) // get load average 2
	la3 := strings.TrimSpace(ss[4]) // get load average 3

	loadAver1, err := strconv.ParseFloat(la1, 64)
	if err != nil {
		return nil, fmt.Errorf("loadAverage1: %v", err)
	}

	loadAver2, err := strconv.ParseFloat(la2, 64)
	if err != nil {
		return nil, fmt.Errorf("loadAverage2: %v", err)
	}

	loadAver3, err := strconv.ParseFloat(la3, 64)
	if err != nil {
		return nil, fmt.Errorf("loadAverage3: %v", err)
	}

	u = &Uptime{}
	u.CurTime = curTime.Format("15:04:05")
	u.Up = "up " + x[Hours] + " " + x[Min] + ","
	u.User = strings.TrimSpace(ss[User] + ",")
	u.Label = "load average:"
	u.LoadAver = append(u.LoadAver, loadAver1, loadAver2, loadAver3)
	return u, nil
}
