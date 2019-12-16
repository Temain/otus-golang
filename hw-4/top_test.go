package top

import (
	"io/ioutil"
	"testing"
)

func TestTop10(t *testing.T) {
	data, err := ioutil.ReadFile("book.txt")
	if err != nil {
		t.Fatal("Read file error.")
	}

	text := string(data)
	e := []string{"the", "a", "to", "you", "of", "is", "we", "and", "go", "it"}
	r := Top10(text)
	if !testEq(r, e) {
		t.Fatalf("Wrong top 10 words: got %v expected %v", r, e)
	}
}

func testEq(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
