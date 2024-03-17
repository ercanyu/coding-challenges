package handler

import (
	"testing"
)

func TestHandleWcCommandWithOptionC(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfBytes := HandleWcCommand(filename, "c")

	// then
	if numberOfBytes != 342190 {
		t.Errorf("Expected 342190, got %d", numberOfBytes)
	}
}

func TestHandleWcCommandWithOptionL(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := HandleWcCommand(filename, "l")

	// then
	if numberOfLines != 7145 {
		t.Errorf("Expected 7145, got %d", numberOfLines)
	}
}

func TestHandleWcCommandWithOptionW(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := HandleWcCommand(filename, "w")

	// then
	if numberOfLines != 58164 {
		t.Errorf("Expected 58164, got %d", numberOfLines)
	}
}

func TestHandleWcCommandWithOptionM(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := HandleWcCommand(filename, "m")

	// then
	if numberOfLines != 339292 {
		t.Errorf("Expected 339292, got %d", numberOfLines)
	}
}
