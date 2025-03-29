package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello, World!" {
		t.Error("Expected 'Hello, World!', got ", out.Text)
	}

	test.Type(in, "Troy")
	if out.Text != "Hello Troy!" {
		t.Error("Expected 'Hello, Troy', got ", out.Text)
	}
}
