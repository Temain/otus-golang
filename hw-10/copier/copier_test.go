package copier

import (
	"io/ioutil"
	"testing"
)

func TestCopyAll(t *testing.T) {
	const (
		from    = "from.txt"
		fromStr = "from file data"
		to      = "to_all.txt"
		limit   = 0
		offset  = 0
	)

	err := prepareFile(from, fromStr)
	if err != nil {
		t.Fatalf("error on prepare source file")
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}

	targetBytes, err := ioutil.ReadFile(to)
	if err != nil {
		t.Fatalf("error on read target file")
	}
	targetStr := string(targetBytes)
	if targetStr != fromStr {
		t.Fatalf("bad data in target file after copy")
	}
}

func TestCopyWithLimit(t *testing.T) {
	const (
		from    = "from.txt"
		fromStr = "from file data"
		to      = "to_limit.txt"
		limit   = 6
		offset  = 0
	)

	err := prepareFile(from, fromStr)
	if err != nil {
		t.Fatalf("error on prepare source file")
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}

	targetBytes, err := ioutil.ReadFile(to)
	if err != nil {
		t.Fatalf("error on read target file")
	}
	targetStr := string(targetBytes)
	if targetStr != "from f" {
		t.Fatalf("bad data in target file after copy")
	}
}

func TestCopyWithOffset(t *testing.T) {
	const (
		from    = "from.txt"
		fromStr = "from file data"
		to      = "to_offset.txt"
		limit   = 0
		offset  = 2
	)

	err := prepareFile(from, fromStr)
	if err != nil {
		t.Fatalf("error on prepare source file")
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}

	targetBytes, err := ioutil.ReadFile(to)
	if err != nil {
		t.Fatalf("error on read target file")
	}
	targetStr := string(targetBytes)
	if targetStr != "om file data" {
		t.Fatalf("bad data in target file after copy")
	}
}

func TestCopyWithLimitAndOffset(t *testing.T) {
	const (
		from    = "from.txt"
		fromStr = "from file data"
		to      = "to_limit_offset.txt"
		limit   = 4
		offset  = 2
	)

	err := prepareFile(from, fromStr)
	if err != nil {
		t.Fatalf("error on prepare source file")
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}

	targetBytes, err := ioutil.ReadFile(to)
	if err != nil {
		t.Fatalf("error on read target file")
	}
	targetStr := string(targetBytes)
	if targetStr != "om f" {
		t.Fatalf("bad data in target file after copy")
	}
}

func prepareFile(name string, data string) error {
	fileBytes := []byte(data)
	err := ioutil.WriteFile(name, fileBytes, 0644)
	return err
}
