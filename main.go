package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	
		arr := strings.Split(args[1], " ")
		fmt.Println(len(arr) - 1)

}
