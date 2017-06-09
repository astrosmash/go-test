// measure the time 'range' takes to process.
// if we print to the pipe we won't read, e.g. 2> /dev/null, execution time reduces

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for ind, arg := range os.Args[0:] {
		fmt.Fprintln(os.Stderr, ind, arg)
	}
	fmt.Println(time.Since(start))
}
