// +build windows
// +build !appengine

package isatty

import (
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	fileNameInfo uintptr = 2
	fileTypePipe         = 3
)

var (
	kernel32                         = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode              