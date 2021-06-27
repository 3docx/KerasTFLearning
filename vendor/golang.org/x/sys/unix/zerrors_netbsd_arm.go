// mkerrors.sh -marm
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build arm,netbsd

// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -marm _const.go

package unix

import "syscall"

const (
	AF_APPLETALK                      = 0x10
	AF_ARP                            = 0x1c
	AF_BLUETOOTH                      = 0x1f
	AF_CCITT                          = 0xa
	AF_CHAOS                          = 0x5
	AF_CNT                            = 0x15
	AF_COIP                           = 0x14
	AF_DATAKIT                        = 0x9
	AF_DECnet                         = 0xc
	AF_DLI                            = 0xd
	AF_E164                           = 0x1a
	AF_ECMA                           = 0x8
	AF_HYLINK                         = 0xf
	AF_IEEE80211                      = 0x20
	AF_IMPLINK                        = 0x3
	AF_INET                           = 0x2
	AF_INET6                          = 0x18
	AF_IPX                            = 0x17
	AF_ISDN                           = 0x1a
	AF_ISO                            = 0x7
	AF_LAT                            = 0xe
	AF_LINK                           = 0x12
	AF_LOCAL                          = 0x1
	AF_MAX                            = 0x23
	AF_MPLS                           = 0x21
	AF_NATM                           = 0x1b
	AF_NS                             = 0x6
	AF_OROUTE                         = 0x11
	AF_OSI                            = 0x7
	AF_PUP                            = 0x4
	AF_ROUTE                          = 0x22
	AF_SNA                            = 0xb
	AF_UNIX                           = 0x1
	AF_UNSPEC                         = 0x0
	ARPHRD_ARCNET                     = 0x7
	ARPHRD_ETHER                      = 0x1
	ARPHRD_FRELAY                     = 0xf
	ARPHRD_IEEE1394                   = 0x18
	ARPHRD_IEEE802                    = 0x6
	ARPHRD_STRIP                      = 0x17
	B0                                = 0x0
	B110                              = 0x6e
	B115200                           = 0x1c200
	B1200                             = 0x4b0
	B134                              = 0x86
	B14400                            = 0x3840
	B150                              = 0x96
	B1800                             = 0x708
	B19200                            = 0x4b00
	B200                              = 0xc8
	B230400                           = 0x38400
	B2400                             = 0x960
	B28800                            = 0x7080
	B300                              = 0x12c
	B38400                            = 0x9600
	B460800                           = 0x70800
	B4800                             = 0x12c0
	B50                               = 0x32
	B57600                            = 0xe100
	B600                              = 0x258
	B7200                             = 0x1c20
	B75                               = 0x4b
	B76800                            = 0x12c00
	B921600                           = 0xe1000
	B9600                             = 0x2580
	BIOCFEEDBACK                      = 0x8004427d
	BIOCFLUSH                         = 0x20004268
	BIOCGBLEN                         = 0x40044266
	BIOCGDLT                          = 0x4004426a
	BIOCGDLTLIST                      = 0xc0084277
	BIOCGETIF                         = 0x4090426b
	BIOCGFEEDBACK                     = 0x4004427c
	BIOCGHDRCMPLT                     = 0x40044274
	BIOCGRTIMEOUT                     = 0x400c427b
	BIOCGSEESENT                      = 0x40044278
	BIOCGSTATS                        = 0x4080426f
	BIOCGSTATSOLD                     = 0x4008426f
	BIOCIMMEDIATE                     = 0x80044270
	BIOCPROMISC                       = 0x20004269
	BIOCSBLEN                         = 0xc0044266
	BIOCSDLT                          = 0x80044276
	BIOCSETF                          = 0x80084267
	BIOCSETIF                         = 0x8090426c
	BIOCSFEEDBACK                     = 0x8004427d
	BIOCSHDRCMPLT                     = 0x80044275
	BIOCSRTIMEOUT                     = 0x800c427a
	BIOCSSEESENT                      = 0x80044279
	BIOCSTCPF                         = 0x80084272
	BIOCSUDPF                         = 0x80084273
	BIOCVERSION                       = 0x40044271
	BPF_A                             = 0x10
	BPF_ABS                           = 0x20
	BPF_ADD                           = 0x0
	BPF_ALIGNMENT                     = 0x4
	BPF_ALIGNMENT32                   = 0x4
	BPF_ALU                           = 0x4
	BPF_AND                           = 0x50
	BPF_B                             = 0x10
	BPF_DFLTBUFSIZE                   = 0x100000
	BPF_DIV                           = 0x30
	BPF_H                             = 0x8
	BPF_IMM                           = 0x0
	BPF_IND                           = 0x40
	BPF_JA                            = 0x0
	BPF_JEQ                           = 0x10
	BPF_JGE                           = 0x30
	BPF_JGT                           = 0x20
	BPF_JMP                           = 0x5
	BPF_JSET                          = 0x40
	BPF_K                             = 0x0
	BPF_LD                            = 0x0
	BPF_LDX                           = 0x1
	BPF_LEN                           = 0x80
	BPF_LSH                           = 0x60
	BPF_MAJOR_VERSION                 = 0x1
	BPF_MAXBUFSIZE                    = 0x1000000
	BPF_MAXINSNS                      = 0x200
	BPF_MEM                           = 0x60
	BPF_MEMWORDS                      = 0x10
	BPF_MINBUFSIZE                    = 0x20
	BPF_MINOR_VERSION                 = 0x1
	BPF_MISC                          = 0x7
	BPF_MSH                           = 0xa0
	BPF_MUL                           = 0x20
	BPF_NEG                           = 0x80
	BPF_OR                            = 0x40
	BPF_RELEASE                       = 0x30bb6
	BPF_RET                           = 0x6
	BPF_RSH                           = 0x70
	BPF_ST                            = 0x2
	BPF_STX                           = 0x3
	BPF_SUB                           = 0x10
	BPF_TAX                           = 0x0
	BPF_TXA                           = 0x80
	BPF_W                             = 0x0
	BPF_X                             = 0x8
	BRKINT                            = 0x2
	CFLUSH                            = 0xf
	CLOCAL                            = 0x8000
	CREAD                             = 0x800
	CRTSCTS                           = 0x10000
	CS5                               = 0x0
	CS6                               = 0x100
	CS7                          