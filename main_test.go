package main

import (
	"log"
	"testing"
)

func TestTree(t *testing.T) {
	result := tree(".", "")
	got := result
	if got != nil {
		log.Fatalf("Did not exit without an error wanted: %v, got %v", nil, got)
	}

}
