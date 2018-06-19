//Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"os"
	"fmt"
	"bytes"
)

type Counter struct {
	files map[string]int
	n     int
}

func countLines(file *os.File, counts map[string]Counter) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counter := counts[line]
		counter.n++
		//map must init
		if counter.files == nil {
			counter.files = make(map[string]int)
		}
		counter.files[file.Name()]++
		//write back to map
		counts[line] = counter
	}
}

func main() {
	counts := make(map[string]Counter)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, counter := range counts {
		var buf bytes.Buffer
		for file := range counter.files {
			if buf.Len() != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(file)
		}
		fmt.Printf("%d\t%s\t%s\n", counter.n, line, buf.String())
	}
}
