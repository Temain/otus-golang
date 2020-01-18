package main

import (
	"flag"
	"github.com/Temain/otus-golang/hw-10/copier"
	"log"
)

var from string
var to string
var limit int
var offset int

func init() {
	flag.StringVar(&from, "from", "", "source file")
	flag.StringVar(&to, "to", "", "target file")
	flag.IntVar(&limit, "limit", 0, "limit in source file")
	flag.IntVar(&offset, "offset", 0, "offset in source file")
}

func main() {
	flag.Parse()

	//log.Printf("%v", from)
	//log.Printf("%v", to)
	//log.Printf("%v", limit)
	//log.Printf("%v", offset)

	err := copier.Copy(from, to, limit, offset)
	if err != nil {
		log.Fatal(err)
	}
}
