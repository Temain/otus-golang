package copier

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestCopyAll(t *testing.T) {
	const (
		from   = "from.txt"
		to     = "to.txt"
		limit  = 6
		offset = 1
	)

	fromStr := "from file data"
	err := prepareFile(from, fromStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}
}

func TestCopyWithZeroLimit(t *testing.T) {
	const (
		from   = "from.txt"
		to     = "to.txt"
		limit  = 10000
		offset = 1
	)

	fromStr := strings.Repeat("from file data", 100000)
	err := prepareFile(from, fromStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}
}

func prepareFile(name string, data string) error {
	fileBytes := []byte(data)
	err := ioutil.WriteFile(name, fileBytes, 0644)
	return err
}
