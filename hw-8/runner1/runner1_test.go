package runner1

import (
	"errors"
	"fmt"
	"testing"
)

func TestAllSuccess(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}
	err := Run(tasks, 2, 2)
	if err != nil {
		t.Fatalf("bad result, should be all success")
	}
}

func TestFirstErrors(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return errors.New("task with error") })
	}
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}
	n, m := 2, 2
	err := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}

func TestLastErrors(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return errors.New("task with error") })
	}

	n, m := 2, 2
	err := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}

func TestRandomErrors(t *testing.T) {
	tasks := []func() error{
		func() error {
			fmt.Println("task #1 with error")
			return errors.New("error #1")
		},
		func() error { fmt.Println("task #1"); return nil },
		func() error { fmt.Println("task #2"); return nil },
		func() error { fmt.Println("task #3"); return nil },
		func() error { fmt.Println("task #4"); return nil },
		func() error { fmt.Println("task #5"); return nil },
		func() error { fmt.Println("task #6"); return nil },
		func() error {
			fmt.Println("task #4 with error")
			return errors.New("error #2")
		},
		func() error { fmt.Println("task #7"); return nil },
		func() error { fmt.Println("task #8"); return nil },
		func() error { fmt.Println("task #9"); return nil },
		func() error {
			fmt.Println("task #4 with error")
			return errors.New("error #3")
		},
		func() error { fmt.Println("task #10"); return nil },
		func() error { fmt.Println("task #11"); return nil },
		func() error { fmt.Println("task #12"); return nil },
		func() error {
			fmt.Println("task #4 with error")
			return errors.New("error #4")
		},
		func() error {
			fmt.Println("task #4 with error")
			return errors.New("error #5")
		},
		func() error { fmt.Println("task #12"); return nil },
		func() error { fmt.Println("task #13"); return nil },
		func() error { fmt.Println("task #14"); return nil },
	}

	n, m := 2, 2
	err := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}

func TestNegativeNParam(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}

	n, m := -10, 10
	err := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}

func TestNegativeMParam(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}

	n, m := 10, -10
	err := Run(tasks, n, m)
	if err == nil {
		t.Fatalf("bad result expected error")
	}
}

func TestTasksCountLessThenN(t *testing.T) {
	var tasks []func() error
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error { return nil })
	}
	n, m := 10, 2
	err := Run(tasks, n, m)
	if err != nil {
		t.Fatalf("bad result, %v", err)
	}
}

func TestEmptyTasks(t *testing.T) {
	var tasks []func() error
	n, m := 10, 10
	err := Run(tasks, n, m)
	if err != nil {
		t.Fatalf("bad result, %v", err)
	}
}
