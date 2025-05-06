package uptime

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	now := time.Now()
	u, err := New()
	if err != nil {
		t.Errorf("New: %v\n", err)
	}

	if u.CurTime != time.Now().Format("15:04:05") {
		t.Fatalf("u.CurTime(%s) != time.Now(%s)", u.CurTime, now)
	}
	if len(u.Up) == 0 {
		t.Fatal("len(u.Up) == 0\n")
	}

	if len(u.LoadAver) != 3 {
		t.Fatal("len(u.LoadAver) != 3\n")
	}

	for i, la := range u.LoadAver {
		if la == 0.0 {
			t.Fatalf("u.LoadAver[%d] == 0.0\n", i)
		}
	}
	t.Logf("%v", u)
}
