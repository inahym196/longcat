//go:build linux

package terminal

import "golang.org/x/sys/unix"

func termiosReadReq() uint {
	return unix.TCGETS
}

func termiosWriteReq() uint {
	return unix.TCSETS
}
