package p2p_test

import (
	"testing"

	"github.com/machine-moon/gogo/p2p"
)

func TestNewTCPTransport(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		listenAddr string
		want       p2p.Transport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p2p.NewTCPTransport(tt.listenAddr)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewTCPTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTCPTransport(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		listenAddr string
		want       p2p.Transport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := p2p.NewTCPTransport(tt.listenAddr)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewTCPTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
