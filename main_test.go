package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, CI/CD World!"
	if expected != "Hello, CI/CD World!" {
		t.Errorf("expected %s, got %s", expected, "Hello, CI/CD World!")
	}
}
