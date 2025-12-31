package portablesyscall

import (
	"syscall"
	"testing"
)

// TestGetuid checks Getuid
func TestGetuid(t *testing.T) {

	uid, err := Getuid()

	if OSName == "windows" {
		if err == nil {
			t.Error("expected an error")
		}
		if err != syscall.EWINDOWS {
			t.Errorf("want syscall.EWINDOWS (%v) got %v", syscall.EWINDOWS, err)
		}
	} else {
		if err != nil {
			t.Error(err)
		}

		if uid < 1 {
			t.Errorf("want a uid which is greater than 1, got %d", uid)
		}
	}
}

// TestSetuid checks Setuid under Windows
func TestSetuid(t *testing.T) {
	err := Setuid(1)

	if OSName == "windows" {
		if err != syscall.EWINDOWS {
			t.Errorf("want syscall.EWINDOWS (%v) got %v", syscall.EWINDOWS, err)
		}
	}
}
