
#!/usr/bin/env perl

# Copyright 2011 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

#
# Parse the header files for OpenBSD and generate a Go usable sysctl MIB.
#
# Build a MIB with each entry being an array containing the level, type and
# a hash that will contain additional entries if the current entry is a node.
# We then walk this MIB and create a flattened sysctl name to OID hash.
#

use strict;

if($ENV{'GOARCH'} eq "" || $ENV{'GOOS'} eq "") {
	print STDERR "GOARCH or GOOS not defined in environment\n";
	exit 1;
}

my $debug = 0;
my %ctls = ();

my @headers = qw (
	sys/sysctl.h
	sys/socket.h
	sys/tty.h
	sys/malloc.h
	sys/mount.h
	sys/namei.h
	sys/sem.h
	sys/shm.h
	sys/vmmeter.h
	uvm/uvm_param.h
	uvm/uvm_swap_encrypt.h
	ddb/db_var.h
	net/if.h
	net/if_pfsync.h
	net/pipex.h
	netinet/in.h
	netinet/icmp_var.h
	netinet/igmp_var.h
	netinet/ip_ah.h
	netinet/ip_carp.h
	netinet/ip_divert.h
	netinet/ip_esp.h
	netinet/ip_ether.h
	netinet/ip_gre.h
	netinet/ip_ipcomp.h
	netinet/ip_ipip.h
	netinet/pim_var.h
	netinet/tcp_var.h
	netinet/udp_var.h
	netinet6/in6.h
	netinet6/ip6_divert.h
	netinet6/pim6_var.h
	netinet/icmp6.h
	netmpls/mpls.h
);

my @ctls = qw (
	kern
	vm
	fs
	net
	#debug				# Special handling required
	hw
	#machdep			# Arch specific
	user
	ddb
	#vfs				# Special handling required
	fs.posix
	kern.forkstat
	kern.intrcnt
	kern.malloc
	kern.nchstats
	kern.seminfo
	kern.shminfo
	kern.timecounter
	kern.tty
	kern.watchdog
	net.bpf
	net.ifq
	net.inet
	net.inet.ah
	net.inet.carp
	net.inet.divert
	net.inet.esp
	net.inet.etherip
	net.inet.gre
	net.inet.icmp
	net.inet.igmp
	net.inet.ip
	net.inet.ip.ifq
	net.inet.ipcomp
	net.inet.ipip
	net.inet.mobileip
	net.inet.pfsync
	net.inet.pim
	net.inet.tcp
	net.inet.udp
	net.inet6
	net.inet6.divert
	net.inet6.ip6
	net.inet6.icmp6
	net.inet6.pim6
	net.inet6.tcp6
	net.inet6.udp6
	net.mpls
	net.mpls.ifq
	net.key
	net.pflow
	net.pfsync
	net.pipex
	net.rt
	vm.swapencrypt
	#vfsgenctl			# Special handling required
);

# Node name "fixups"
my %ctl_map = (
	"ipproto" => "net.inet",
	"net.inet.ipproto" => "net.inet",
	"net.inet6.ipv6proto" => "net.inet6",
	"net.inet6.ipv6" => "net.inet6.ip6",
	"net.inet.icmpv6" => "net.inet6.icmp6",
	"net.inet6.divert6" => "net.inet6.divert",
	"net.inet6.tcp6" => "net.inet.tcp",
	"net.inet6.udp6" => "net.inet.udp",
	"mpls" => "net.mpls",
	"swpenc" => "vm.swapencrypt"
);

# Node mappings
my %node_map = (
	"net.inet.ip.ifq" => "net.ifq",
	"net.inet.pfsync" => "net.pfsync",
	"net.mpls.ifq" => "net.ifq"