package runner3

import (
	"errors"
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	tasks := []func() error{
		func() error { fmt.Println("task #1"); return nil },
		func() error { fmt.Println("task #2"); return nil },
		func() error { fmt.Println("task #3"); return nil },
		func() error { fmt.Println("task #4"); return nil },
		func() error { fmt.Println("task #5"); return nil },
		func() error { fmt.Println("task #6"); return nil },
		func() error { fmt.Println("task #7"); return nil },
		func() error { fmt.Println("task #8"); return nil },
		func() error { fmt.Println("task #9"); return nil },
		func() error { fmt.Println("task #10"); return nil },
	}
	_, s, e := Run(tasks, 2, 2)
	ln := len(tasks)
	if int(s) != ln {
		t.Fatalf("bad success tasks %d expected %d", s, ln)
	}
	if e != 0 {
		t.Fatalf("bad error tasks expected 0")
	}
}

func TestErrors(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("task #1 with error")
			return errors.New("error #1")
		},
		func() error {
			fmt.Println("task #2 with error")
			return errors.New("error #2")
		},
		func() error { fmt.Println("task #3"); return nil },
		func() error {
			fmt.Println("task #4 with error")
			return errors.New("error #4")
		},
		func() error { fmt.Println("task #5"); return nil },
		func() error { fmt.Println("task #6"); return nil },
		func() error { fmt.Println("task #7"); return nil },
		func() error { fmt.Println("task #8"); return nil },
		func() error { fmt.Println("task #9"); return nil },
		func() error { fmt.Println("task #10"); return nil },
	}
	n, m := 2, 2
	err, s, e := Run(tasks, n, m)
	if e+s > n+m {
		t.Fatalf("bad completed tasks %d expected max %d", e+s, n+m)
	}
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}
