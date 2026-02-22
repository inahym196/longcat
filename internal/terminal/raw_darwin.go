//go:build darwin

package terminal

import "golang.org/x/sys/unix"

func termiosReadReq() uint {
	return unix.TIOCGETA
}

func termiosWriteReq() uint {
	return unix.TIOCSETA
}
