package top

import (
	"sort"
	"strings"
)

func Top10(s string) (res []string) {
	s = clear(s)
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	p := make(PairList, len(counts))

	i := 0
	for k, v := range counts {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))

	for j := 0; j < len(p) && j < 10; j++ {
		res = append(res, p[j].Key)
	}

	return res
}

func clear(s string) string {
	s = strings.ToLower(s)
	replacer := strings.NewReplacer(",", "", ".", "", ";", "", "-", "", "{", "", "}", "")
	return replacer.Replace(s)
}
