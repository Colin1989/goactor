package goactor

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
	}{
		// TODO: Add test cases.
		{name: "test-all", opts: []Option{WithDebug(), WithSeverMode(Standalone)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApp(tt.opts...)
			if got == nil {
				return
			}
		})
	}
}
