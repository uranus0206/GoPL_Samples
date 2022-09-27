// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of names files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFile := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "", dupFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, arg, dupFile)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, dupFile[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file string, dupFile map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		shouldAppend := true
		arr := dupFile[input.Text()]
		for _, v := range arr {
			if v == file {
				shouldAppend = false
			}
		}

		if shouldAppend {
			dupFile[input.Text()] = append(dupFile[input.Text()], file)
		}

	}
}
