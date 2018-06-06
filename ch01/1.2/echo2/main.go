// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

//range
func main() {
	//4 ways to declare a string variable
	//s := ""       //a short variable declaration, is the most compact, but it may be used only within a function, not for package-level variables.
	//var s string  //default initialization
	//var s = ""    //rarely used except when declaring multiple variables.
	//var s string = "" //Use when variable’s type is not the same as initial value
	s, sep := "", ""
	//1. Range produces a pair of values: the index and the value of the element at that index.
	//2. Go does not permit unused local variables. Use the blank identifier (_) whenever syntax requires a variable name but program logic does not.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
