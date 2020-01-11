package runner

import (
	"errors"
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("task #1")
			return nil
		},
		func() error {
			fmt.Println("task #2")
			return nil
		},
		func() error {
			fmt.Println("task #3")
			return nil
		},
		func() error {
			fmt.Println("task #4")
			return nil
		},
		func() error {
			fmt.Println("task #5")
			return nil
		},
		func() error {
			fmt.Println("task #6")
			return nil
		},
	}
	_, s, e := Run(tasks, 2, 2)
	ln := len(tasks)
	if s != ln {
		t.Fatalf("bad success tasks #{s} expected #{ln}")
	}
	if e != 0 {
		t.Fatalf("bad error tasks expected 0")
	}
}

func TestErrors(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("task #1")
			return errors.New("error #1")
		},
		func() error {
			fmt.Println("task #2")
			return errors.New("error #2")
		},
		func() error {
			fmt.Println("task #3")
			return errors.New("error #3")
		},
		func() error {
			fmt.Println("task #4")
			return nil
		},
		func() error {
			fmt.Println("task #5")
			return nil
		},
		func() error {
			fmt.Println("task #6")
			return nil
		},
	}
	n, m := 2, 2
	err, _, e := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
	if e != m {
		t.Fatalf("bad error tasks #{e} expected #{m}")
	}
}
