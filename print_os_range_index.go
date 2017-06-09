package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, arg := range os.Args[0:] {
		fmt.Println(ind, arg)
	}
}
