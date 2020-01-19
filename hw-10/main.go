package main

import (
	"flag"
	"github.com/Temain/otus-golang/hw-10/copier"
	"log"
	"strings"
)

var (
	from   string
	to     string
	limit  int
	offset int
)

func init() {
	flag.StringVar(&from, "from", "", "source file")
	flag.StringVar(&to, "to", "", "target file")
	flag.IntVar(&limit, "limit", 0, "limit in source file")
	flag.IntVar(&offset, "offset", 0, "offset in source file")
}

func main() {
	flag.Parse()

	if len(strings.TrimSpace(from)) == 0 {
		log.Fatal("empty -from arg, see --help")
	}
	if len(strings.TrimSpace(to)) == 0 {
		log.Fatal("empty -to arg, see --help")
	}

	err := hw_10.Copy(from, to, limit, offset)
	if err != nil {
		log.Fatal(err)
	}
}
