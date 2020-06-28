// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(rsc): Rewrite all nn(SP) references into name+(nn-8)(FP)
// so that go vet can check that they are correct.

// +build 386,linux

package unix

import (
	"syscall"
	"unsafe"
)

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
}

//sysnb	pipe(p *[2]_C_int) (err error)

func Pipe(p []int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe(&pp)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}

//sysnb pipe2(p *[2]_C_int, flags int) (err error)

func Pipe2(p []int, flags int) (err error) {
	if len(p) != 2 {
		return EINVAL
	}
	var pp [2]_C_int
	err = pipe2(&pp, flags)
	p[0] = int(pp[0])
	p[1] = int(pp[1])
	return
}

// 64-bit file system and 32-bit uid calls
// (386 default is 32-bit file system and 16-bit uid).
//sys	Dup2(oldfd int, newfd int) (err error)
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64_64
//sys	Fchown(fd int, uid int, gid int) (err error) = SYS_FCHOWN32
//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnb	Getegid() (egid int) = SYS_GETEGID32
//sysnb	Geteuid() (euid int) = SYS_GETEUID32
//sysnb	Getgid() (gid int) = SYS_GETGID32
//sysnb	Getuid() (uid int) = SYS_GETUID32
//sysnb	InotifyInit() (fd int, err error)
//sys	Ioperm(from int, num int, on int) (err error)
//sys	Iopl(level int) (err error)
//sys	Lchown(path string, uid int, gid int) (err error) = SYS_LCHOWN32
//sys	Lstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
//sys	Pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	Pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64
//sys	Setfsgid(gid int) (err error) = SYS_SETFSGID32
//sys	Setfsuid(uid int) (err error) = SYS_SETFSUID32
//sysnb	Setregid(rgid int, egid int) (err error) = SYS_SETREGID32
//sysnb	Setresgid(rgid int, egid int, sgid int) (err error) = SYS_SETRESGID32
//sysnb	Setresuid(ruid int, euid int, suid int) (err error) = SYS_SETRESUID32
//sysnb	Setreuid(ruid int, euid int) (err error) = SYS_SETREUID32
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sys	Stat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error) = SYS_GETGROUPS32
//sysnb	setgroups(n int, list *_Gid_t) (err error) = SYS_SETGROUPS32
//sys	Se