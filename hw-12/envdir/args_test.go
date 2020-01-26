package envdir

import (
	"testing"
)

func TestDirArg(t *testing.T) {
	args := []string{"/home", "cmd"}
	_, _, err := ParseArgs(args)
	if err != nil {
		t.Fatalf("bad result when dir passed")
	}
}

func TestIsWrongDirArg(t *testing.T) {
	args := []string{"dir", "cmd"}
	_, _, err := ParseArgs(args)
	if err == nil {
		t.Fatalf("bad result when not dir passed, should be error")
	}
}

func TestEmptySliceArgs(t *testing.T) {
	var args []string
	_, _, err := ParseArgs(args)
	if err == nil {
		t.Fatalf("bad result on empty args, should be error")
	}
}

func TestEmptyOneArg(t *testing.T) {
	args := []string{""}
	_, _, err := ParseArgs(args)
	if err == nil {
		t.Fatalf("bad result on empty one arg, should be error")
	}
}

func TestEmptyTwoArgs(t *testing.T) {
	args := []string{"", ""}
	_, _, err := ParseArgs(args)
	if err == nil {
		t.Fatalf("bad result on empty two args, should be error")
	}
}
