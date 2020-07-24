// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package unix

import (
	"bytes"
	"runtime"
	"sort"
	"sync"
	"syscall"
	"unsafe"
)

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

const (
	darwin64Bit    = runtime.GOOS == "darwin" && sizeofPtr == 8
	dragonfly64Bit = runtime.GOOS == "dragonfly" && sizeofPtr == 8
	netbsd32Bit    = runtime.GOOS == "netbsd" && sizeofPtr == 4
	solaris64Bit   = runtime.GOOS == "solaris" && sizeofPtr == 8
)

// Do the interface allocations only once for common
// Errno values.
var (
	errEAGAIN error = syscall.EAGAIN
	errEINVAL error = syscall.EINVAL
	errENOENT error = syscall.ENOENT
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case EAGAIN:
		return errEAGAIN
	case EINVAL:
		return errEINVAL
	case ENOENT:
		return errENOENT
	}
	return e
}

// ErrnoName returns the error name for error number e.
func ErrnoName(e syscall.Errno) string {
	i := sort.Search(len(errorList), func(i int) bool {
		return errorList[i].num >= e
	})
	if i < len(errorList) && errorList[i].num == e {
		return errorList[i].name
	}
	return ""
}

// SignalName returns the signal name for signal number s.
func SignalName(s syscall.Signal) string {
	i := sort.Search(len(signalList), func(i int) bool {
		return signalList[i].num >= s
	})
	if i < len(signalList) && signalList[i].num == s {
		return signalList[i].name
	}
	return ""
}

// clen returns the index of the first NULL byte in n or len(n) if n contains no NULL byte.
func clen(n []byte) int {
	i := bytes.IndexByte(n, 0)
	if i == -1 {
		i = len(n)
	}
	return i
}

// Mmap manager, for use by operating system-specific implementations.

type mmapper struct {
	sync.Mutex
	active map[*byte][]byte // active mappings; key is last byte in mapping
	mmap   func(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error)
	munmap func(addr uintptr, length uintptr) error
}

func (m *mmapper) Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
	if length <= 0 {
		return nil, EINVAL
	}

	// Map the requested memory.
	addr, errno := m.mmap(0, uintptr(length), prot, flags, fd, offset)
	if errno != nil {
		return nil, errno
	}

	// Slice memory layout
	var sl = struct {
		addr uintptr
		le