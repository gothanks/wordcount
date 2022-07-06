package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		fmt.Println("вы не ввели дынные")
	} else {
		arr := strings.Split(args[1], " ")
		fmt.Println(len(arr) - 1)
	}
}
