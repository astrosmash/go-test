package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type dupCount struct {
	count     int
	fileCount map[string]int
}

func main() {
	counts := make(map[string]dupCount)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%#v\t %#v\n", n, line)
	}
}

func countLines(r io.Reader, counts map[string]dupCount, filename string) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		counts[line] = addDupe(counts[line], filename)
	}
}

func addDupe(dupes dupCount, filename string) dupCount {
	if dupes.count == 0 {
		dupes.fileCount = make(map[string]int)
	}
	dupes.fileCount[filename]++
	dupes.count++
	return dupes
}
