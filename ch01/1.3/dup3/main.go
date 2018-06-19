// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

//read the entire input into memory in one big gulp,
//split it into lines all at once, then process the lines.
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			//If there is an empty line at the end of file, there will be an empty string
			if line == "" {
				continue
			}
			counts[line]++
		}
	}
	for line, n := range counts {
		//more than once
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
