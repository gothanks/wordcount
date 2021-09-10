package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	orig := strings.Join(flag.Args(), "")
	words := strings.Fields(orig)
	fmt.Println(len(words))
}
