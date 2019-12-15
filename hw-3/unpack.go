package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Unpack выполняет распаковку строки, содержащую повторяющиеся символы.
func Unpack(s string) (r string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return r, errors.New("некорректная строка")
	}

	var prev rune
	var escaped bool
	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			mult := int(char - '0')
			r = r + strings.Repeat(string(prev), mult-1)
		} else {
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				r = r + string(char)
			}
			prev = char
		}
	}

	return r, err
}
