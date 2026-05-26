//go:build windows
// +build windows

package atreugo

// ListenAndServe serves requests from the given network and address in the atreugo configuration.
//
// Pass custom listener to Serve/ServeGracefully if you want to use it.
func (s *Atreugo) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

// nolint:wrapcheck
