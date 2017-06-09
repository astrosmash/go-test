/*
There are two more efficient ways of producing the output.
If we have big amount of args, current implementation may be not efficient due to the need for garbage collector to collect old values of var s.
1) We may separate the whole os.Args slice with spaces by fmt.Println(strings.Join(os.Args[1:], " "))
2) We don't care much about formatting and are letting fmt to do the job, by fmt.Println(os.Args[1:]), but formatting may add [] braces
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		// fmt.Print()
	}
	fmt.Println(s)
}
