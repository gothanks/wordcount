package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if string(args[1]) == "" {
		fmt.Println(0)
	}
	arr := strings.Split(args[1], " ")
	fmt.Println(len(arr))

}
