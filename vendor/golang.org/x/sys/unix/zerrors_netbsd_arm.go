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
	CTL_QUERY                         = -0x2
	DIOCBSFLUSH                       = 0x20006478
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
	DLT_CISCO_IOS                     = 0x76
	DLT_C_HDLC                        = 0x68
	DLT_C_HDLC_WITH_DIR               = 0xcd
	DLT_DECT                          = 0xdd
	DLT_DOCSIS                        = 0x8f
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
	DLT_HDLC                          = 0x10
	DLT_HHDLC                         = 0x79
	DLT_HIPPI                         = 0xf
	DLT_IBM_SN                        = 0x92
	DLT_IBM_SP                        = 0x91
	DLT_IEEE802                       = 0x6
	DLT_IEEE802_11                    = 0x69
	DLT_IEEE802_11_RADIO              = 0x7f
	DLT_IEEE802_11_RADIO_AVS          = 0xa3
	DLT_IEEE802_15_4                  = 0xc3
	DLT_IEEE802_15_4_LINUX            = 0xbf
	DLT_IEEE802_15_4_NONASK_PHY       = 0xd7
	DLT_IEEE802_16_MAC_CPS            = 0xbc
	DLT_IEEE802_16_MAC_CPS_RADIO      = 0xc1
	DLT_IPMB                          = 0xc7
	DLT_IPMB_LINUX                    = 0xd1
	DLT_IPNET                         = 0xe2
	DLT_IPV4                          = 0xe4
	DLT_IPV6                          = 0xe5
	DLT_IP_OVER_FC                    = 0x7a
	DLT_JUNIPER_ATM1                  = 0x89
	DLT_JUNIPER_ATM2                  = 0x87
	DLT_JUNIPER_CHDLC                 = 0xb5
	DLT_JUNIPER_ES                    = 0x84
	DLT_JUNIPER_ETHER                 = 0xb2
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
	DLT_JUNIPER_ST                    = 0xc8
	DLT_JUNIPER_VP                    = 0xb7
	DLT_LAPB_WITH_DIR                 = 0xcf
	DLT_LAPD                          = 0xcb
	DLT_LIN                           = 0xd4
	DLT_LINUX_EVDEV                   = 0xd8
	DLT_LINUX_IRDA                    = 0x90
	DLT_LINUX_LAPD                    = 0xb1
	DLT_LINUX_SLL                     = 0x71
	DLT_LOOP                          = 0x6c
	DLT_LTALK                         = 0x72
	DLT_MFR                           = 0xb6
	DLT_MOST                          = 0xd3
	DLT_MPLS                          = 0xdb
	DLT_MTP2                          = 0x8c
	DLT_MTP2_WITH_PHDR                = 0x8b
	DLT_MTP3                          = 0x8d
	DLT_NULL                          = 0x0
	DLT_PCI_EXP                       = 0x7d
	DLT_PFLOG                         = 0x75
	DLT_PFSYNC                        = 0x12
	DLT_PPI                           = 0xc0
	DLT_PPP                           = 0x9
	DLT_PPP_BSDOS                     = 0xe
	DLT_PPP_ETHER                     = 0x33
	DLT_PPP_PPPD                      = 0xa6
	DLT_PPP_SERIAL                    = 0x32
	DLT_PPP_WITH_DIR                  = 0xcc
	DLT_PRISM_HEADER                  = 0x77
	DLT_PRONET                        = 0x4
	DLT_RAIF1                         = 0xc6
	DLT_RAW                           = 0xc
	DLT_RAWAF_MASK                    = 0x2240000
	DLT_RIO                           = 0x7c
	DLT_SCCP                          = 0x8e
	DLT_SITA                          = 0xc4
	DLT_SLIP                          = 0x8
	DLT_SLIP_BSDOS                    = 0xd
	DLT_SUNATM                        = 0x7b
	DLT_SYMANTEC_FIREWALL             = 0x63
	DLT_TZSP                          = 0x80
	DLT_USB                           = 0xba
	DLT_USB_LINUX                     = 0xbd
	DLT_USB_LINUX_MMAPPED             = 0xdc
	DLT_WIHART                        = 0xdf
	DLT_X2E_SERIAL                    = 0xd5
	DLT_X2E_XORAYA                    = 0xd6
	DT_BLK                            = 0x6
	DT_CHR                            = 0x2
	DT_DIR                            = 0x4
	DT_FIFO                           = 0x1
	DT_LNK                            = 0xa
	DT_REG                            = 0x8
	DT_SOCK                           = 0xc
	DT_UNKNOWN                        = 0x0
	DT_WHT                            = 0xe
	ECHO                              = 0x8
	ECHOCTL                           = 0x40
	ECHOE                             = 0x2
	ECHOK                             = 0x4
	ECHOKE                            = 0x1
	ECHONL                            = 0x10
	ECHOPRT                           = 0x20
	EMUL_LINUX                        = 0x1
	EMUL_LINUX32                      = 0x5
	EMUL_MAXID                        = 0x6
	ETHERCAP_JUMBO_MTU                = 0x4
	ETHERCAP_VLAN_HWTAGGING           = 0x2
	ETHERCAP_VLAN_MTU                 = 0x1
	ETHERMIN                          = 0x2e
	ETHERMTU                          = 0x5dc
	ETHERMTU_JUMBO                    = 0x2328
	ETHERTYPE_8023                    = 0x4
	ETHERTYPE_AARP                    = 0x80f3
	ETHERTYPE_ACCTON                  = 0x8390
	ETHERTYPE_AEONIC                  = 0x8036
	ETHERTYPE_ALPHA                   = 0x814a
	ETHERTYPE_AMBER                   = 0x6008
	ETHERTYPE_AMOEBA                  = 0x8145
	ETHERTYPE_APOLLO                  = 0x80f7
	ETHERTYPE_APOLLODOMAIN            = 0x8019
	ETHERTYPE_APPLETALK               = 0x809b
	ETHERTYPE_APPLITEK                = 0x80c7
	ETHERTYPE_ARGONAUT                = 0x803a
	ETHERTYPE_ARP                     = 0x806
	ETHERTYPE_AT                      = 0x809b
	ETHERTYPE_ATALK                   = 0x809b
	ETHERTYPE_ATOMIC                  = 0x86df
	ETHERTYPE_ATT                     = 0x8069
	ETHERTYPE_ATTSTANFORD             = 0x8008
	ETHERTYPE_AUTOPHON                = 0x806a
	ETHERTYPE_AXIS                    = 0x8856
	ETHERTYPE_BCLOOP                  = 0x9003
	ETHERTYPE_BOFL                    = 0x8102
	ETHERTYPE_CABLETRON               = 0x7034
	ETHERTYPE_CHAOS                   = 0x804
	ETHERTYPE_COMDESIGN               = 0x806c
	ETHERTYPE_COMPUGRAPHIC            = 0x806d
	ETHERTYPE_COUNTERPOINT            = 0x8062
	ETHERTYPE_CRONUS                  = 0x8004
	ETHERTYPE_CRONUSVLN               = 0x8003
	ETHERTYPE_DCA                     = 0x1234
	ETHERTYPE_DDE                     = 0x807b
	ETHERTYPE_DEBNI                   = 0xaaaa
	ETHERTYPE_DECAM                   = 0x8048
	ETHERTYPE_DECCUST                 = 0x6006
	ETHERTYPE_DECDIAG                 = 0x6005
	ETHERTYPE_DECDNS                  = 0x803c
	ETHERTYPE_DECDTS                  = 0x803e
	ETHERTYPE_DECEXPER                = 0x6000
	ETHERTYPE_DECLAST                 = 0x8041
	ETHERTYPE_DECLTM                  = 0x803f
	ETHERTYPE_DECMUMPS                = 0x6009
	ETHERTYPE_DECNETBIOS              = 0x8040
	ETHERTYPE_DELTACON                = 0x86de
	ETHERTYPE_DIDDLE                  = 0x4321
	ETHERTYPE_DLOG1                   = 0x660
	ETHERTYPE_DLOG2                   = 0x661
	ETHERTYPE_DN                      = 0x6003
	ETHERTYPE_DOGFIGHT                = 0x1989
	ETHERTYPE_DSMD                    = 0x8039
	ETHERTYPE_ECMA                    = 0x803
	ETHERTYPE_ENCRYPT                 = 0x803d
	ETHERTYPE_ES                      = 0x805d
	ETHERTYPE_EXCELAN                 = 0x8010
	ETHERTYPE_EXPERDATA               = 0x8049
	ETHERTYPE_FLIP                    = 0x8146
	ETHERTYPE_FLOWCONTROL             = 0x8808
	ETHERTYPE_FRARP                   = 0x808
	ETHERTYPE_GENDYN                  = 0x8068
	ETHERTYPE_HAYES                   = 0x8130
	ETHERTYPE_HIPPI_FP                = 0x8180
	ETHERTYPE_HITACHI                 = 0x8820
	ETHERTYPE_HP                      = 0x8005
	ETHERTYPE_IEEEPUP                 = 0xa00
	ETHERTYPE_IEEEPUPAT               = 0xa01
	ETHERTYPE_IMLBL                   = 0x4c42
	ETHERTYPE_IMLBLDIAG               = 0x424c
	ETHERTYPE_IP                      = 0x800
	ETHERTYPE_IPAS                    = 0x876c
	ETHERTYPE_IPV6                    = 0x86dd
	ETHERTYPE_IPX                     = 0x8137
	ETHERTYPE_IPXNEW                  = 0x8037
	ETHERTYPE_KALPANA                 = 0x8582
	ETHERTYPE_LANBRIDGE               = 0x8038
	ETHERTYPE_LANPROBE                = 0x8888
	ETHERTYPE_LAT                     = 0x6004
	ETHERTYPE_LBACK                   = 0x9000
	ETHERTYPE_LITTLE                  = 0x8060
	ETHERTYPE_LOGICRAFT               = 0x8148
	ETHERTYPE_LOOPBACK                = 0x9000
	ETHERTYPE_MATRA                   = 0x807a
	ETHERTYPE_MAX                     = 0xffff
	ETHERTYPE_MERIT                   = 0x807c
	ETHERTYPE_MICP                    = 0x873a
	ETHERTYPE_MOPDL                   = 0x6001
	ETHERTYPE_MOPRC                   = 0x6002
	ETHERTYPE_MOTOROLA                = 0x818d
	ETHERTYPE_MPLS                    = 0x8847
	ETHERTYPE_MPLS_MCAST              = 0x8848
	ETHERTYPE_MUMPS                   = 0x813f
	ETHERTYPE_NBPCC                   = 0x3c04
	ETHERTYPE_NBPCLAIM                = 0x3c09
	ETHERTYPE_NBPCLREQ                = 0x3c05
	ETHERTYPE_NBPCLRSP                = 0x3c06
	ETHERTYPE_NBPCREQ                 = 0x3c02
	ETHERTYPE_NBPCRSP                 = 0x3c03
	ETHERTYPE_NBPDG                   = 0x3c07
	ETHERTYPE_NBPDGB                  = 0x3c08
	ETHERTYPE_NBPDLTE                 = 0x3c0a
	ETHERTYPE_NBPRAR                  = 0x3c0c
	ETHERTYPE_NBPRAS                  = 0x3c0b
	ETHERTYPE_NBPRST                  = 0x3c0d
	ETHERTYPE_NBPSCD                  = 0x3c01
	ETHERTYPE_NBPVCD                  = 0x3c00
	ETHERTYPE_NBS                     = 0x802
	ETHERTYPE_NCD                     = 0x8149
	ETHERTYPE_NESTAR                  = 0x8006
	ETHERTYPE_NETBEUI                 = 0x8191
	ETHERTYPE_NOVELL                  = 0x8138
	ETHERTYPE_NS                      = 0x600
	ETHERTYPE_NSAT                    = 0x601
	ETHERTYPE_NSCOMPAT                = 0x807
	ETHERTYPE_NTRAILER                = 0x10
	ETHERTYPE_OS9                     = 0x7007
	ETHERTYPE_OS9NET                  = 0x7009
	ETHERTYPE_PACER                   = 0x80c6
	ETHERTYPE_PAE                     = 0x888e
	ETHERTYPE_PCS                     = 0x4242
	ETHERTYPE_PLANNING                = 0x8044
	ETHERTYPE_PPP                     = 0x880b
	ETHERTYPE_PPPOE                   = 0x8864
	ETHERTYPE_PPPOEDISC               = 0x8863
	ETHERTYPE_PRIMENTS                = 0x7031
	ETHERTYPE_PUP                     = 0x200
	ETHERTYPE_PUPAT                   = 0x200
	ETHERTYPE_RACAL                   = 0x7030
	ETHERTYPE_RATIONAL                = 0x8150
	ETHERTYPE_RAWFR                   = 0x6559
	ETHERTYPE_RCL                     = 0x1995
	ETHERTYPE_RDP                     = 0x8739
	ETHERTYPE_RETIX                   = 0x80f2
	ETHERTYPE_REVARP                  = 0x8035
	ETHERTYPE_SCA                     = 0x6007
	ETHERTYPE_SECTRA                  = 0x86db
	ETHERTYPE_SECUREDATA              = 0x876d
	ETHERTYPE_SGITW                   = 0x817e
	ETHERTYPE_SG_BOUNCE               = 0x8016
	ETHERTYPE_SG_DIAG                 = 0x8013
	ETHERTYPE_SG_NETGAMES             = 0x8014
	ETHERTYPE_SG_RESV                 = 0x8015
	ETHERTYPE_SIMNET                  = 0x5208
	ETHERTYPE_SLOWPROTOCOLS           = 0x8809
	ETHERTYPE_SNA                     = 0x80d5
	ETHERTYPE_SNMP                    = 0x814c
	ETHERTYPE_SONIX                   = 0xfaf5
	ETHERTYPE_SPIDER                  = 0x809f
	ETHERTYPE_SPRITE                  = 0x500
	ETHERTYPE_STP                     = 0x8181
	ETHERTYPE_TALARIS                 = 0x812b
	ETHERTYPE_TALARISMC               = 0x852b
	ETHERTYPE_TCPCOMP                 = 0x876b
	ETHERTYPE_TCPSM                   = 0x9002
	ETHERTYPE_TEC                     = 0x814f
	ETHERTYPE_TIGAN                   = 0x802f
	ETHERTYPE_TRAIL                   = 0x1000
	ETHERTYPE_TRANSETHER              = 0x6558
	ETHERTYPE_TYMSHARE                = 0x802e
	ETHERTYPE_UBBST                   = 0x7005
	ETHERTYPE_UBDEBUG                 = 0x900
	ETHERTYPE_UBDIAGLOOP              = 0x7002
	ETHERTYPE_UBDL                    = 0x7000
	ETHERTYPE_UBNIU                   = 0x7001
	ETHERTYPE_UBNMC                   = 0x7003
	ETHERTYPE_VALID                   = 0x1600
	ETHERTYPE_VARIAN                  = 0x80dd
	ETHERTYPE_VAXELN                  = 0x803b
	ETHERTYPE_VEECO                   = 0x8067
	ETHERTYPE_VEXP                    = 0x805b
	ETHERTYPE_VGLAB                   = 0x8131
	ETHERTYPE_VINES                   = 0xbad
	ETHERTYPE_VINESECHO               = 0xbaf
	ETHERTYPE_VINESLOOP               = 0xbae
	ETHERTYPE_VITAL                   = 0xff00
	ETHERTYPE_VLAN                    = 0x8100
	ETHERTYPE_VLTLMAN                 = 0x8080
	ETHERTYPE_VPROD                   = 0x805c
	ETHERTYPE_VURESERVED              = 0x8147
	ETHERTYPE_WATERLOO                = 0x8130
	ETHERTYPE_WELLFLEET               = 0x8103
	ETHERTYPE_X25                     = 0x805
	ETHERTYPE_X75                     = 0x801
	ETHERTYPE_XNSSM                   = 0x9001
	ETHERTYPE_XTP                     = 0x817d
	ETHER_ADDR_LEN                    = 0x6
	ETHER_CRC_LEN                     = 0x4
	ETHER_CRC_POLY_BE                 = 0x4c11db6
	ETHER_CRC_POLY_LE                 = 0xedb88320
	ETHER_HDR_LEN                     = 0xe
	ETHER_MAX_LEN                     = 0x5ee
	ETHER_MAX_LEN_JUMBO               = 0x233a
	ETHER_MIN_LEN                     = 0x40
	ETHER_PPPOE_ENCAP_LEN             = 0x8
	ETHER_TYPE_LEN                    = 0x2
	ETHER_VLAN_ENCAP_LEN              = 0x4
	EVFILT_AIO                        = 0x2
	EVFILT_PROC                       = 0x4
	EVFILT_READ                       = 0x0
	EVFILT_SIGNAL                     = 0x5
	EVFILT_SYSCOUNT                   = 0x7
	EVFILT_TIMER                      = 0x6
	EVFILT_VNODE                      = 0x3
	EVFILT_WRITE                      = 0x1
	EV_ADD                            = 0x1
	EV_CLEAR                          = 0x20
	EV_DELETE                         = 0x2
	EV_DISABLE                        = 0x8
	EV_ENABLE                         = 0x4
	EV_EOF                            = 0x8000
	EV_ERROR                          = 0x4000
	EV_FLAG1                          = 0x2000
	EV_ONESHOT                        = 0x10
	EV_SYSFLAGS                       = 0xf000
	EXTA                              = 0x4b00
	EXTB                              = 0x9600
	EXTPROC                           = 0x800
	FD_CLOEXEC                        = 0x1
	FD_SETSIZE                        = 0x100
	FLUSHO                            = 0x800000
	F_CLOSEM                          = 0xa
	F_DUPFD                           = 0x0
	F_DUPFD_CLOEXEC                   = 0xc
	F_FSCTL                           = -0x80000000
	F_FSDIRMASK                       = 0x70000000
	F_FSIN                            = 0x10000000
	F_FSINOUT                         = 0x30000000
	F_FSOUT                           = 0x20000000
	F_FSPRIV                          = 0x8000
	F_FSVOID                          = 0x40000000
	F_GETFD                           = 0x1
	F_GETFL                           = 0x3
	F_GETLK                           = 0x7
	F_GETNOSIGPIPE                    = 0xd
	F_GETOWN                          = 0x5
	F_MAXFD                           = 0xb
	F_OK                              = 0x0
	F_PARAM_MASK                      = 0xfff
	F_PARAM_MAX                       = 0xfff
	F_RDLCK                           = 0x1
	F_SETFD                           = 0x2
	F_SETFL                           = 0x4
	F_SETLK                           = 0x8
	F_SETLKW                          = 0x9
	F_SETNOSIGPIPE                    = 0xe
	F_SETOWN                          = 0x6
	F_UNLCK                           = 0x2
	F_WRLCK                           = 0x3
	HUPCL                             = 0x4000
	HW_MACHINE                        = 0x1
	ICANON                            = 0x100
	ICMP6_FILTER                      = 0x12
	ICRNL                             = 0x100
	IEXTEN                            = 0x400
	IFAN_ARRIVAL                      = 0x0
	IFAN_DEPARTURE                    = 0x1
	IFA_ROUTE                         = 0x1
	IFF_ALLMULTI                      = 0x200
	IFF_BROADCAST                     = 0x2
	IFF_CANTCHANGE                    = 0x8f52
	IFF_DEBUG                         = 0x4
	IFF_LINK0                         = 0x1000
	IFF_LINK1                         = 0x2000
	IFF_LINK2                         = 0x4000
	IFF_LOOPBACK                      = 0x8
	IFF_MULTICAST                     = 0x8000
	IFF_NOARP                         = 0x80
	IFF_NOTRAILERS                    = 0x20
	IFF_OACTIVE                       = 0x400
	IFF_POINTOPOINT                   = 0x10
	IFF_PROMISC                       = 0x100
	IFF_RUNNING                       = 0x40
	IFF_SIMPLEX                       = 0x800
	IFF_UP                            = 0x1
	IFNAMSIZ                          = 0x10
	IFT_1822                          = 0x2
	IFT_A12MPPSWITCH                  = 0x82
	IFT_AAL2                          = 0xbb
	IFT_AAL5                          = 0x31
	IFT_ADSL                          = 0x5e
	IFT_AFLANE8023                    = 0x3b
	IFT_AFLANE8025                    = 0x3c
	IFT_ARAP                          = 0x58
	IFT_ARCNET                        = 0x23
	IFT_ARCNETPLUS                    = 0x24
	IFT_ASYNC                         = 0x54
	IFT_ATM                           = 0x25
	IFT_ATMDXI                        = 0x69
	IFT_ATMFUNI                       = 0x6a
	IFT_ATMIMA                        = 0x6b
	IFT_ATMLOGICAL                    = 0x50
	IFT_ATMRADIO                      = 0xbd
	IFT_ATMSUBINTERFACE               = 0x86
	IFT_ATMVCIENDPT                   = 0xc2
	IFT_ATMVIRTUAL                    = 0x95
	IFT_BGPPOLICYACCOUNTING           = 0xa2
	IFT_BRIDGE                        = 0xd1
	IFT_BSC                           = 0x53
	IFT_CARP                          = 0xf8
	IFT_CCTEMUL                       = 0x3d
	IFT_CEPT                          = 0x13
	IFT_CES                           = 0x85
	IFT_CHANNEL                       = 0x46
	IFT_CNR                           = 0x55
	IFT_COFFEE                        = 0x84
	IFT_COMPOSITELINK                 = 0x9b
	IFT_DCN                           = 0x8d
	IFT_DIGITALPOWERLINE              = 0x8a
	IFT_DIGITALWRAPPEROVERHEADCHANNEL = 0xba
	IFT_DLSW                          = 0x4a
	IFT_DOCSCABLEDOWNSTREAM           = 0x80
	IFT_DOCSCABLEMACLAYER             = 0x7f
	IFT_DOCSCABLEUPSTREAM             = 0x81
	IFT_DOCSCABLEUPSTREAMCHANNEL      = 0xcd
	IFT_DS0                           = 0x51
	IFT_DS0BUNDLE                     = 0x52
	IFT_DS1FDL                        = 0xaa
	IFT_DS3                           = 0x1e
	IFT_DTM                           = 0x8c
	IFT_DVBASILN                      = 0xac
	IFT_DVBASIOUT                     = 0xad
	IFT_DVBRCCDOWNSTREAM              = 0x93
	IFT_DVBRCCMACLAYER                = 0x92
	IFT_DVBRCCUPSTREAM                = 0x94
	IFT_ECONET                        = 0xce
	IFT_EON                           = 0x19
	IFT_EPLRS                         = 0x57
	IFT_ESCON                         = 0x49
	IFT_ETHER                         = 0x6
	IFT_FAITH                         = 0xf2
	IFT_FAST                          = 0x7d
	IFT_FASTETHER                     = 0x3e
	IFT_FASTETHERFX                   = 0x45
	IFT_FDDI                          = 0xf
	IFT_FIBRECHANNEL                  = 0x38
	IFT_FRAMERELAYINTERCONNECT        = 0x3a
	IFT_FRAMERELAYMPI                 = 0x5c
	IFT_FRDLCIENDPT                   = 0xc1
	IFT_FRELAY                        = 0x20
	IFT_FRELAYDCE                     = 0x2c
	IFT_FRF16MFRBUNDLE                = 0xa3
	IFT_FRFORWARD                     = 0x9e
	IFT_G703AT2MB                     = 0x43
	IFT_G703AT64K                     = 0x42
	IFT_GIF                           = 0xf0
	IFT_GIGABITETHERNET               = 0x75
	IFT_GR303IDT                      = 0xb2
	IFT_GR303RDT                      = 0xb1
	IFT_H323GATEKEEPER                = 0xa4
	IFT_H323PROXY                     = 0xa5
	IFT_HDH1822                       = 0x3
	IFT_HDLC                          = 0x76
	IFT_HDSL2                         = 0xa8
	IFT_HIPERLAN2                     = 0xb7
	IFT_HIPPI                         = 0x2f
	IFT_HIPPIINTERFACE                = 0x39
	IFT_HOSTPAD                       = 0x5a
	IFT_HSSI                          = 0x2e
	IFT_HY                            = 0xe
	IFT_IBM370PARCHAN                 = 0x48
	IFT_IDSL                          = 0x9a
	IFT_IEEE1394                      = 0x90
	IFT_IEEE80211                     = 0x47
	IFT_IEEE80212                     = 0x37
	IFT_IEEE8023ADLAG                 = 0xa1
	IFT_IFGSN                         = 0x91
	IFT_IMT                           = 0xbe
	IFT_INFINIBAND                    = 0xc7
	IFT_INTERLEAVE                    = 0x7c
	IFT_IP                            = 0x7e
	IFT_IPFORWARD                     = 0x8e
	IFT_IPOVERATM                     = 0x72
	IFT_IPOVERCDLC                    = 0x6d
	IFT_IPOVERCLAW                    = 0x6e
	IFT_IPSWITCH                      = 0x4e
	IFT_ISDN                          = 0x3f
	IFT_ISDNBASIC                     = 0x14
	IFT_ISDNPRIMARY                   = 0x15
	IFT_ISDNS                         = 0x4b
	IFT_ISDNU                         = 0x4c
	IFT_ISO88022LLC                   = 0x29
	IFT_ISO88023                      = 0x7
	IFT_ISO88024                      = 0x8
	IFT_ISO88025                      = 0x9
	IFT_ISO88025CRFPINT               = 0x62
	IFT_ISO88025DTR                   = 0x56
	IFT_ISO88025FIBER                 = 0x73
	IFT_ISO88026                      = 0xa
	IFT_ISUP                          = 0xb3
	IFT_L2VLAN                        = 0x87
	IFT_L3IPVLAN                      = 0x88
	IFT_L3IPXVLAN                     = 0x89
	IFT_LAPB                          = 0x10
	IFT_LAPD                          = 0x4d
	IFT_LAPF                          = 0x77
	IFT_LINEGROUP                     = 0xd2
	IFT_LOCALTALK                     = 0x2a
	IFT_LOOP                          = 0x18
	IFT_MEDIAMAILOVERIP               = 0x8b
	IFT_MFSIGLINK                     = 0xa7
	IFT_MIOX25                        = 0x26
	IFT_MODEM                         = 0x30
	IFT_MPC                           = 0x71
	