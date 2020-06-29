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
//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT

//sys	mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sys	Pause() (err error)

func mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
	page := uintptr(offset / 4096)
	if offset != int64(page)*4096 {
		return 0, EINVAL
	}
	return mmap2(addr, length, prot, flags, fd, page)
}

type rlimit32 struct {
	Cur uint32
	Max uint32
}

//sysnb getrlimit(resource int, rlim *rlimit32) (err error) = SYS_GETRLIMIT

const rlimInf32 = ^uint32(0)
const rlimInf64 = ^uint64(0)

func Getrlimit(resource int, rlim *Rlimit) (err error) {
	err = prlimit(0, resource, nil, rlim)
	if err != ENOSYS {
		return err
	}

	rl := rlimit32{}
	err = getrlimit(resource, &rl)
	if err != nil {
		return
	}

	if rl.Cur == rlimInf32 {
		rlim.Cur = rlimInf64
	} else {
		rlim.Cur = uint64(rl.Cur)
	}

	if rl.Max == rlimInf32 {
		rlim.Max = rlimInf64
	} else {
		rlim.Max = uint64(rl.Max)
	}
	return
}

//sysnb setrlimit(resource int, rlim *rlimit32) (err error) = SYS_SETRLIMIT

func Setrlimit(resource int, rlim *Rlimit) (err error) {
	err = prlimit(0, resource, rlim, nil)
	if err != ENOSYS {
		return err
	}

	rl := rlimit32{}
	if rlim.Cur == rlimInf64 {
		rl.Cur = rlimInf32
	} else if rlim.Cur < uint64(rlimInf32) {
		rl.Cur = uint32(rlim.Cur)
	} else {
		return EINVAL
	}
	if rlim.Max == rlimInf64 {
		rl.Max = rlimInf32
	} else if rlim.Max < uint64(rlimInf32) {
		rl.Max = uint32(rlim.Max)
	} else {
		return EINVAL
	}

	return setrlimit(resource, &rl)
}

// Underlying system call writes to newoffset via pointer.
// Implemented in assembly to avoid allocation.
func seek(fd int, offset int64, whence int) (newoffset int64, err syscall.Errno)

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	newoffset, errno := seek(fd, offset, whence)
	if errno != 0 {
		return 0, errno
	}
	return newoffset, nil
}

// Vsyscalls on amd64.
//sysnb	Gettimeofday(tv *Timeval) (err error)
//sysnb	Time(t *Time_t) (tt Time_t, err error)

//sys	Utime(path string, buf *Utimbuf) (err error)

// On x86 Linux, all the socket calls go through an extra indirection,
// I think because the 5-register system call interface can't handle
// the 6-argument calls like sendto and recvfrom. Instead the
// arguments to the underlying system call are the number below
// and a pointer to an array of uintptr. We hide the pointer in the
// socketcall assembly to avoid allocation on every system call.

const (
	// see linux/net.h
	_SOCKET      = 1
	_BIND        = 2
	_CONNECT     = 3
	_LISTEN      = 4
	_ACCEPT      = 5
	_GETSOCKNAME = 6
	_GETPEERNAME = 7
	_SOCKETPAIR  = 8
	_SEND        = 9
	_RECV        = 10
	_SENDTO      = 11
	_RECVFROM    = 12
	_SHUTDOWN    = 13
	_SETSOCKOPT  = 14
	_GETSOCKOPT  = 15
	_SENDMSG     = 16
	_RECVMSG     = 17
	_ACCEPT4     = 18
	_RECVMMSG    = 19
	_SENDMMSG    = 20
)

func socketcall(call int, a0, a1, a2, a3, a4, a5 uintptr) (n int, err syscall.Errno)
func rawsocketcall(call int, a0, a1, a2, a3, a4, a5 uintptr) (n int, err syscall.Errno)

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	fd, e := socketcall(_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e != 0 {
		err = e
	}
	return
}

func accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) {
	fd, e := socketcall(_ACCEPT4, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), uintptr(flags), 0, 0)
	if e != 0 {
		err = e
	}
	return
}

func getsockname(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, e := rawsocketcall(_GETSOCKNAME, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e != 0 {
		err = e
	}
	return
}

func getpeername(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (err