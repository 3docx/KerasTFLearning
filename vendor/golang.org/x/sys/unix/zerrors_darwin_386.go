// mkerrors.sh -m32
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,darwin

// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -m32 _const.go

package unix

import "syscall"

const (
	AF_APPLETALK                      = 0x10
	AF_CCITT                          = 0xa
	AF_CHAOS                          = 0x5
	AF_CNT                            = 0x15
	AF_COIP                           = 0x14
	AF_DATAKIT                        = 0x9
	AF_DECnet                         = 0xc
	AF_DLI                            = 0xd
	AF_E164                           = 0x1c
	AF_ECMA                           = 0x8
	AF_HYLINK                         = 0xf
	AF_IEEE80211                      = 0x25
	AF_IMPLINK                        = 0x3
	AF_INET                           = 0x2
	AF_INET6                          = 0x1e
	AF_IPX                            = 0x17
	AF_ISDN                           = 0x1c
	AF_ISO                            = 0x7
	AF_LAT                            = 0xe
	AF_LINK                           = 0x12
	AF_LOCAL                          = 0x1
	AF_MAX                            = 0x28
	AF_NATM                           = 0x1f
	AF_NDRV                           = 0x1b
	AF_NETBIOS                        = 0x21
	AF_NS                             = 0x6
	AF_OSI                            = 0x7
	AF_PPP                            = 0x22
	AF_PUP                            = 0x4
	AF_RESERVED_36                    = 0x24
	AF_ROUTE                          = 0x11
	AF_SIP                            = 0x18
	AF_SNA                            = 0xb
	AF_SYSTEM                         = 0x20
	AF_UNIX                           = 0x1
	AF_UNSPEC                         = 0x0
	AF_UTUN                           = 0x26
	ALTWERASE                         = 0x200
	ATTR_BIT_MAP_COUNT                = 0x5
	ATTR_CMN_ACCESSMASK               = 0x20000
	ATTR_CMN_ACCTIME                  = 0x1000
	ATTR_CMN_ADDEDTIME                = 0x10000000
	ATTR_CMN_BKUPTIME                 = 0x2000
	ATTR_CMN_CHGTIME                  = 0x800
	ATTR_CMN_CRTIME                   = 0x200
	ATTR_CMN_DATA_PROTECT_FLAGS       = 0x40000000
	ATTR_CMN_DEVID                    = 0x2
	ATTR_CMN_DOCUMENT_ID              = 0x100000
	ATTR_CMN_ERROR                    = 0x20000000
	ATTR_CMN_EXTENDED_SECURITY        = 0x400000
	ATTR_CMN_FILEID                   = 0x2000000
	ATTR_CMN_FLAGS                    = 0x40000
	ATTR_CMN_FNDRINFO                 = 0x4000
	ATTR_CMN_FSID                     = 0x4
	ATTR_CMN_FULLPATH                 = 0x8000000
	ATTR_CMN_GEN_COUNT                = 0x80000
	ATTR_CMN_GRPID                    = 0x10000
	ATTR_CMN_GRPUUID                  = 0x1000000
	ATTR_CMN_MODTIME                  = 0x400
	ATTR_CMN_NAME                     = 0x1
	ATTR_CMN_NAMEDATTRCOUNT           = 0x80000
	ATTR_CMN_NAMEDATTRLIST            = 0x100000
	ATTR_CMN_OBJID                    = 0x20
	ATTR_CMN_OBJPERMANENTID           = 0x40
	ATTR_CMN_OBJTAG                   = 0x10
	ATTR_CMN_OBJTYPE                  = 0x8
	ATTR_CMN_OWNERID                  = 0x8000
	ATTR_CMN_PARENTID                 = 0x4000000
	ATTR_CMN_PAROBJID                 = 0x80
	ATTR_CMN_RETURNED_ATTRS           = 0x80000000
	ATTR_CMN_SCRIPT                   = 0x100
	ATTR_CMN_SETMASK                  = 0x41c7ff00
	ATTR_CMN_USERACCESS               = 0x200000
	ATTR_CMN_UUID                     = 0x800000
	ATTR_CMN_VALIDMASK                = 0xffffffff
	ATTR_CMN_VOLSETMASK               = 0x6700
	ATTR_FILE_ALLOCSIZE               = 0x4
	ATTR_FILE_CLUMPSIZE               = 0x10
	ATTR_FILE_DATAALLOCSIZE           = 0x400
	ATTR_FILE_DATAEXTENTS             = 0x800
	ATTR_FILE_DATALENGTH              = 0x200
	ATTR_FILE_DEVTYPE                 = 0x20
	ATTR_FILE_FILETYPE                = 0x40
	ATTR_FILE_FORKCOUNT               = 0x80
	ATTR_FILE_FORKLIST                = 0x100
	ATTR_FILE_IOBLOCKSIZE             = 0x8
	ATTR_FILE_LINKCOUNT               = 0x1
	ATTR_FILE_RSRCALLOCSIZE           = 0x2000
	ATTR_FILE_RSRCEXTENTS             = 0x4000
	ATTR_FILE_RSRCLENGTH              = 0x1000
	ATTR_FILE_SETMASK                 = 0x20
	ATTR_FILE_TOTALSIZE               = 0x2
	ATTR_FILE_VALIDMASK               = 0x37ff
	ATTR_VOL_ALLOCATIONCLUMP          = 0x40
	ATTR_VOL_ATTRIBUTES               = 0x40000000
	ATTR_VOL_CAPABILITIES             = 0x20000
	ATTR_VOL_DIRCOUNT                 = 0x400
	ATTR_VOL_ENCODINGSUSED            = 0x10000
	ATTR_VOL_FILECOUNT                = 0x200
	ATTR_VOL_FSTYPE                   = 0x1
	ATTR_VOL_INFO                     = 0x80000000
	ATTR_VOL_IOBLOCKSIZE              = 0x80
	ATTR_VOL_MAXOBJCOUNT              = 0x800
	ATTR_VOL_MINALLOCATION            = 0x20
	ATTR_VOL_MOUNTEDDEVICE            = 0x8000
	ATTR_VOL_MOUNTFLAGS               = 0x4000
	ATTR_VOL_MOUNTPOINT               = 0x1000
	ATTR_VOL_NAME                     = 0x2000
	ATTR_VOL_OBJCOUNT                 = 0x100
	ATTR_VOL_QUOTA_SIZE               = 0x10000000
	ATTR_VOL_RESERVED_SIZE            = 0x20000000
	ATTR_VOL_SETMASK                  = 0x80002000
	ATTR_VOL_SIGNATURE                = 0x2
	ATTR_VOL_SIZE                     = 0x4
	ATTR_VOL_SPACEAVAIL               = 0x10
	ATTR_VOL_SPACEFREE                = 0x8
	ATTR_VOL_UUID                     = 0x40000
	ATTR_VOL_VALIDMASK                = 0xf007ffff
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
	B4800                             = 0x12c0
	B50                               = 0x32
	B57600                            = 0xe100
	B600                              = 0x258
	B7200                             = 0x1c20
	B75                               = 0x4b
	B76800                            = 0x12c00
	B9600                             = 0x2580
	BIOCFLUSH                         = 0x20004268
	BIOCGBLEN                         = 0x40044266
	BIOCGDLT                          = 0x4004426a
	BIOCGDLTLIST                      = 0xc00c4279
	BIOCGETIF                         = 0x4020426b
	BIOCGHDRCMPLT                     = 0x40044274
	BIOCGRSIG                         = 0x40044272
	BIOCGRTIMEOUT                     = 0x4008426e
	BIOCGSEESENT                      = 0x40044276
	BIOCGSTATS                        = 0x4008426f
	BIOCIMMEDIATE                     = 0x80044270
	BIOCPROMISC                       = 0x20004269
	BIOCSBLEN                         = 0xc0044266
	BIOCSDLT                          = 0x80044278
	BIOCSETF                          = 0x80084267
	BIOCSETFNR                        = 0x8008427e
	BIOCSETIF                         = 0x8020426c
	BIOCSHDRCMPLT                     = 0x80044275
	BIOCSRSIG                         = 0x80044273
	BIOCSRTIMEOUT                     = 0x8008426d
	BIOCSSEESENT                      = 0x80044277
	BIOCVERSION                       = 0x40044271
	BPF_A                             = 0x10
	BPF_ABS                           = 0x20
	BPF_ADD                           = 0x0
	BPF_ALIGNMENT                     = 0x4
	BPF_ALU                           = 0x4
	BPF_AND                           = 0x50
	BPF_B                             = 0x10
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
	BPF_MAXBUFSIZE                    = 0x80000
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
	BS0                               = 0x0
	BS1                               = 0x8000
	BSDLY                             = 0x8000
	CFLUSH                            = 0xf
	CLOCAL                            = 0x8000
	CLOCK_MONOTONIC                   = 0x6
	CLOCK_MONOTONIC_RAW               = 0x4
	CLOCK_MONOTONIC_RAW_APPROX        = 0x5
	CLOCK_PROCESS_CPUTIME_ID          = 0xc
	CLOCK_REALTIME                    = 0x0
	CLOCK_THREAD_CPUTIME_ID           = 0x10
	CLOCK_UPTIME_RAW                  = 0x8
	CLOCK_UPTIME_RAW_APPROX           = 0x9
	CR0                               = 0x0
	CR1                               = 0x1000
	CR2                               = 0x2000
	CR3                               = 0x3000
	CRDLY                             = 0x3000
	CREAD                             = 0x800
	CRTSCTS                           = 0x30000
	CS5                               = 0x0
	CS6                               = 0x100
	CS7                               = 0x200
	CS8                               = 0x300
	CSIZE                             = 0x300
	CSTART                            = 0x11
	CSTATUS                           = 0x14
	CSTOP                             = 0x13
	CSTOPB                            = 0x400
	CSUSP                             = 0x1a
	CTL_HW                            = 0x6
	CTL_KERN                          = 0x1
	CTL_MAXNAME                       = 0xc
	CTL_NET                           = 0x4
	DLT_A429                          = 0xb8
	DLT_A653_ICM                      = 0xb9
	DLT_AIRONET_HEADER                = 0x78
	DLT_AOS                           = 0xde
	DLT_APPLE_IP_OVER_IEEE1394        = 0x8a
	DLT_ARCNET                        = 0x7
	DLT_ARCNET_LINUX                  = 0x81
	DLT_ATM_CLIP                      = 0x13
	DLT_ATM_RFC1483                   = 0xb
	DLT_AURORA                        = 0x7e
	DLT_AX25                          = 0x3
	DLT_AX25_KISS                     = 0xca
	DLT_BACNET_MS_TP                  = 0xa5
	DLT_BLUETOOTH_HCI_H4              = 0xbb
	DLT_BLUETOOTH_HCI_H4_WITH_PHDR    = 0xc9
	DLT_CAN20B                        = 0xbe
	DLT_CAN_SOCKETCAN                 = 0xe3
	DLT_CHAOS                         = 0x5
	DLT_CHDLC                         = 0x68
	DLT_CISCO_IOS                     = 0x76
	DLT_C_HDLC                        = 0x68
	DLT_C_HDLC_WITH_DIR               = 0xcd
	DLT_DBUS                          = 0xe7
	DLT_DECT                          = 0xdd
	DLT_DOCSIS                        = 0x8f
	DLT_DVB_CI                        = 0xeb
	DLT_ECONET                        = 0x73
	DLT_EN10MB                        = 0x1
	DLT_EN3MB                         = 0x2
	DLT_ENC                           = 0x6d
	DLT_ERF                           = 0xc5
	DLT_ERF_ETH                       = 0xaf
	DLT_ERF_POS                       = 0xb0
	DLT_FC_2                          = 0xe0
	DLT_FC_2_WITH_FRAME_DELIMS        = 0xe1
	DLT_FDDI                          = 0xa
	DLT_FLEXRAY                       = 0xd2
	DLT_FRELAY                        = 0x6b
	DLT_FRELAY_WITH_DIR               = 0xce
	DLT_GCOM_SERIAL                   = 0xad
	DLT_GCOM_T1E1                     = 0xac
	DLT_GPF_F                         = 0xab
	DLT_GPF_T                         = 0xaa
	DLT_GPRS_LLC                      = 0xa9
	DLT_GSMTAP_ABIS                   = 0xda
	DLT_GSMTAP_UM                     = 0xd9
	DLT_HHDLC                         = 0x79
	DLT_IBM_SN                        = 0x92
	DLT_IBM_SP                        = 0x91
	DLT_IEEE802                       = 0x6
	DLT_IEEE802_11                    = 0x69
	DLT_IEEE802_11_RADIO              = 0x7f
	DLT_IEEE802_11_RADIO_AVS          = 0xa3
	DLT_IEEE802_15_4                  = 0xc3
	DLT_IEEE802_15_4_LINUX            = 0xbf
	DLT_IEEE802_15_4_NOFCS            = 0xe6
	DLT_IEEE802_15_4_NONASK_PHY       = 0xd7
	DLT_IEEE802_16_MAC_CPS            = 0xbc
	DLT_IEEE802_16_MAC_CPS_RADIO      = 0xc1
	DLT_IPFILTER                      = 0x74
	DLT_IPMB                          = 0xc7
	DLT_IPMB_LINUX                    = 0xd1
	DLT_IPNET                         = 0xe2
	DLT_IPOIB                         = 0xf2
	DLT_IPV4                          = 0xe4
	DLT_IPV6                          = 0xe5
	DLT_IP_OVER_FC                    = 0x7a
	DLT_JUNIPER_ATM1                  = 0x89
	DLT_JUNIPER_ATM2                  = 0x87
	DLT_JUNIPER_ATM_CEMIC             = 0xee
	DLT_JUNIPER_CHDLC                 = 0xb5
	DLT_JUNIPER_ES                    = 0x84
	DLT_JUNIPER_ETHER                 = 0xb2
	DLT_JUNIPER_FIBRECHANNEL          = 0xea
	DLT_JUNIPER_FRELAY                = 0xb4
	DLT_JUNIPER_GGSN                  = 0x85
	DLT_JUNIPER_ISM                   = 0xc2
	DLT_JUNIPER_MFR                   = 0x86
	DLT_JUNIPER_MLFR                  = 0x83
	DLT_JUNIPER_MLPPP                 = 0x82
	DLT_JUNIPER_MONITOR               = 0xa4
	DLT_JUNIPER_PIC_PEER              = 0xae
	DLT_JUNIPER_PPP                   = 0xb3
	DLT_JUNIPER_PPPOE                 = 0xa7
	DLT_JUNIPER_PPPOE_ATM             = 0xa8
	DLT_JUNIPER_SERVICES              = 0x88
	DLT_JUNIPER_SRX_E2E               = 0xe9
	DLT_JUNIPER_ST                    = 0xc8
	DLT_JUNIPER_VP                    = 0xb7
	DLT_JUNIPER_VS                    = 0xe8
	DLT_LAPB_WITH_DIR                 = 0xcf
	DLT_LAPD                          = 0xcb
	DLT_LIN                           = 0xd4
	DLT_LINUX_EVDEV                   = 0xd8
	DLT_LINUX_IRDA                    = 0x90
	DLT_LINUX_LAPD                    = 0xb1
	DLT_LINUX_PPP_WITHDIRECTION       = 0xa6
	DLT_LINUX_SLL                     = 0x71
	DLT_LOOP                          = 0x6c
	DLT_LTALK                         = 0x72
	DLT_MATCHING_MAX                  = 0xf5
	DLT_MATCHING_MIN                  = 0x68
	DLT_MFR                           = 0xb6
	DLT_MOST                          = 0xd3
	DLT_MPEG_2_TS                     = 0xf3
	DLT_MPLS                          = 0xdb
	DLT_MTP2                          = 0x8c
	DLT_MTP2_WITH_PHDR                = 0x8b
	DLT_MTP3                          = 0x8d
	DLT_MUX27010                      = 0xec
	DLT_NETANALYZER                   = 0xf0
	DLT_NETANALYZER_TRANSPARENT       = 0xf1
	DLT_NFC_LLCP              