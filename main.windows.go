//go:build windows
// +build windows

// The syscall package provides a common set of interfaces to system calls on
// Windows and Linux.  All versions offer a set of functions with the same signatures.
// The constant OSName contains the string "windows" or "linux", the same name as the
// build tag for the target system.  The Windows version of the functions all return
// a syscall.EWINDOWS ("not supported by windows") error when called.
//
// This allows source code that is intended to run under Linux and uses system calls to
// at least compile under Windows.  The result can also be run in a limited fashion
// under Windows as long as it uses the OSName constant to avoid calling the syscall
// functions (which would break if called in that environment).  This allows system
// testing under Windows of other parts of the solution that don't use this package.
package sys

import (
	"syscall"
)

// OSName contains the name of the target operating system.  It's the same value as the build tag
// for that system ("windows", "linux" or whatever).
const OSName = "windows"

// Getuid gets the current user ID.  The Windows version always returns -1
// and the error "not implemented".
func Getuid() (int, error) {
	return -1, syscall.EWINDOWS
}

// Setuid switches the effective user to the user with the given user ID.  The
// Windows version always returns a "not impleented" error.
func Setuid(targetID int) error {
	return syscall.EWINDOWS
}
