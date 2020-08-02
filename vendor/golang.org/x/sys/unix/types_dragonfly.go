// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

/*
Input to cgo -godefs.  See README.md
*/

// +godefs map struct_in_addr [4]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

package unix

/*
#define KERNEL
#include <dirent.h>
#include <fcntl.h>
#include <poll.h>
#include <signal.h>
#include <termios.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/event.h>
#include <sys/mman.h>
#include <sys/mount.h>
#include <sys/param.h>
#include <sys/ptrace.h>
#include <sys/resource.h>
#include <sys/select.h>
#include <sys/signal.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <sys/types.h>
#include <sys/un.h>
#include <sys/utsname.h>
#include <sys/wait.h>
#include <net/bpf.h>
#include <net/if.h>
#include <net/if_dl.h>
#include <net/route.h>
#include <netinet/in.h>
#include <netinet/icmp6.h>
#include <netinet/tcp.h>

enum {
	sizeofPtr = sizeof(void*),
};

union sockaddr_all {
	struct sockaddr s1;	// this one gets used for fields
	struct sockaddr_in s2;	// these pad it out
	struct sockaddr_in6 s3;
	struct sockaddr_un s4;
	struct sockaddr_dl s5;
};

struct sockaddr_any {
	struct sockaddr addr;
	char pad[sizeof(union sockaddr_all) - sizeof(struct sockaddr)];
};

*/
import "C"

// Machine characteristics; for internal use.

const (
	sizeofPtr      = C.sizeofPtr
	sizeofShort    = C.sizeof_short
	sizeofInt      = C.sizeof_int
	sizeofLong     = C.sizeof_long
	sizeofLongLong = C.sizeof_longlong
)

// Basic types

type (
	_C_short     C.short
	_C_int       C.int
	_C_long      C.long
	_C_long_long C.longlong
)

// Time

type Timespec C.struct_timespec

type Timeval C.struct_timeval

// Processes

type Rusage C.struct_rusage

type Rlimit C.struct_rlimit

type _Gid_t C.gid_t

// Files

const ( // Directory mode bits
	S_IFMT   = C.S_IFMT
	S_IFIFO  = C.S_IFIFO
	S_IFCHR  = C.S_IFCHR
	S_IFDIR  = C.S_IFDIR
	S_IFBLK  = C.S_IFBLK
	S_IFREG  = C.S_IFREG
	S_IFLNK  = C.S_IFLNK
	S_IFSOCK = C.S_IFSOCK
	S_ISUID  = C.S_ISUID
	S_ISGID  = C.S_ISGID
	S_ISVTX  = C.S_ISVTX
	S_IRUSR  = C.S_IRUSR
	S_IWUSR  = C.S_IWUSR
	S_IXUSR  = C.S_IXUSR
)

type Stat_t C.struct_stat

type Statfs_t C.struct_statfs

type Flock_t C.struct_flock

type