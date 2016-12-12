package main

import (
	"flag"
	"github.com/narita-takeru/sqlintercept"
)

var (
	src = flag.String("src", ":3305", "source")
	dst = flag.String("dst", "localhost:3306", "destination")
)

func main() {

	flag.Parse()
	sqlintercept.Start(*src, *dst)
}
