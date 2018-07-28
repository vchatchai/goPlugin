package main

import (
	"testing"

	"github.com/vchatchai/goPlugin/plugins"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPlugin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugins.Plugin("thai")
			plugins.Plugin("eng")
		})
	}
}
