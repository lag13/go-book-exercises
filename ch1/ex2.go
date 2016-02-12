// Print index and value of each argument
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("value => %v\n", v)
		fmt.Printf("key => %v\n", i)
	}
}
