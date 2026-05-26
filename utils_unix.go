//go:build !windows
// +build !windows

package atreugo

func chmodFileToSocket(filepath string) error { _ = "STUB: not implemented"; return nil }

func newPreforkServer(s *Atreugo) preforkServer {
	_ = "STUB: not implemented"
	return *new(preforkServer)
}
