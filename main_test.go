package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	t.Run("check valid path", func(t *testing.T) {
		result := tree(".", "")
		got := result
		if got != nil {
			t.Errorf("Did not exit without an error wanted: %v, got %v", nil, got)
		}

	})

	t.Run("Check error on path that doesn't exist", func(t *testing.T) {
		got := tree("path that doesn't exist", "")
		if got == nil {
			t.Errorf("did not exit with incorrect path, got: %q, want: %v", got, nil)
		}

	})
}
