// mksysnum_dragonfly.pl
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build amd64,dragonfly

package unix

const (
	// SYS_NOSYS = 0;  // { int nosys(void); } syscall nosys_args int
	SYS_EXIT          = 1   // { void exit(int rval); }
	SYS_FORK          = 2   // { int fork(void); }
	SYS_READ          = 3   // { ssize_t read(int fd, void *buf, size_t nbyte); }
	SYS_WRITE         = 4   // { ssize_t write(int fd, const void *buf, size_t nbyte); }
	SYS_OPEN          = 5   // { int open(char *path, int flags, int mode); }
	SYS_CLOSE         = 6   // { int close(int fd); }
	SYS_WAIT4         = 7   // { int wait4(int pid, int *status, int options, \
	SYS_LINK          = 9   // { int link(char *path, char *link); }
	SYS_UNLINK        = 10  // { int unlink(char *path); }
	SYS_CHDIR         = 12  // { int chdir(char *path); }
	SYS_FCHDIR        = 13  // { int fchdir(int fd); }
	SYS_MKNOD         = 14  // { int mknod(char *path, int mode, int dev); }
	SYS_CHMOD         = 15  // { int chmod(char *path, int mode); }
	SYS_CHOWN         = 16  // { int chown(char *path, int uid, int gid); }
	SYS_OBREAK        = 17  // { int obreak(char *nsize); } break obreak_args int
	SYS_GETFSSTAT     = 18  // { int getfsstat(struct statfs *buf, long bufsize, \
	SYS_GETPID        = 20  // { pid_t getpid(void); }
	SYS_MOUNT         = 21  // { int mount(char *type, char *path, int flags, \
	SYS_UNMOUNT       = 22  // { int unmount(char *path, int flags); }
	SYS_SETUID        = 23  // { int setuid(uid_t uid); }
	SYS_GETUID        = 24  // { uid_t getuid(void); }
	SYS_GETEUID       = 25  // { uid_t geteuid(void); }
	SYS_PTRACE        = 26  // { int ptrace(int req, pid_t pid, caddr_t addr, \
	SYS_RECVMSG       = 27  // { int recvmsg(int s, struct msghdr *msg, int flags); }
	SYS_SENDMSG       = 28  // { int sendmsg(int s, caddr_t msg, int flags); }
	SYS_RECVFROM      = 29  // { int recvfrom(int s, caddr_t buf, size_t len, \
	SYS_ACCEPT        = 30  // { int accept(int s, caddr_t name, int *anamelen); }
	SYS_GETPEERNAME   = 31  // { int getpeername(int fdes, caddr_t asa, int *alen); }
	SYS_GETSOCKNAME   = 32  // { int getsockname(int fdes, caddr_t asa, int *alen); }
	SYS_ACCESS        = 33  // { int access(char *path, int flags); }
	SYS_CHFLAGS       = 34  // { int chflags(char *path, int flags); }
	SYS_FCHFLAGS      = 35  // { int fchflags(int fd, int flags); }
	SYS_SYNC          = 36  // { int sync(void); }
	SYS_KILL          = 37  // { int kill(int pid, int signum); }
	SYS_GETPPID       = 39  // { pid_t getppid(void); }
	SYS_DUP           = 41  // { int dup(int fd); }
	SYS_PIPE          = 42  // { int pipe(void); }
	SYS_GETEGID       = 43  // { gid_t getegid(void); }
	SYS_PROFIL        = 44  // { int profil(caddr_t samples, size_t size, \
	SYS_KTRACE        = 45  // { int ktrace(const char *fname, int ops, int facs, \
	SYS_GETGID        = 47  // { gid_t getgid(void); }
	SYS_GETLOGIN      = 49  // { int getlogin(char *namebuf, u_int namelen); }
	SYS_SETLOGIN      = 50  // { int setlogin(char *namebuf); }
	SYS_ACCT          = 51  // { int acct(char *path); }
	SYS_SIGALTSTACK   = 53  // { int sigaltstack(stack_t *ss, stack_t *oss); }
	SYS_IOCTL         = 54  // { int ioctl(int fd, u_long com, caddr_t data); }
	SYS_REBOOT        = 55  // { int reboot(int opt); }
	SYS_REVOKE        = 56  // { int revoke(char *path); }
	SYS_SYMLINK       = 57  // { int symlink(char *path, char *link); }
	SYS_READLINK      = 58  // { int readlink(char *path, char *buf, int count); }
	SYS_EXECVE        = 59  // { int execve(char *fname, char **argv, char **envv); }
	SYS_UMASK         = 60  // { int umask(int newmask); } umask umask_args int
	SYS_CHROOT        = 61  // { int chroot(char *path); }
	SYS_MSYNC         = 65  // { int msync(void *addr, size_t len, int flags); }
	SYS_VFORK         = 66  // { pid_t vfork(void); }
	SYS_SBRK          = 69  // { int sbrk(int incr); }
	SYS_SSTK          = 70  // { int sstk(int incr); }
	SYS_MUNMAP        = 73  // { int munmap(void *addr, size_t len); }
	SYS_MPROTECT      = 74  // { int mprotect(void *addr, size_t len, int prot); }
	SYS_MADVISE       = 75  // { int madvise(void *addr, size_t len, int behav); }
	SYS_MINCORE       = 78  // { int mincore(const void *addr, size_t len, \
	SYS_GETGROUPS     = 79  // { int getgroups(u_int gidsetsize, gid_t *gidset); }
	SYS_SETGROUPS     = 80  // { int setgroups(u_int gidsetsize, gid_t *gidset); }
	SYS_GETPGRP  