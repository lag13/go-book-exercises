// Measure time difference between the potentially inefficient version of the
// program and the one which uses strings.Join()
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type timeFn func() int64

func main() {
	const (
		numRuns = 500
	)
	timeFirstVersion := avgTime(timeFirstVersion, numRuns)
	timeSecondVersion := avgTime(timeSecondVersion, numRuns)
	var which string
	var dt int64
	if timeFirstVersion < timeSecondVersion {
		which = "first"
		dt = timeSecondVersion - timeFirstVersion
	} else {
		which = "second"
		dt = timeFirstVersion - timeSecondVersion
	}
	fmt.Printf("First version time = %v\n", timeFirstVersion)
	fmt.Printf("Second version time = %v\n", timeSecondVersion)
	fmt.Printf("The %v method is %v nanoseconds faster\n", which, dt)
}

func avgTime(fn timeFn, numRuns int) int64 {
	var avgTime int64
	avgTime = 0
	for i := 0; i < numRuns; i++ {
		avgTime += fn()
	}
	return avgTime / int64(numRuns)
}

func timeFirstVersion() int64 {
	var s1, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s1 += sep + os.Args[i]
		sep = " "
	}
	return time.Since(start).Nanoseconds()
}

func timeSecondVersion() int64 {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	return time.Since(start).Nanoseconds()
}
