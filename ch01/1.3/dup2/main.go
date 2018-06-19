// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"os"
	"fmt"
)

func countLines(files *os.File, countMap map[string]int) {
	input := bufio.NewScanner(files)
	for input.Scan() {
		countMap[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//读取标准输入或者文件，按行计数
//有参数则读取文件；没有参数，则读取标准输入
func main() {
	counts := make(map[string]int)
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

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
