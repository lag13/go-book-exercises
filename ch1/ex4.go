// Prints all files in which the duplicated lines occur.
package main

// I had some difficulties getting this implementation to work that revolved
// around how maps work. The crux of the problem is illustrated by the code in
// this post:
// 		https://github.com/golang/go/issues/3117
// In particular if you have a map like var `m map[string]struct{one, two int}`
// then you cannot do things like `m["key"].one` even "key" is an existing key
// in the map. More generally, regardless of the value that the map holds this
// expression is invalid: `&m["key"]`. Googling this issue gave good results.
// 		https://www.google.com/search?sourceid=chrome-psyapi2&ion=1&espv=2&ie=UTF-8&q=cannot%20assign%20to%20map%20golang&oq=cannot%20assign%20to%20&aqs=chrome.1.69i57j0l5.5002j0j7
// The issue revolves around the fact that we cannot take the address of a
// value stored in a map key. A more proper way of phrasing that is that a map
// index expression is not "addressable". The reason for that is because as a
// map changes (grows/shrinks etc...) how the data is stored can be changed. So
// if you had a struct A stored in a map: `m["here"] = A`. The address of the
// value at m["here"] might get copied/moved to a new address. The way around
// this is either to have a map whose values are pointers to structs (then you
// *can* do things like m["here"].structField because the value of m["here"] is
// in fact an address and if a variable stores the address of a struct or the
// struct value itself, you can access the member variables with the same dot
// notation). Or you could have a map to struct values and whenever you wanted
// to modify the struct you have to copy it into a temporary variable, make
// the changes, then copy that value back into the map.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type duplicateCount struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]*duplicateCount)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%s:%d\t%s\n", strings.Join(n.files, ","), n.count, line)
		}
	}
}

func countLines(fileName string, f *os.File, counts map[string]*duplicateCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var dc *duplicateCount
		dc, ok := counts[input.Text()]
		if !ok {
			counts[input.Text()] = new(duplicateCount)
			dc = counts[input.Text()]
		}
		dc.count++
		if !contains(dc.files, fileName) {
			dc.files = append(dc.files, fileName)
		}
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
