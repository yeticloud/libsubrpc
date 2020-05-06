package subrpc

import "testing"

// TestNewManager func
func TestNewManager(t *testing.T) {
	m, err := NewManager("/tmp/")
	if err != nil {
		t.Error(err)
	}
	if m.SockPath == "" {
		t.Error("SockPath not set")
	}
	if len(m.Procs) != 0 {
		t.Error("Procs initialized with contents")
	}
	if m.OutBuffer == nil || m.ErrBuffer == nil {
		t.Error("buffers must be instantiated")
	}
}
