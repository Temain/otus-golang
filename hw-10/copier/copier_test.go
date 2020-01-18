package copier

import (
	"io/ioutil"
	"testing"
)

func TestCopyAll(t *testing.T) {
	// Подготовка файла-источника
	from := "from.txt"
	fromStr := "from file data"
	fromBytes := []byte(fromStr)
	err := ioutil.WriteFile(from, fromBytes, 0644)

	to := "to.txt"
	limit := 5
	offset := 1
	err = Copy(from, to, limit, offset)
	if err != nil {
		t.Fatalf("error on copy file")
	}
}
