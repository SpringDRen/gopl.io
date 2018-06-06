// Echo1 prints its command-line arguments on a single line
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//1. The := symbol is part of a short variable declaration
	//2. i++、i--   These are statements, not expressions. So j = i++ is illegal,and they are postfix only, so --i is not legal either.
	//3. The for loop is the only loop statement in Go.
	for i := 1; i < len(os.Args); i++ {
		//字符串链接 + 效率较低，消耗资源较大
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
