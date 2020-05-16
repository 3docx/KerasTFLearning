// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd

package unix

import (
	"errors"
	"fmt"
)

// Go implementation of C mostly found in /usr/src/sys/kern/subr_capability.c

const (
	// This is the version of CapRights this package understands. See C implementation for parallels.
	capRightsGoVersion = CAP_RIGHTS_VERSION_00
	capArSizeMin       = CAP_RIGHTS_VERSION_00 + 2
	capArSizeMax       = capRightsGoVersion + 2
)

var (
	bit2idx = []int{
		-1, 0, 1, -1, 2, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1,
		4, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}
)

func capidxbit(right uint64) int {
	return int((right >> 57) & 0x1f)
}

func rightToIndex(right uint64) (int, error) {
	idx := capidxbit(right)
	if idx < 0 || idx >= len(bit2idx) {
		return -2, fmt.Errorf("index for right 0x%x out of range", right)
	}
	return bit2idx[idx], nil
}

func caprver(right uint64) int {
	return int(right >> 62)
}

func capver(rights *CapRights) int {
	return caprver(rights.Rights[0])
}

func caparsize(rights *CapRights) int {
	return capver(rights) + 2
}

// CapRightsSet sets the permissions in setrights in rights.
func CapRightsSet(rights *CapRights, setrights []uint64) error {
	// This is essentially a copy of cap_rights_vset()
	