package top

import (
	"sort"
	"strings"
	"unicode"
)

// Top10 возвращает 10 самых частовстречающихся в тексте слов.
func Top10(s string) []string {
	words := prepare(s)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	p := make(PairList, 0, len(counts))
	for k, v := range counts {
		p = append(p, Pair{k, v})
	}

	sort.Sort(sort.Reverse(p))

	res := make([]string, 0, 10)
	for j, pair := range p {
		if j >= 10 {
			break
		}
		res = append(res, pair.Key)
	}

	return res
}

func prepare(s string) []string {
	s = strings.ToLower(s)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(s, f)
}
