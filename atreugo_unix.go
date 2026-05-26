//go:build !windows
// +build !windows

package atreugo

import (
	"net"
)

// ServeGracefully serves incoming connections from the given listener with graceful shutdown
//
// It's blocked until the given listener returns permanent error.
func (s *Atreugo) ServeGracefully(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

// ListenAndServe serves requests from the given network and address in the atreugo configuration.
//
// Pass custom listener to Serve/ServeGracefully if you want to use it.
func (s *Atreugo) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

// nolint:wrapcheck
