package goget

import (
	"testing"

	"github.com/newacorn/depend2"
)

func TestGet(t *testing.T) {
	depend2.Depend2()
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get()
		})
	}
}
