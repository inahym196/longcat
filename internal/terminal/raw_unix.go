//go:build darwin || linux

package terminal

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func EnterRawMode(file *os.File) (func() error, error) {
	if file == nil {
		return nil, fmt.Errorf("file is nil")
	}
	if !isTerminal(file) {
		return func() error { return nil }, nil
	}

	fd := file.Fd()
	state, err := unix.IoctlGetTermios(int(fd), termiosReadReq())
	if err != nil {
		return nil, fmt.Errorf("get termios: %w", err)
	}
	old := *state

	state.Lflag &^= unix.ICANON | unix.ECHO
	state.Cc[unix.VMIN] = 1
	state.Cc[unix.VTIME] = 0

	if err := unix.IoctlSetTermios(int(fd), termiosWriteReq(), state); err != nil {
		return nil, fmt.Errorf("set termios(raw): %w", err)
	}

	return func() error {
		restore := old
		if err := unix.IoctlSetTermios(int(fd), termiosWriteReq(), &restore); err != nil {
			return fmt.Errorf("restore termios: %w", err)
		}
		return nil
	}, nil
}

func isTerminal(file *os.File) bool {
	fi, err := file.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}
