// Prints all files in which the duplicated lines occur.
package main

// TODO: Get just the file names excluding any paths

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occursIn := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts, occursIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
				continue
			}
			countLines(arg, f, counts, occursIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:%d\t%s\n", occursIn[line], n, line)
		}
	}
}

// TODO: Make this a map from strings to an array of strings. That way we can
// filter out multiple occurrences of a file name. Currently this includes the
// file every time we count a letter. Make it so it just brings in files one
// time. I think I'll have to keep a map from strings to a slice of strings and
// check if the string is not present in that slice before adding it.
func countLines(fileName string, f *os.File, counts map[string]int, occursIn map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		occursIn[input.Text()] += " " + fileName
	}
}
