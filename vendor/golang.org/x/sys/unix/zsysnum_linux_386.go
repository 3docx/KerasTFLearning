// linux/mksysnum.pl -Wall -Werror -static -I/tmp/include -m32 /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,linux

package unix

const (
	SYS_RESTART_SYSCALL        = 0
	SYS_EXIT                   = 1
	SYS_FORK                   = 2
	SYS_READ                   = 3
	SYS_WRITE                  = 4
	SYS_OPEN                   = 5
	SYS_CLOSE                  = 6
	SYS_WAITPID                = 7
	SYS_CREAT                  = 8
	SYS_LINK                   = 9
	SYS_UNLINK                 = 10
	SYS_EXECVE                 = 11
	SYS_CHDIR                  = 12
	SYS_TIME                   = 13
	SYS_MKNOD                  = 14
	SYS_CHMOD                  = 15
	SYS_LCHOWN                 = 16
	SYS_BREAK                  = 17
	SYS_OLDSTAT                = 18
	SYS_LSEEK                  = 19
	SYS_GETPID                 = 20
	SYS_MOUNT                  = 21
	SYS_UMOUNT                 = 22
	SYS_SETUID                 = 23
	SYS_GETUID                 = 24
	SYS_STIME                  = 25
	SYS_PTRACE                 = 26
	SYS_ALARM                  = 27
	SYS_OLDFSTAT               = 28
	SYS_PAUSE                  = 29
	SYS_UTIME                  = 30
	SYS_STTY                   = 31
	SYS_GTTY                   = 32
	SYS_ACCESS                 = 33
	SYS_NICE                   = 34
	SYS_FTIME                  = 35
	SYS_SYNC                   = 36
	SYS_KILL                   = 37
	SYS_RENAME                 = 38
	SYS_MKDIR                  = 39
	SYS_RMDIR                  = 40
	SYS_DUP                    = 41
	SYS_PIPE                   = 42
	SYS_TIMES                  = 43
	SYS_PROF                   = 44
	SYS_BRK                    = 45
	SYS_SETGID                 = 46
	SYS_GETGID                 = 47
	SYS_SIGNAL                 = 48
	SYS_GETEUID                = 49
	SYS_GETEGID                = 50
	SYS_ACCT                   = 51
	SYS_UMOUNT2                = 52
	SYS_LOCK                   = 53
	SYS_IOCTL                  = 54
	SYS_FCNTL                  = 55
	SYS_MPX                    = 56
	SYS_SETPGID                = 57
	SYS_ULIMIT                 = 58
	SYS_OLDOLDUNAME            = 59
	SYS_UMASK                  = 60
	SYS_CHROOT                 = 61
	SYS_USTAT                  = 62
	SYS_DUP2                   = 63
	SYS_GETPPID                = 64
	SYS_GETPGRP                = 65
	SYS_SETSID                 = 66
	SYS_SIGACTION              = 67
	SYS_SGETMASK               = 68
	SYS_SSETMASK               = 69
	SYS_SETREUID               = 70
	SYS_SETREGID               = 71
	SYS_SIGSUSPEND             = 72
	SYS_SIGPENDING             = 73
	SYS_SETHOSTNAME            = 74
	SYS_SETRLIMIT              = 75
	SYS_GETRLIMIT              = 76
	SYS_GETRUSAGE              = 77
	SYS_GETTIMEOFDAY           = 78
	SYS_SETTIMEOFDAY           = 79
	SYS_GETGROUPS              = 80
	SYS_SETGROUPS              = 81
	SYS_SELECT                 = 82
	SYS_SYMLINK                = 83
	SYS_OLDLSTAT               = 84
	SYS_READLINK               = 85
	SYS_USELIB                 = 86
	SYS_SWAPON                 = 87
	SYS_REBOOT                 = 88
	SYS_READDIR                = 89
	SYS_MMAP                   = 90
	SYS_MUNMAP                 = 91
	SYS_TRUNCATE               = 92
	SYS_FTRUNCATE              = 93
	SYS_FCHMOD                 = 94
	SYS_FCHOWN                 = 95
	SYS_GETPRIORITY            = 96
	SYS_SETPRIORITY            = 97
	SYS_PROFIL                 = 98
	SYS_STATFS                 = 99
	SYS_FSTATFS                = 100
	SYS_IOPERM                 = 101
	SYS_SOCKETCALL             = 102
	SYS_SYSLOG                 = 103
	SYS_SETITIMER              = 104
	SYS_GETITIMER              = 105
	SYS_STAT                   = 106
	SYS_LSTAT                  = 107
	SYS_FSTAT                  = 108
	SYS_OLDUNAME               = 109
	SYS_IOPL                   = 110
	SYS_VHANGUP                = 111
	SYS_IDLE                   = 112
	SYS_VM86OLD                = 113
	SYS_WAIT4                  = 114
	SYS_SWAPOFF                = 115
	SYS_SYSINFO                = 116
	SYS_IPC                    = 117
	SYS_FSYNC                  = 118
	SYS_SIGRETURN              = 119
	SYS_CLONE                  = 120
	SYS_SETDOMAINNAME          = 121
	SYS_UNAME                  = 122
	SYS_MODIFY_LDT             = 123
	SYS_ADJTIMEX               = 124
	SYS_MPROTECT               = 125
	SYS_SIGPROCMASK            = 126
	SYS_CREATE_MODULE          = 127
	SYS_INIT_MODULE            = 128
	SYS_DELETE_MODULE          = 129
	SYS_GET_KERNEL_SYMS        = 130
	SYS_QUOTACTL               = 131
	SYS_GETPGID                = 132
	SYS_FCHDIR                 = 133
	SYS_BDFLUSH                = 134
	SYS_SYSFS                  = 135
	SYS_PERSONALITY            = 136
	SYS_AFS_SYSCALL            = 137
	SYS_SETFSUID               = 138
	SYS_SETFSGID               = 139
	SYS__LLSEEK                = 140
	SYS_GETDENTS               = 141
	SYS__NEWSELECT             = 142
	SYS_FLOCK                  = 143
	SYS_MSYNC                  = 144
	SYS_READV                  = 145
	SYS_WRITEV                 = 146
	SYS_GETSID                 = 147
	SYS_FDATASYNC              = 148
	SYS__SYSCTL                = 149
	SYS_MLOCK                  = 150
	SYS_MUNLOCK                = 151
	SYS_MLOCKALL               = 152
	SYS_MUNLOCKALL             = 153
	SYS_SCHED_SETPARAM         = 154
	SYS_SCHED_GETPARAM         = 155
	SYS_SCHED_SETSCHEDULER     = 156
	SYS_SCHED_GETSCHEDULER     = 157
	SYS_SCHED_YIELD            = 158
	SYS_SCHED_GET_PRIORITY_MAX = 159
	SYS_SCHED_GET_PRIORITY_MIN = 160
	SYS_SCHED_RR_GET_INTERVAL  = 161
	SYS_NANOSLEEP              = 162
	SYS_MREMAP                 = 163
	SYS_SETRESUID              = 164
	SYS_GETRESUID              = 165
	SYS_VM86                   = 166
	SYS_QUERY_MODULE           = 167
	SYS_POLL                   = 168
	SYS_NFSSERVCTL             = 169
	SYS_SETRESGID              = 170
	SYS_GETRESGID              = 171
	SYS_PRCTL                  = 172
	SYS_RT_SIGRETURN           = 173
	SYS_RT_SIGACTION           = 174
	SYS_RT_SIGPROCMASK         = 175
	SYS_RT_SIGPENDING          = 176
	SYS_RT_SIGTIMEDWAIT        = 177
	SYS_RT_SIGQUEUEINFO        = 178
	SYS_RT_SIGSUSPEND          = 179
	SYS_PREAD64                = 180
	SYS_PWRITE64               = 181
	SYS_CHOWN                  = 182
	SYS_GETCWD                 = 183
	SYS_CAPGET     