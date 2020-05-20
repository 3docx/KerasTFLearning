// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gccgo

package unix

import "syscall"

// We can't use the gc-syntax .s files for gccgo. On the plus side
// much of the functionality can be written directly in Go.

//extern gccgoRealSyscallNoError
