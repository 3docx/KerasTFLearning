// mkerrors.sh -Wall -Werror -static -I/tmp/include -fsigned-char
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build s390x,linux

// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -Wall -Werror -static -I/tmp/include -fsigned-char _const.go

package unix

import "syscall"

const (
	AAFS_MAGIC                           = 0x5a3c69f0
	ADFS_SUPER_MAGIC                     = 0xadf5
	AFFS_SUPER_MAGIC                     = 0xadff
	AFS_FS_MAGIC                         = 0x6b414653
	AFS_SUPER_MAGIC                      = 0x5346414f
	AF_ALG                               = 0x26
	AF_APPLETALK                         = 0x5
	AF_ASH                               = 0x12
	AF_ATMPVC                            = 0x8
	AF_ATMSVC                            = 0x14
	AF_AX25                              = 0x3
	AF_BLUETOOTH                         = 0x1f
	AF_BRIDGE                            = 0x7
	AF_CAIF                              = 0x25
	AF_CAN                               = 0x1d
	AF_DECnet                            = 0xc
	AF_ECONET                            = 0x13
	AF_FILE                              = 0x1
	AF_IB                                = 0x1b
	AF_IEEE802154                        = 0x24
	AF_INET                              = 0x2
	AF_INET6                             = 0xa
	AF_IPX                               = 0x4
	AF_IRDA                              = 0x17
	AF_ISDN                              = 0x22
	AF_IUCV                              = 0x20
	AF_KCM                               = 0x29
	AF_KEY                               = 0xf
	AF_LLC                               = 0x1a
	AF_LOCAL                             = 0x1
	AF_MAX                               = 0x2c
	AF_MPLS                              = 0x1c
	AF_NETBEUI                           = 0xd
	AF_NETLINK                           = 0x10
	AF_NETROM                            = 0x6
	AF_NFC                               = 0x27
	AF_PACKET                            = 0x11
	AF_PHONET                            = 0x23
	AF_PPPOX                             = 0x18
	AF_QIPCRTR                           = 0x2a
	AF_RDS                               = 0x15
	AF_ROSE                              = 0xb
	AF_ROUTE                             = 0x10
	AF_RXRPC                             = 0x21
	AF_SECURITY                          = 0xe
	AF_SMC                               = 0x2b
	AF_SNA                               = 0x16
	AF_TIPC                              = 0x1e
	AF_UNIX                              = 0x1
	AF_UNSPEC                            = 0x0
	AF_VSOCK                             = 0x28
	AF_WANPIPE                           = 0x19
	AF_X25                               = 0x9
	ALG_OP_DECRYPT                       = 0x0
	ALG_OP_ENCRYPT                       = 0x1
	ALG_SET_AEAD_ASSOCLEN                = 0x4
	ALG_SET_AEAD_AUTHSIZE                = 0x5
	ALG_SET_IV                           = 0x2
	ALG_SET_KEY                          = 0x1
	ALG_SET_OP                           = 0x3
	ANON_INODE_FS_MAGIC                  = 0x9041934
	ARPHRD_6LOWPAN                       = 0x339
	ARPHRD_ADAPT                         = 0x108
	ARPHRD_APPLETLK                      = 0x8
	ARPHRD_ARCNET                        = 0x7
	ARPHRD_ASH                           = 0x30d
	ARPHRD_ATM                           = 0x13
	ARPHRD_AX25                          = 0x3
	ARPHRD_BIF                           = 0x307
	ARPHRD_CAIF                          = 0x336
	ARPHRD_CAN                           = 0x118
	ARPHRD_CHAOS                         = 0x5
	ARPHRD_CISCO                         = 0x201
	ARPHRD_CSLIP                         = 0x101
	ARPHRD_CSLIP6                        = 0x103
	ARPHRD_DDCMP                         = 0x205
	ARPHRD_DLCI                          = 0xf
	ARPHRD_ECONET                        = 0x30e
	ARPHRD_EETHER                        = 0x2
	ARPHRD_ETHER                         = 0x1
	ARPHRD_EUI64                         = 0x1b
	ARPHRD_FCAL                          = 0x311
	ARPHRD_FCFABRIC                      = 0x313
	ARPHRD_FCPL                          = 0x312
	ARPHRD_FCPP                          = 0x310
	ARPHRD_FDDI                          = 0x306
	ARPHRD_FRAD                          = 0x302
	ARPHRD_HDLC                          = 0x201
	ARPHRD_HIPPI                         = 0x30c
	ARPHRD_HWX25                         = 0x110
	ARPHRD_IEEE1394                      = 0x18
	ARPHRD_IEEE802                       = 0x6
	ARPHRD_IEEE80211                     = 0x321
	ARPHRD_IEEE80211_PRISM               = 0x322
	ARPHRD_IEEE80211_RADIOTAP            = 0x323
	ARPHRD_IEEE802154                    = 0x324
	ARPHRD_IEEE802154_MONITOR            = 0x325
	ARPHRD_IEEE802_TR                    = 0x320
	ARPHRD_INFINIBAND                    = 0x20
	ARPHRD_IP6GRE                        = 0x337
	ARPHRD_IPDDP                         = 0x309
	ARPHRD_IPGRE                         = 0x30a
	ARPHRD_IRDA                          = 0x30f
	ARPHRD_LAPB                          = 0x204
	ARPHRD_LOCALTLK                      = 0x305
	ARPHRD_LOOPBACK                      = 0x304
	ARPHRD_METRICOM                      = 0x17
	ARPHRD_NETLINK                       = 0x338
	ARPHRD_NETROM                        = 0x0
	ARPHRD_NONE                          = 0xfffe
	ARPHRD_PHONET                        = 0x334
	ARPHRD_PHONET_PIPE                   = 0x335
	ARPHRD_PIMREG                        = 0x30b
	ARPHRD_PPP                           = 0x200
	ARPHRD_PRONET                        = 0x4
	ARPHRD_RAWHDLC                       = 0x206
	ARPHRD_RAWIP                         = 0x207
	ARPHRD_ROSE                          = 0x10e
	ARPHRD_RSRVD                         = 0x104
	ARPHRD_SIT                           = 0x308
	ARPHRD_SKIP                          = 0x303
	ARPHRD_SLIP                          = 0x100
	ARPHRD_SLIP6                         = 0x102
	ARPHRD_TUNNEL                        = 0x300
	ARPHRD_TUNNEL6                       = 0x301
	ARPHRD_VOID                          = 0xffff
	ARPHRD_VSOCKMON                      = 0x33a
	ARPHRD_X25                           = 0x10f
	AUTOFS_SUPER_MAGIC                   = 0x187
	B0                                   = 0x0
	B1000000                             = 0x1008
	B110                                 = 0x3
	B115200                              = 0x1002
	B1152000                             = 0x1009
	B1200                                = 0x9
	B134                                 = 0x4
	B150                                 = 0x5
	B1500000                             = 0x100a
	B1800                                = 0xa
	B19200                               = 0xe
	B200                                 = 0x6
	B2000000                             = 0x100b
	B230400                              = 0x1003
	B2400                                = 0xb
	B2500000                             = 0x100c
	B300                                 = 0x7
	B3000000                             = 0x100d
	B3500000                             = 0x100e
	B38400                               = 0xf
	B4000000                             = 0x100f
	B460800                              = 0x1004
	B4800                                = 0xc
	B50                                  = 0x1
	B500000                              = 0x1005
	B57600                               = 0x1001
	B576000                              = 0x1006
	B600                                 = 0x8
	B75                                  = 0x2
	B921600                              = 0x1007
	B9600                                = 0xd
	BALLOON_KVM_MAGIC                    = 0x13661366
	BDEVFS_MAGIC                         = 0x62646576
	BINFMTFS_MAGIC                       = 0x42494e4d
	BLKBSZGET                            = 0x80081270
	BLKBSZSET                            = 0x40081271
	BLKFLSBUF                            = 0x1261
	BLKFRAGET                            = 0x1265
	BLKFRASET                            = 0x1264
	BLKGETSIZE                           = 0x1260
	BLKGETSIZE64                         = 0x80081272
	BLKPBSZGET                           = 0x127b
	BLKRAGET                             = 0x1263
	BLKRASET                             = 0x1262
	BLKROGET                             = 0x125e
	BLKROSET                             = 0x125d
	BLKRRPART                            = 0x125f
	BLKSECTGET                           = 0x1267
	BLKSECTSET                           = 0x1266
	BLKSSZGET                            = 0x1268
	BOTHER                               = 0x1000
	BPF_A                                = 0x10
	BPF_ABS                              = 0x20
	BPF_ADD                              = 0x0
	BPF_ALU                              = 0x4
	BPF_AND                              = 0x50
	BPF_B                                = 0x10
	BPF_DIV                              = 0x30
	BPF_FS_MAGIC                         = 0xcafe4a11
	BPF_H                                = 0x8
	BPF_IMM                              = 0x0
	BPF_IND                              = 0x40
	BPF_JA                               = 0x0
	BPF_JEQ                              = 0x10
	BPF_JGE                              = 0x30
	BPF_JGT                              = 0x20
	BPF_JMP                              = 0x5
	BPF_JSET                             = 0x40
	BPF_K                                = 0x0
	BPF_LD                               = 0x0
	BPF_LDX                              = 0x1
	BPF_LEN                              = 0x80
	BPF_LL_OFF                           = -0x200000
	BPF_LSH                              = 0x60
	BPF_MAJOR_VERSION                    = 0x1
	BPF_MAXINSNS                         = 0x1000
	BPF_MEM                              = 0x60
	BPF_MEMWORDS                         = 0x10
	BPF_MINOR_VERSION                    = 0x1
	BPF_MISC                             = 0x7
	BPF_MOD                              = 0x90
	BPF_MSH                              = 0xa0
	BPF_MUL                              = 0x20
	BPF_NEG                              = 0x80
	BPF_NET_OFF                          = -0x100000
	BPF_OR                               = 0x40
	BPF_RET                              = 0x6
	BPF_RSH                              = 0x70
	BPF_ST                               = 0x2
	BPF_STX                              = 0x3
	BPF_SUB                              = 0x10
	BPF_TAX                              = 0x0
	BPF_TXA                              = 0x80
	BPF_W                                = 0x0
	BPF_X                                = 0x8
	BPF_XOR                              = 0xa0
	BRKINT                               = 0x2
	BS0                                  = 0x0
	BS1                                  = 0x2000
	BSDLY                                = 0x2000
	BTRFS_SUPER_MAGIC                    = 0x9123683e
	BTRFS_TEST_MAGIC                     = 0x73727279
	CAN_BCM                              = 0x2
	CAN_EFF_FLAG                         = 0x80000000
	CAN_EFF_ID_BITS                      = 0x1d
	CAN_EFF_MASK                         = 0x1fffffff
	CAN_ERR_FLAG                         = 0x20000000
	CAN_ERR_MASK                         = 0x1fffffff
	CAN_INV_FILTER                       = 0x20000000
	CAN_ISOTP                            = 0x6
	CAN_MAX_DLC                          = 0x8
	CAN_MAX_DLEN                         = 0x8
	CAN_MCNET                            = 0x5
	CAN_MTU                              = 0x10
	CAN_NPROTO                           = 0x7
	CAN_RAW                              = 0x1
	CAN_RAW_FILTER_MAX                   = 0x200
	CAN_RTR_FLAG                         = 0x40000000
	CAN_SFF_ID_BITS                      = 0xb
	CAN_SFF_MASK                         = 0x7ff
	CAN_TP16                             = 0x3
	CAN_TP20                             = 0x4
	CBAUD                                = 0x100f
	CBAUDEX                              = 0x1000
	CFLUSH                               = 0xf
	CGROUP2_SUPER_MAGIC                  = 0x63677270
	CGROUP_SUPER_MAGIC                   = 0x27e0eb
	CIBAUD                               = 0x100f0000
	CLOCAL                               = 0x800
	CLOCK_BOOTTIME                       = 0x7
	CLOCK_BOOTTIME_ALARM                 = 0x9
	CLOCK_DEFAULT                        = 0x0
	CLOCK_EXT                            = 0x1
	CLOCK_INT                            = 0x2
	CLOCK_MONOTONIC                      = 0x1
	CLOCK_MONOTONIC_COARSE               = 0x6
	CLOCK_MONOTONIC_RAW                  = 0x4
	CLOCK_PROCESS_CPUTIME_ID             = 0x2
	CLOCK_REALTIME                       = 0x0
	CLOCK_REALTIME_ALARM                 = 0x8
	CLOCK_REALTIME_COARSE                = 0x5
	CLOCK_TAI                            = 0xb
	CLOCK_THREAD_CPUTIME_ID              = 0x3
	CLOCK_TXFROMRX                       = 0x4
	CLOCK_TXINT                          = 0x3
	CLONE_CHILD_CLEARTID                 = 0x200000
	CLONE_CHILD_SETTID                   = 0x1000000
	CLONE_DETACHED                       = 0x400000
	CLONE_FILES                          = 0x400
	CLONE_FS                             = 0x200
	CLONE_IO                             = 0x80000000
	CLONE_NEWCGROUP                      = 0x2000000
	CLONE_NEWIPC                         = 0x8000000
	CLONE_NEWNET                         = 0x40000000
	CLONE_NEWNS                          = 0x20000
	CLONE_NEWPID                         = 0x20000000
	CLONE_NEWUSER                        = 0x10000000
	CLONE_NEWUTS                         = 0x4000000
	CLONE_PARENT                         = 0x8000
	CLONE_PARENT_SETTID                  = 0x100000
	CLONE_PTRACE                         = 0x2000
	CLONE_SETTLS                         = 0x80000
	CLONE_SIGHAND                        = 0x800
	CLONE_SYSVSEM                        = 0x40000
	CLONE_THREAD                         = 0x10000
	CLONE_UNTRACED                       = 0x800000
	CLONE_VFORK                          = 0x4000
	CLONE_VM                             = 0x100
	CMSPAR                               = 0x40000000
	CODA_SUPER_MAGIC                     = 0x73757245
	CR0                                  = 0x0
	CR1                                  = 0x200
	CR2                                  = 0x400
	CR3                                  = 0x600
	CRAMFS_MAGIC                         = 0x28cd3d45
	CRDLY                                = 0x600
	CREAD                                = 0x80
	CRTSCTS                              = 0x80000000
	CS5                                  = 0x0
	CS6                                  = 0x10
	CS7                                  = 0x20
	CS8                                  = 0x30
	CSIGNAL                              = 0xff
	CSIZE                                = 0x30
	CSTART                               = 0x11
	CSTATUS                              = 0x0
	CSTOP                                = 0x13
	CSTOPB                               = 0x40
	CSUSP                                = 0x1a
	DAXFS_MAGIC                          = 0x64646178
	DEBUGFS_MAGIC                        = 0x64626720
	DEVPTS_SUPER_MAGIC                   = 0x1cd1
	DT_BLK                               = 0x6
	DT_CHR                               = 0x2
	DT_DIR                               = 0x4
	DT_FIFO                              = 0x1
	DT_LNK                               = 0xa
	DT_REG                               = 0x8
	DT_SOCK                              = 0xc
	DT_UNKNOWN                           = 0x0
	DT_WHT                               = 0xe
	ECHO                                 = 0x8
	ECHOCTL                              = 0x200
	ECHOE                                = 0x10
	ECHOK                                = 0x20
	ECHOKE                               = 0x800
	ECHONL                               = 0x40
	ECHOPRT                              = 0x400
	ECRYPTFS_SUPER_MAGIC                 = 0xf15f
	EFD_CLOEXEC                          = 0x80000
	EFD_NONBLOCK                         = 0x800
	EFD_SEMAPHORE                        = 0x1
	EFIVARFS_MAGIC                       = 0xde5e81e4
	EFS_SUPER_MAGIC                      = 0x414a53
	ENCODING_DEFAULT                     = 0x0
	ENCODING_FM_MARK                     = 0x3
	ENCODING_FM_SPACE                    = 0x4
	ENCODING_MANCHESTER                  = 0x5
	ENCODING_NRZ                         = 0x1
	ENCODING_NRZI                        = 0x2
	EPOLLERR                             = 0x8
	EPOLLET                              = 0x80000000
	EPOLLEXCLUSIVE                       = 0x10000000
	EPOLLHUP                             = 0x10
	EPOLLIN                              = 0x1
	EPOLLMSG                             = 0x400
	EPOLLONESHOT                         = 0x40000000
	EPOLLOUT                             = 0x4
	EPOLLPRI                             = 0x2
	EPOLLRDBAND                          = 0x80
	EPOLLRDHUP                           = 0x2000
	EPOLLRDNORM                          = 0x40
	EPOLLWAKEUP                          = 0x20000000
	EPOLLWRBAND                          = 0x200
	EPOLLWRNORM                          = 0x100
	EPOLL_CLOEXEC                        = 0x80000
	EPOLL_CTL_ADD                        = 0x1
	EPOLL_CTL_DEL                        = 0x2
	EPOLL_CTL_MOD                        = 0x3
	ETH_P_1588                           = 0x88f7
	ETH_P_8021AD                         = 0x88a8
	ETH_P_8021AH                         = 0x88e7
	ETH_P_8021Q                          = 0x8100
	ETH_P_80221                          = 0x8917
	ETH_P_802_2                          = 0x4
	ETH_P_802_3                          = 0x1
	ETH_P_802_3_MIN                      = 0x600
	ETH_P_802_EX1                        = 0x88b5
	ETH_P_AARP                           = 0x80f3
	ETH_P_AF_IUCV                        = 0xfbfb
	ETH_P_ALL                            = 0x3
	ETH_P_AOE                            = 0x88a2
	ETH_P_ARCNET                         = 0x1a
	ETH_P_ARP                            = 0x806
	ETH_P_ATALK                          = 0x809b
	ETH_P_ATMFATE                        = 0x8884
	ETH_P_ATMMPOA                        = 0x884c
	ETH_P_AX25                           = 0x2
	ETH_P_BATMAN                         = 0x4305
	ETH_P_BPQ                            = 0x8ff
	ETH_P_CAIF                           = 0xf7
	ETH_P_CAN                            = 0xc
	ETH_P_CANFD                          = 0xd
	ETH_P_CONTROL                        = 0x16
	ETH_P_CUST                           = 0x6006
	ETH_P_DDCMP                          = 0x6
	ETH_P_DEC                            = 0x6000
	ETH_P_DIAG                           = 0x6005
	ETH_P_DNA_DL                         = 0x6001
	ETH_P_DNA_RC                         = 0x6002
	ETH_P_DNA_RT                         = 0x6003
	ETH_P_DSA                            = 0x1b
	ETH_P_ECONET                         = 0x18
	ETH_P_EDSA                           = 0xdada
	ETH_P_ERSPAN                         = 0x88be
	ETH_P_ERSPAN2                        = 0x22eb
	ETH_P_FCOE                           = 0x8906
	ETH_P_FIP                            = 0x8914
	ETH_P_HDLC                           = 0x19
	ETH_P_HSR                            = 0x892f
	ETH_P_IBOE                           = 0x8915
	ETH_P_IEEE802154                     = 0xf6
	ETH_P_IEEEPUP                        = 0xa00
	ETH_P_IEEEPUPAT                      = 0xa01
	ETH_P_IFE                            = 0xed3e
	ETH_P_IP                             = 0x800
	ETH_P_IPV6                           = 0x86dd
	ETH_P_IPX                            = 0x8137
	ETH_P_IRDA                           = 0x17
	ETH_P_LAT                            = 0x6004
	ETH_P_LINK_CTL                       = 0x886c
	ETH_P_LOCALTALK                      = 0x9
	ETH_P_LOOP                           = 0x60
	ETH_P_LOOPBACK                       = 0x9000
	ETH_P_MACSEC                         = 0x88e5
	ETH_P_MAP                            = 0xf9
	ETH_P_MOBITEX                        = 0x15
	ETH_P_MPLS_MC                        = 0x8848
	ETH_P_MPLS_UC                      