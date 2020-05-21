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
my $err