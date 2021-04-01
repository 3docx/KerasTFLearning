// mkerrors.sh -m64
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// +build sparc64,linux

// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -m64 _const.go

package unix

import "syscall"

const (
	AF_ALG                           = 0x26
	AF_APPLETALK                     = 0x5
	AF_ASH                           = 0x12
	AF_ATMPVC                        = 0x8
	AF_ATMSVC                        = 0x14
	AF_AX25                          = 0x3
	AF_BLUETOOTH                     = 0x1f
	AF_BRIDGE                        = 0x7
	AF_CAIF                          = 0x25
	AF_CAN                           = 0x1d
	AF_DECnet                        = 0xc
	AF_ECONET                        = 0x13
	AF_FILE                          = 0x1
	AF_IB                            = 0x1b
	AF_IEEE802154                    = 0x24
	AF_INET                          = 0x2
	AF_INET6                         = 0xa
	AF_IPX                           = 0x4
	AF_IRDA                          = 0x17
	AF_ISDN                          = 0x22
	AF_IUCV                          = 0x20
	AF_KCM                           = 0x29
	AF_KEY                           = 0xf
	AF_LLC                           = 0x1a
	AF_LOCAL                         = 0x1
	AF_MAX                           = 0x2a
	AF_MPLS                          = 0x1c
	AF_NETBEUI                       = 0xd
	AF_NETLINK                       = 0x10
	AF_NETROM                        = 0x6
	AF_NFC                           = 0x27
	AF_PACKET                        = 0x11
	AF_PHONET                        = 0x23
	AF_PPPOX                         = 0x18
	AF_RDS                           = 0x15
	AF_ROSE                          = 0xb
	AF_ROUTE                         = 0x10
	AF_RXRPC                         = 0x21
	AF_SECURITY                      = 0xe
	AF_SNA                           = 0x16
	AF_TIPC                          = 0x1e
	AF_UNIX                          = 0x1
	AF_UNSPEC                        = 0x0
	AF_VSOCK                         = 0x28
	AF_WANPIPE                       = 0x19
	AF_X25                           = 0x9
	ALG_OP_DECRYPT                   = 0x0
	ALG_OP_ENCRYPT                   = 0x1
	ALG_SET_AEAD_ASSOCLEN            = 0x4
	ALG_SET_AEAD_AUTHSIZE            = 0x5
	ALG_SET_IV                       = 0x2
	ALG_SET_KEY                      = 0x1
	ALG_SET_OP                       = 0x3
	ARPHRD_6LOWPAN                   = 0x339
	ARPHRD_ADAPT                     = 0x108
	ARPHRD_APPLETLK                  = 0x8
	ARPHRD_ARCNET                    = 0x7
	ARPHRD_ASH                       = 0x30d
	ARPHRD_ATM                       = 0x13
	ARPHRD_AX25                      = 0x3
	ARPHRD_BIF                       = 0x307
	ARPHRD_CAIF                      = 0x336
	ARPHRD_CAN                       = 0x118
	ARPHRD_CHAOS                     = 0x5
	ARPHRD_CISCO                     = 0x201
	ARPHRD_CSLIP                     = 0x101
	ARPHRD_CSLIP6                    = 0x103
	ARPHRD_DDCMP                     = 0x205
	ARPHRD_DLCI                      = 0xf
	ARPHRD_ECONET                    = 0x30e
	ARPHRD_EETHER                    = 0x2
	ARPHRD_ETHER                     = 0x1
	ARPHRD_EUI64                     = 0x1b
	ARPHRD_FCAL                      = 0x311
	ARPHRD_FCFABRIC                  = 0x313
	ARPHRD_FCPL                      = 0x312
	ARPHRD_FCPP                      = 0x310
	ARPHRD_FDDI                      = 0x306
	ARPHRD_FRAD                      = 0x302
	ARPHRD_HDLC                      = 0x201
	ARPHRD_HIPPI                     = 0x30c
	ARPHRD_HWX25                     = 0x110
	ARPHRD_IEEE1394                  = 0x18
	ARPHRD_IEEE802                   = 0x6
	ARPHRD_IEEE80211                 = 0x321
	ARPHRD_IEEE80211_PRISM           = 0x322
	ARPHRD_IEEE80211_RADIOTAP        = 0x323
	ARPHRD_IEEE802154                = 0x324
	ARPHRD_IEEE802154_MONITOR        = 0x325
	ARPHRD_IEEE802_TR                = 0x320
	ARPHRD_INFINIBAND                = 0x20
	ARPHRD_IP6GRE                    = 0x337
	ARPHRD_IPDDP                     = 0x309
	ARPHRD_IPGRE                     = 0x30a
	ARPHRD_IRDA                      = 0x30f
	ARPHRD_LAPB                      = 0x204
	ARPHRD_LOCALTLK                  = 0x305
	ARPHRD_LOOPBACK                  = 0x304
	ARPHRD_METRICOM                  = 0x17
	ARPHRD_NETLINK                   = 0x338
	ARPHRD_NETROM                    = 0x0
	ARPHRD_NONE                      = 0xfffe
	ARPHRD_PHONET                    = 0x334
	ARPHRD_PHONET_PIPE               = 0x335
	ARPHRD_PIMREG                    = 0x30b
	ARPHRD_PPP                       = 0x200
	ARPHRD_PRONET                    = 0x4
	ARPHRD_RAWHDLC                   = 0x206
	ARPHRD_ROSE                      = 0x10e
	ARPHRD_RSRVD                     = 0x104
	ARPHRD_SIT                       = 0x308
	ARPHRD_SKIP                      = 0x303
	ARPHRD_SLIP                      = 0x100
	ARPHRD_SLIP6                     = 0x102
	ARPHRD_TUNNEL                    = 0x300
	ARPHRD_TUNNEL6                   = 0x301
	ARPHRD_VOID                      = 0xffff
	ARPHRD_X25                       = 0x10f
	ASI_LEON_DFLUSH                  = 0x11
	ASI_LEON_IFLUSH                  = 0x10
	ASI_LEON_MMUFLUSH                = 0x18
	B0                               = 0x0
	B1000000                         = 0x100c
	B110                             = 0x3
	B115200                          = 0x1002
	B1152000                         = 0x100d
	B1200                            = 0x9
	B134                             = 0x4
	B150                             = 0x5
	B1500000                         = 0x100e
	B153600                          = 0x1006
	B1800                            = 0xa
	B19200                           = 0xe
	B200                             = 0x6
	B2000000                         = 0x100f
	B230400                          = 0x1003
	B2400                            = 0xb
	B300                             = 0x7
	B307200                          = 0x1007
	B38400                           = 0xf
	B460800                          = 0x1004
	B4800                            = 0xc
	B50                              = 0x1
	B500000                          = 0x100a
	B57600                           = 0x1001
	B576000                          = 0x100b
	B600                             = 0x8
	B614400                          = 0x1008
	B75                              = 0x2
	B76800                           = 0x1005
	B921600                          = 0x1009
	B9600                            = 0xd
	BLKBSZGET                        = 0x80081270
	BLKBSZSET                        = 0x40081271
	BLKFLSBUF                        = 0x1261
	BLKFRAGET                        = 0x1265
	BLKFRASET                        = 0x1264
	BLKGETSIZE                       = 0x1260
	BLKGETSIZE64                     = 0x80081272
	BLKRAGET                         = 0x1263
	BLKRASET                         = 0x1262
	BLKROGET                         = 0x125e
	BLKROSET                         = 0x125d
	BLKRRPART                        = 0x125f
	BLKSECTGET                       = 0x1267
	BLKSECTSET                       = 0x1266
	BLKSSZGET                        = 0x1268
	BOTHER                           = 0x1000
	BPF_A                            = 0x10
	BPF_ABS                          = 0x20
	BPF_ADD                          = 0x0
	BPF_ALU                          = 0x4
	BPF_AND                          = 0x50
	BPF_B                            = 0x10
	BPF_DIV                          = 0x30
	BPF_H                            = 0x8
	BPF_IMM                          = 0x0
	BPF_IND                          = 0x40
	BPF_JA                           = 0x0
	BPF_JEQ                          = 0x10
	BPF_JGE                          = 0x30
	BPF_JGT                          = 0x20
	BPF_JMP                          = 0x5
	BPF_JSET                         = 0x40
	BPF_K                            = 0x0
	BPF_LD                           = 0x0
	BPF_LDX                          = 0x1
	BPF_LEN                          = 0x80
	BPF_LL_OFF                       = -0x200000
	BPF_LSH                          = 0x60
	BPF_MAJOR_VERSION                = 0x1
	BPF_MAXINSNS                     = 0x1000
	BPF_MEM                          = 0x60
	BPF_MEMWORDS                     = 0x10
	BPF_MINOR_VERSION                = 0x1
	BPF_MISC                         = 0x7
	BPF_MOD                          = 0x90
	BPF_MSH                          = 0xa0
	BPF_MUL                          = 0x20
	BPF_NEG                          = 0x80
	BPF_NET_OFF                      = -0x100000
	BPF_OR                           = 0x40
	BPF_RET                          = 0x6
	BPF_RSH                          = 0x70
	BPF_ST                           = 0x2
	BPF_STX                          = 0x3
	BPF_SUB                          = 0x10
	BPF_TAX                          = 0x0
	BPF_TXA                          = 0x80
	BPF_W                            = 0x0
	BPF_X                            = 0x8
	BPF_XOR                          = 0xa0
	BRKINT                           = 0x2
	BS0                              = 0x0
	BS1                              = 0x2000
	BSDLY                            = 0x2000
	CAN_BCM                          = 0x2
	CAN_EFF_FLAG                     = 0x80000000
	CAN_EFF_ID_BITS                  = 0x1d
	CAN_EFF_MASK                     = 0x1fffffff
	CAN_ERR_FLAG                     = 0x20000000
	CAN_ERR_MASK                     = 0x1fffffff
	CAN_INV_FILTER                   = 0x20000000
	CAN_ISOTP                        = 0x6
	CAN_MAX_DLC                      = 0x8
	CAN_MAX_DLEN                     = 0x8
	CAN_MCNET                        = 0x5
	CAN_MTU                          = 0x10
	CAN_NPROTO                       = 0x7
	CAN_RAW                          = 0x1
	CAN_RTR_FLAG                     = 0x40000000
	CAN_SFF_ID_BITS                  = 0xb
	CAN_SFF_MASK                     = 0x7ff
	CAN_TP16                         = 0x3
	CAN_TP20                         = 0x4
	CBAUD                            = 0x100f
	CBAUDEX                          = 0x1000
	CFLUSH                           = 0xf
	CIBAUD                           = 0x100f0000
	CLOCAL                           = 0x800
	CLOCK_BOOTTIME                   = 0x7
	CLOCK_BOOTTIME_ALARM             = 0x9
	CLOCK_DEFAULT                    = 0x0
	CLOCK_EXT                        = 0x1
	CLOCK_INT                        = 0x2
	CLOCK_MONOTONIC                  = 0x1
	CLOCK_MONOTONIC_COARSE           = 0x6
	CLOCK_MONOTONIC_RAW              = 0x4
	CLOCK_PROCESS_CPUTIME_ID         = 0x2
	CLOCK_REALTIME                   = 0x0
	CLOCK_REALTIME_ALARM             = 0x8
	CLOCK_REALTIME_COARSE            = 0x5
	CLOCK_TAI                        = 0xb
	CLOCK_THREAD_CPUTIME_ID          = 0x3
	CLOCK_TXFROMRX                   = 0x4
	CLOCK_TXINT                      = 0x3
	CLONE_CHILD_CLEARTID             = 0x200000
	CLONE_CHILD_SETTID               = 0x1000000
	CLONE_DETACHED                   = 0x400000
	CLONE_FILES                      = 0x400
	CLONE_FS                         = 0x200
	CLONE_IO                         = 0x80000000
	CLONE_NEWCGROUP                  = 0x2000000
	CLONE_NEWIPC                     = 0x8000000
	CLONE_NEWNET                     = 0x40000000
	CLONE_NEWNS                      = 0x20000
	CLONE_NEWPID                     = 0x20000000
	CLONE_NEWUSER                    = 0x10000000
	CLONE_NEWUTS                     = 0x4000000
	CLONE_PARENT                     = 0x8000
	CLONE_PARENT_SETTID              = 0x100000
	CLONE_PTRACE                     = 0x2000
	CLONE_SETTLS                     = 0x80000
	CLONE_SIGHAND                    = 0x800
	CLONE_SYSVSEM                    = 0x40000
	CLONE_THREAD                     = 0x10000
	CLONE_UNTRACED                   = 0x800000
	CLONE_VFORK                      = 0x4000
	CLONE_VM                         = 0x100
	CMSPAR                           = 0x40000000
	CR0                              = 0x0
	CR1                              = 0x200
	CR2                              = 0x400
	CR3                              = 0x600
	CRDLY                            = 0x600
	CREAD                            = 0x80
	CRTSCTS                          = 0x80000000
	CS5                              = 0x0
	CS6                              = 0x10
	CS7                              = 0x20
	CS8                              = 0x30
	CSIGNAL                          = 0xff
	CSIZE                            = 0x30
	CSTART                           = 0x11
	CSTATUS                          = 0x0
	CSTOP                            = 0x13
	CSTOPB                           = 0x40
	CSUSP                            = 0x1a
	DT_BLK                           = 0x6
	DT_CHR                           = 0x2
	DT_DIR                           = 0x4
	DT_FIFO                          = 0x1
	DT_LNK                           = 0xa
	DT_REG                           = 0x8
	DT_SOCK                          = 0xc
	DT_UNKNOWN                       = 0x0
	DT_WHT                           = 0xe
	ECHO                             = 0x8
	ECHOCTL                          = 0x200
	ECHOE                            = 0x10
	ECHOK                            = 0x20
	ECHOKE                           = 0x800
	ECHONL                           = 0x40
	ECHOPRT                          = 0x400
	EMT_TAGOVF                       = 0x1
	ENCODING_DEFAULT                 = 0x0
	ENCODING_FM_MARK                 = 0x3
	ENCODING_FM_SPACE                = 0x4
	ENCODING_MANCHESTER              = 0x5
	ENCODING_NRZ                     = 0x1
	ENCODING_NRZI                    = 0x2
	EPOLLERR                         = 0x8
	EPOLLET                          = 0x80000000
	EPOLLEXCLUSIVE                   = 0x10000000
	EPOLLHUP                         = 0x10
	EPOLLIN                          = 0x1
	EPOLLMSG                         = 0x400
	EPOLLONESHOT                     = 0x40000000
	EPOLLOUT                         = 0x4
	EPOLLPRI                         = 0x2
	EPOLLRDBAND                      = 0x80
	EPOLLRDHUP                       = 0x2000
	EPOLLRDNORM                      = 0x40
	EPOLLWAKEUP                      = 0x20000000
	EPOLLWRBAND                      = 0x200
	EPOLLWRNORM                      = 0x100
	EPOLL_CLOEXEC                    = 0x400000
	EPOLL_CTL_ADD                    = 0x1
	EPOLL_CTL_DEL                    = 0x2
	EPOLL_CTL_MOD                    = 0x3
	ETH_P_1588                       = 0x88f7
	ETH_P_8021AD                     = 0x88a8
	ETH_P_8021AH                     = 0x88e7
	ETH_P_8021Q                      = 0x8100
	ETH_P_80221                      = 0x8917
	ETH_P_802_2                      = 0x4
	ETH_P_802_3                      = 0x1
	ETH_P_802_3_MIN                  = 0x600
	ETH_P_802_EX1                    = 0x88b5
	ETH_P_AARP                       = 0x80f3
	ETH_P_AF_IUCV                    = 0xfbfb
	ETH_P_ALL                        = 0x3
	ETH_P_AOE                        = 0x88a2
	ETH_P_ARCNET                     = 0x1a
	ETH_P_ARP                        = 0x806
	ETH_P_ATALK                      = 0x809b
	ETH_P_ATMFATE                    = 0x8884
	ETH_P_ATMMPOA                    = 0x884c
	ETH_P_AX25                       = 0x2
	ETH_P_BATMAN                     = 0x4305
	ETH_P_BPQ                        = 0x8ff
	ETH_P_CAIF                       = 0xf7
	ETH_P_CAN                        = 0xc
	ETH_P_CANFD                      = 0xd
	ETH_P_CONTROL                    = 0x16
	ETH_P_CUST                       = 0x6006
	ETH_P_DDCMP                      = 0x6
	ETH_P_DEC                        = 0x6000
	ETH_P_DIAG                       = 0x6005
	ETH_P_DNA_DL                     = 0x6001
	ETH_P_DNA_RC                     = 0x6002
	ETH_P_DNA_RT                     = 0x6003
	ETH_P_DSA                        = 0x1b
	ETH_P_ECONET                     = 0x18
	ETH_P_EDSA                       = 0xdada
	ETH_P_FCOE                       = 0x8906
	ETH_P_FIP                        = 0x8914
	ETH_P_HDLC                       = 0x19
	ETH_P_HSR                        = 0x892f
	ETH_P_IEEE802154                 = 0xf6
	ETH_P_IEEEPUP                    = 0xa00
	ETH_P_IEEEPUPAT                  = 0xa01
	ETH_P_IP                         = 0x800
	ETH_P_IPV6                       = 0x86dd
	ETH_P_IPX                        = 0x8137
	ETH_P_IRDA                       = 0x17
	ETH_P_LAT                        = 0x6004
	ETH_P_LINK_CTL                   = 0x886c
	ETH_P_LOCALTALK                  = 0x9
	ETH_P_LOOP                       = 0x60
	ETH_P_LOOPBACK                   = 0x9000
	ETH_P_MACSEC                     = 0x88e5
	ETH_P_MOBITEX                    = 0x15
	ETH_P_MPLS_MC                    = 0x8848
	ETH_P_MPLS_UC                    = 0x8847
	ETH_P_MVRP                       = 0x88f5
	ETH_P_PAE                        = 0x888e
	ETH_P_PAUSE                      = 0x8808
	ETH_P_PHONET                     = 0xf5
	ETH_P_PPPTALK                    = 0x10
	ETH_P_PPP_DISC                   = 0x8863
	ETH_P_PPP_MP                     = 0x8
	ETH_P_PPP_SES                    = 0x8864
	ETH_P_PRP                        = 0x88fb
	ETH_P_PUP                        = 0x200
	ETH_P_PUPAT                      = 0x201
	ETH_P_QINQ1                      = 0x9100
	ETH_P_QINQ2                      = 0x9200
	ETH_P_QINQ3                      = 0x9300
	ETH_P_RARP                       = 0x8035
	ETH_P_SCA                        = 0x6007
	ETH_P_SLOW                       = 0x8809
	ETH_P_SNAP                       = 0x5
	ETH_P_TDLS                       = 0x890d
	ETH_P_TEB                        = 0x6558
	ETH_P_TIPC                       = 0x88ca
	ETH_P_TRAILER                    = 0x1c
	ETH_P_TR_802_2                   = 0x11
	ETH_P_TSN                        = 0x22f0
	ETH_P_WAN_PPP                    = 0x7
	ETH_P_WCCP                       = 0x883e
	ETH_P_X25                        = 0x805
	ETH_P_XDSA                       = 0xf8
	EXTA                             = 0xe
	EXTB                             = 0xf
	EXTPROC                          = 0x10000
	FALLOC_FL_COLLAPSE_RANGE         = 0x8
	FALLOC_FL_INSERT_RANGE           = 0x20
	FALLOC_FL_KEEP_SIZE              = 0x1
	FALLOC_FL_NO_HIDE_STALE          = 0x4
	FALLOC_FL_PUNCH_HOLE             = 0x2
	FALLOC_FL_ZERO_RANGE             = 0x10
	FD_CLOEXEC                       = 0x1
	FD_SETSIZE                       = 0x400
	FF0                              = 0x0
	FF1                              = 0x8000
	FFDLY                            = 0x8000
	FLUSHO                           = 0x2000
	F_DUPFD                          = 0x0
	F_DUPFD_CLOEXEC                  = 0x406
	F_EXLCK                          = 0x4
	F_GETFD                          = 0x1
	F_GETFL                          = 0x3
	F_GETLEASE                       = 0x401
	F_GETLK                          = 0x7
	F_GETLK64                        = 0x7
	F_GETOWN                         = 0x5
	F_GETOWN_EX                      = 0x10
	F_GETPIPE_SZ                     = 0x408
	F_GETSIG                         = 0xb
	F_LOCK                           = 0x1
	F_NOTIFY                         = 0x402
	F_OFD_GETLK                      = 0x24
	F_OFD_SETLK                      = 0x25
	F_OFD_SETLKW                     = 0x26
	F_OK                             = 0x0
	F_RDLCK                          = 0x1
	F_SETFD                          = 0x2
	F_SETFL                          = 0x4
	F_SETLEASE                       = 0x400
	F_SETLK                          = 0x8
	F_SETLK64                        = 0x8
	F_SETLKW                         = 0x9
	F_SETLKW64                       = 0x9
	F_SETOWN                         = 0x6
	F_SETOWN_EX                      = 0xf
	F_SETPIPE_SZ                     = 0x407
	F_SETSIG                         = 0xa
	F_SHLCK                          = 0x8
	F_TEST                           = 0x3
	F_TLOCK                          = 0x2
	F_ULOCK                          = 0x0
	F_UNLCK                          = 0x3
	F_WRLCK                          = 0x2
	GRND_NONBLOCK                    = 0x1
	GRND_RANDOM                      = 0x2
	HUPCL                            = 0x400
	IBSHIFT                          = 0x10
	ICANON                           = 0x2
	ICMPV6_FILTER                    = 0x1
	ICRNL                            = 0x100
	IEXTEN                           = 0x8000
	IFA_F_DADFAILED                  = 0x8
	IFA_F_DEPRECATED                 = 0x20
	IFA_F_HOMEADDRESS                = 0x10
	IFA_F_MANAGETEMPADDR             = 0x100
	IFA_F_MCAUTOJOIN                 = 0x400
	IFA_F_NODAD                      = 0x2
	IFA_F_NOPREFIXROUTE              = 0x200
	IFA_F_OPTIMISTIC                 = 0x4
	IFA_F_PERMANENT                  = 0x80
	IFA_F_SECONDARY                  = 0x1
	IFA_F_STABLE_PRIVACY             = 0x800
	IFA_F_TEMPORARY                  = 0x1
	IFA_F_TENTATIVE                  = 0x40
	IFA_MAX                          = 0x8
	IFF_ALLMULTI                     = 0x200
	IFF_ATTACH_QUEUE                 = 0x200
	IFF_AUTOMEDIA                    = 0x4000
	IFF_BROADCAST                    = 0x2
	IFF_DEBUG                        = 0x4
	IFF_DETACH_QUEUE                 = 0x400
	IFF_DORMANT                      = 0x20000
	IFF_DYNAMIC                      = 0x8000
	IFF_ECHO                         = 0x40000
	IFF_LOOPBACK                     = 0x8
	IFF_LOWER_UP                     = 0x10000
	IFF_MASTER                       = 0x400
	IFF_MULTICAST                    = 0x1000
	IFF_MULTI_QUEUE                  = 0x100
	IFF_NOARP                        = 0x80
	IFF_NOFILTER                     = 0x1000
	IFF_NOTRAILERS                   = 0x20
	IFF_NO_PI                        = 0x1000
	IFF_ONE_QUEUE                    = 0x2000
	IFF_PERSIST                      = 0x800
	IFF_POINTOPOINT                  = 0x10
	IFF_PORTSEL                      = 0x2000
	IFF_PROMISC                      = 0x100
	IFF_RUNNING                      = 0x40
	IFF_SLAVE                        = 0x800
	IFF_TAP                          = 0x2
	IFF_TUN                          = 0x1
	IFF_TUN_EXCL                     = 0x8000
	IFF_UP                           = 0x1
	IFF_VNET_HDR                     = 0x4000
	IFF_VOLATILE                     = 0x70c5a
	IFNAMSIZ                         = 0x10
	IGNBRK                           = 0x1
	IGNCR                            = 0x80
	IGNPAR                           = 0x4
	IMAXBEL                          = 0x2000
	INLCR                            = 0x40
	INPCK                            = 0x10
	IN_ACCESS                        = 0x1
	IN_ALL_EVENTS                    = 0xfff
	IN_ATTRIB                        = 0x4
	IN_CLASSA_HOST                   = 0xffffff
	IN_CLASSA_MAX                    = 0x80
	IN_CLASSA_NET                    = 0xff000000
	IN_CLASSA_NSHIFT                 = 0x18
	IN_CLASSB_HOST                   = 0xffff
	IN_CLASSB_MAX                    = 0x10000
	IN_CLASSB_NET                    = 0xffff0000
	IN_CLASSB_NSHIFT                 = 0x10
	IN_CLASSC_HOST                   = 0xff
	IN_CLASSC_NET                    = 0xffffff00
	IN_CLASSC_NSHIFT                 = 0x8
	IN_CLOEXEC                       = 0x400000
	IN_CLOSE                         = 0x18
	IN_CLOSE_NOWRITE                 = 0x10
	IN_CLOSE_WRITE                   = 0x8
	IN_CREATE                        = 0x100
	IN_DELETE                        = 0x200
	IN_DELETE_SELF                   = 0x400
	IN_DONT_FOLLOW                   = 0x2000000
	IN_EXCL_UNLINK                   = 0x4000000
	IN_IGNORED                       = 0x8000
	IN_ISDIR                         = 0x40000000
	IN_LOOPBACKNET                   = 0x7f
	IN_MASK_ADD                      = 0x20000000
	IN_MODIFY                        = 0x2
	IN_MOVE                          = 0xc0
	IN_MOVED_FROM                    = 0x40
	IN_MOVED_TO                      = 0x80
	IN_MOVE_SELF                     = 0x800
	IN_NONBLOCK                      = 0x4000
	IN_ONESHOT                       = 0x80000000
	IN_ONLYDIR                       = 0x1000000
	IN_OPEN                          = 0x20
	IN_Q_OVERFLOW                    = 0x4000
	IN_UNMOUNT                       = 0x2000
	IPPROTO_AH                       = 0x33
	IPPROTO_BEETPH                   = 0x5e
	IPPROTO_COMP                     = 0x6c
	IPPROTO_DCCP                     = 0x21
	IPPROTO_DSTOPTS                  = 0x3c
	IPPROTO_EGP                      = 0x8
	IPPROTO_ENCAP                    = 0x62
	IPPROTO_ESP                      = 0x32
	IPPROTO_FRAGMENT                 = 0x2c
	IPPROTO_GRE                      = 0x2f
	IPPROTO_HOPOPTS                  = 0x0
	IPPROTO_ICMP                     = 0x1
	IPPROTO_ICMPV6                   = 0x3a
	IPPROTO_IDP                      = 0x16
	IPPROTO_IGMP                     = 0x2
	IPPROTO_IP                       = 0x0
	IPPROTO_IPIP                     = 0x4
	IPPROTO_IPV6                     = 0x29
	IPPROTO_MH                       = 0x87
	IPPROTO_MPLS                     = 0x89
	IPPROTO_MTP                      = 0x5c
	IPPROTO_NONE                     = 0x3b
	IPPROTO_PIM                      = 0x67
	IPPROTO_PUP                      = 0xc
	IPPROTO_RAW                      = 0xff
	IPPROTO_ROUTING                  = 0x2b
	IPPROTO_RSVP                     = 0x2e
	IPPROTO_SCTP                     = 0x84
	IPPROTO_TCP                      = 0x6
	IPPROTO_TP                       = 0x1d
	IPPROTO_UDP                      = 0x11
	IPPROTO_UDPLITE                  = 0x88
	IPV6_2292DSTOPTS                 = 0x4
	IPV6_2292HOPLIMIT                = 0x8
	IPV6_2292HOPOPTS                 = 0x3
	IPV6_2292PKTINFO                 = 0x2
	IPV6_2292PKTOPTIONS              = 0x6
	IPV6_2292RTHDR                   = 0x5
	IPV6_ADDRFORM                    = 0x1
	IPV6_ADD_MEMBERSHIP              = 0x14
	IPV6_AUTHHDR                     = 0xa
	IPV6_CHECKSUM                    = 0x7
	IPV6_DONTFRAG                    = 0x3e
	IPV6_DROP_MEMBERSHIP             = 0x15
	IPV6_DSTOPTS                     = 0x3b
	IPV6_HDRINCL                     = 0x24
	IPV6_HOPLIMIT                    = 0x34
	IPV6_HOPOPTS                     = 0x36
	IPV6_IPSEC_POLICY                = 0x22
	IPV6_JOIN_ANYCAST                = 0x1b
	IPV6_JOIN_GROUP                  = 0x14
	IPV6_LEAVE_ANYCAST               = 0x1c
	IPV6_LEAVE_GROUP                 = 0x15
	IPV6_MTU                         = 0x18
	IPV6_MTU_DISCOVER                = 0x17
	IPV6_MULTICAST_HOPS              = 0x12
	IPV6_MULTICAST_IF                = 0x11
	IPV6_MULTICAST_LOOP              = 0x13
	IPV6_NEXTHOP                     = 0x9
	IPV6_PATHMTU                     = 0x3d
	IPV6_PKTINFO                     = 0x32
	IPV6_PMTUDISC_DO                 = 0x2
	IPV6_PMTUDISC_DONT               = 0x0
	IPV6_PMTUDISC_INTERFACE          = 0x4
	IPV6_PMTUDISC_OMIT               = 0x5
	IPV6_PMTUDISC_PROBE              = 0x3
	IPV6_PMTUDISC_WANT               = 0x1
	IPV6_RECVDSTOPTS                 = 0x3a
	IPV6_RECVERR                     = 0x19
	IPV6_RECVHOPLIMIT                = 0x33
	IPV6_RECVHOPOPTS                 = 0x35
	IPV6_RECVPATHMTU                 = 0x3c
	IPV6_RECVPKTINFO                 = 0x31
	IPV6_RECVRTHDR                   = 0x38
	IPV6_RECVTCLASS                  = 0x42
	IPV6_ROUTER_ALERT                = 0x16
	IPV6_RTHDR                       = 0x39
	IPV6_RTHDRDSTOPTS                = 0x37
	IPV6_RTHDR_LOOSE                 = 0x0
	IPV6_RTHDR_STRICT                = 0x1
	IPV6_RTHDR_TYPE_0                = 0x0
	IPV6_RXDSTOPTS                   = 0x3b
	IPV6_RXHOPOPTS                   = 0x36
	IPV6_TCLASS                      = 0x43
	IPV6_UNICAST_HOPS                = 0x10
	IPV6_V6ONLY                      = 0x1a
	IPV6_XFRM_POLICY                 = 0x23
	IP_ADD_MEMBERSHIP                = 0x23
	IP_ADD_SOURCE_MEMBERSHIP         = 0x27
	IP_BIND_ADDRESS_NO_PORT          = 0x18
	IP_BLOCK_SOURCE                  = 0x26
	IP_CHECKSUM                      = 0x17
	IP_DEFAULT_MULTICAST_LOOP        = 0x1
	IP_DEFAULT_MULTICAST_TTL         = 0x1
	IP_DF                            = 0x4000
	IP_DROP_MEMBERSHIP               = 0x24
	IP_DROP_SOURCE_MEMBERSHIP        = 0x28
	IP_FREEBIND                      = 0xf
	IP_HDRINCL                       = 0x3
	IP_IPSEC_POLICY                  = 0x10
	IP_MAXPACKET                     = 0xffff
	IP_MAX_MEMBERSHIPS               = 0x14
	IP_MF                            = 0x2000
	IP_MINTTL                        = 0x15
	IP_MSFILTER                      = 0x29
	IP_MSS                           = 0x240
	IP_MTU                           = 0xe
	IP_MTU_DISCOVER                  = 0xa
	IP_MULTICAST_ALL                 = 0x31
	IP_MULTICAST_IF                  = 0x20
	IP_MULTICAST_LOOP                = 0x22
	IP_MULTICAST_TTL                 = 0x21
	IP_NODEFRAG                      = 0x16
	IP_OFFMASK                       = 0x1fff
	IP_OPTIONS                       = 0x4
	IP_ORIGDSTADDR                   = 0x14
	IP_PASSSEC                       = 0x12
	IP_PKTINFO                       = 0x8
	IP_PKTOPTIONS                    = 0x9
	IP_PMTUDISC                      = 0xa
	IP_PMTUDISC_DO                   = 0x2
	IP_PMTUDISC_DONT                 = 0x0
	IP_PMTUDISC_INTERFACE            = 0x4
	IP_PMTUDISC_OMIT                 = 0x5
	IP_PMTUDISC_PROBE                = 0x3
	IP_PMTUDISC_WANT                 = 0x1
	IP_RECVERR                       = 0xb
	IP_RECVOPTS                      = 0x6
	IP_RECVORIGDSTADDR               = 0x14
	IP_RECVRETOPTS                   = 0x7
	IP_RECVTOS                       = 0xd
	IP_RECVTTL                       = 0xc
	IP_RETOPTS                       = 0x7
	IP_RF                            = 0x8000
	IP_ROUTER_ALERT                  = 0x5
	IP_TOS                           = 0x1
	IP_TRANSPARENT                   = 0x13
	IP_TTL                           = 0x2
	IP_UNBLOCK_SOURCE                = 0x25
	IP_UNICAST_IF                    = 0x32
	IP_XFRM_POLICY                   = 0x11
	ISIG                             = 0x1
	ISTRIP                           = 0x20
	IUCLC                            = 0x200
	IUTF8                            = 0x4000
	IXANY                            = 0x800
	IXOFF                            = 0x1000
	IXON                             = 0x400
	LINUX_REBOOT_CMD_CAD_OFF         = 0x0
	LINUX_REBOOT_CMD_CAD_ON          = 0x89abcdef
	LINUX_REBOOT_CMD_HALT            = 0xcdef0123
	LINUX_REBOOT_CMD_KEXEC           = 0x45584543
	LINUX_REBOOT_CMD_POWER_OFF       = 0x4321fedc
	LINUX_REBOOT_CMD_RESTART         = 0x1234567
	LINUX_REBOOT_CMD_RESTART2        = 0xa1b2c3d4
	LINUX_REBOOT_CMD_SW_SUSPEND      = 0xd000fce2
	LINUX_REBOOT_MAGIC1              = 0xfee1dead
	LINUX_REBOOT_MAGIC2              = 0x28121969
	LOCK_EX                          = 0x2
	LOCK_NB                          = 0x4
	LOCK_SH                          = 0x1
	LOCK_UN                          = 0x8
	MADV_DODUMP                      = 0x11
	MADV_DOFORK                      = 0xb
	MADV_DONTDUMP                    = 0x10
	MADV_DONTFORK                    = 0xa
	MADV_DONTNEED                    = 0x4
	MADV_FREE                        = 0x8
	MADV_HUGEPAGE                    = 0xe
	MADV_HWPOISON                    = 0x64
	MADV_MERGEABLE                   = 0xc
	MADV_NOHUGEPAGE                  = 0xf
	MADV_NORMAL                      = 0x0
	MADV_RANDOM                      = 0x1
	MADV_REMOVE                      = 0x9
	MADV_SEQUENTIAL                  = 0x2
	MADV_UNMERGEABLE                 = 0xd
	MADV_WILLNEED                    = 0x3
	MAP_ANON                         = 0x20
	MAP_ANONYMOUS                    = 0x20
	MAP_DENYWRITE                    = 0x800
	MAP_EXECUTABLE                   = 0x1000
	MAP_FILE                         = 0x0
	MAP_FIXED                        = 0x10
	MAP_GROWSDOWN                    = 0x200
	MAP_HUGETLB                      = 0x40000
	MAP_HUGE_MASK                    = 0x3f
	MAP_HUGE_SHIFT                   = 0x1a
	MAP_LOCKED                       = 0x100
	MAP_NONBLOCK                     = 0x10000
	MAP_NORESERVE                    = 0x40
	MAP_POPULATE                     = 0x8000
	MAP_PRIVATE                      = 0x2
	MAP_RENAME                       = 0x20
	MAP_SHARED                       = 0x1
	MAP_STACK                        = 0x20000
	MAP_TYPE                         = 0xf
	MCL_CURRENT                      = 0x2000
	MCL_FUTURE                       = 0x4000
	MCL_ONFAULT                      = 0x8000
	MNT_DETACH                       = 0x2
	MNT_EXPIRE                       = 0x4
	MNT_FORCE                        = 0x1
	MSG_BATCH                        = 0x40000
	MSG_CMSG_CLOEXEC                 = 0x40000000
	MSG_CONFIRM                      = 0x800
	MSG_CTRUNC                       = 0x8
	MSG_DONTROUTE                    = 0x4
	MSG_DONTWAIT                     = 0x40
	MSG_EOR                          = 0x80
	MSG_ERRQUEUE                     = 0x2000
	MSG_FASTOPEN                     = 0x20000000
	MSG_FIN                          = 0x200
	MSG_MORE                         = 0x8000
	MSG_NOSIGNAL                     = 0x4000
	MSG_OOB                          = 0x1
	MSG_PEEK                         = 0x2
	MSG_PROXY                        = 0x10
	MSG_RST                          = 0x1000
	MSG_SYN                          = 0x400
	MSG_TRUNC                        = 0x20
	MSG_TRYHARD                      = 0x4
	MSG_WAITALL                      = 0x100
	MSG_WAITFORONE                   = 0x10000
	MS_ACTIVE                        = 0x40000000
	MS_ASYNC                         = 0x1
	MS_BIND                       