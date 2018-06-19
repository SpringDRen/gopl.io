// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//读取标准输入，然后计数，读到 EOF标志后，停止读取标准输入，输出计数
//EOF：  linux = ctrl+d   windows = ctrl+z. 然后回车
//idea的run console : ctrl+d
//windows下git bash 暂不支持
func main() {
	countMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//NOTE: ignoring potential errors from input.Err()
		//if err := input.Err(); err != nil {
		//	fmt.Fprint(os.Stderr, err)
		//} else {
		countMap[input.Text()]++
		//}
	}
	for line, n := range countMap {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
