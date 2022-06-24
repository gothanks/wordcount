package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	fmt.Println(len(strings.Fields(arg)))
}
