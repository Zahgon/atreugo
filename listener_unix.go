//go:build !windows
// +build !windows

package atreugo

import (
	"net"
)

func (s *Atreugo) getListener() (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// nolint:wrapcheck
