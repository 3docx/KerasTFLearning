#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This program reads a file containing function prototypes
# (like syscall_darwin.go) and generates system call bodies.
# The prototypes are marked by lines beginning with "//sys"
# and read like func declarations if //sys is replaced by func, but:
#	* The parameter lists must give a name for each argument.
#	  This includes return parameters.
#	* The parameter lists must give a type for each argument:
#	  the (x, y, z int) shorthand is not allowed.
#	* If the return parameter is an error number, it must be named errno.

# A line beginning with //sysnb is like //sys, except that the
# goroutine will not be suspended during the execution of the system
# call.  This must only be used for system calls which can never
# block, as otherwise the system call could cause all goroutines to
# hang.

use strict;

my $cmdline = "mksyscall.pl " . join(' ', @ARGV);
my $errors = 0;
my $_32bit = "";
my $plan9 = 0;
my $openbsd = 0;
my $netbsd = 0;
my $dragonfly = 0;
my $arm = 0; # 64-bit value should use (even, odd)-pair
my $tags = "";  # build tags

if($ARGV[0] eq "-b32") {
	$_32bit = "big-endian";
	shift;
} elsif($ARGV[0] eq "-l32") {
	$_32bit = "little-endian";
	shift;
}
if($ARGV[0] eq "-plan9") {
	$plan9 = 1;
	shift;
}
if($ARGV[0] eq "-openbsd") {
	$openbsd = 1;
	shift;
}
if($ARGV[0] eq "-netbsd") {
	$netbsd = 1;
	shift;
}
if($ARGV[0] eq "-dragonfly") {
	$dragonfly = 1;
	shift;
}
if($ARGV[0] eq "-arm") {
	$arm = 1;
	shift;
}
if($ARGV[0] eq "-tags") {
	shift;
	$tags = $ARGV[0];
	shift;
}

if($ARGV[0] =~ /^-/) {
	print STDERR "usage: mksyscall.pl [-b32 | -l32] [-tags x,y] [file ...]\n";
	exit 1;
}

# Check that we are using the new build system if we should
if($ENV{'GOOS'} eq "linux" && $ENV{'GOARCH'} ne "sparc64") {
	if($ENV{'GOLANG_SYS_BUILD'} ne "docker") {
		print STDERR "In the new build system, mksyscall should not be called directly.\n";
		print STDERR "See README.md\n";
		exit 1;
	}
}


sub parseparamlist($) {
	my ($list) = @_;
	$list =~ s/^\s*//;
	$list =~ s/\s*$//;
	if($list eq "") {
		return ();
	}
	return split(/\s*,\s*/, $list);
}

sub parseparam($) {
	my ($p) = @_;
	if($p !~ /^(\S*) (\S*)$/) {
		print STDERR "$ARGV:$.: malformed parameter: $p\n";
		$errors = 1;
		return ("xx", "int");
	}
	return ($1, $2);
}

my $text = "";
while(<>) {
	chomp;
	s/\s+/ /g;
	s/^\s+//;
	s/\s+$//;
	my $nonblock = /^\/\/sysnb /;
	next if !/^\/\/sys / && !$nonblock;

	# Line must be of the form
	#	func Open(path string, mode int, perm int) (fd int, errno error)
	# Split into name, in params, out params.
	if(!/^\/\/sys(nb)? (\w+)\(([^()]*)\)\s*(?:\(([^()]+)\))?\s*(?:=\s*((?i)SYS_[A-Z0-9_]+))?$/) {
		print STDERR "$ARGV:$.: malformed //sys declaration\n";
		$errors = 1;
		next;
	}
	my ($func, $in, $out, $sysname) = ($2, $3, $4, $5);

	# Split argument lists on comma.
	my @in = parseparamlist($in);
	my @out = parseparamlist($out);

	# Try in v