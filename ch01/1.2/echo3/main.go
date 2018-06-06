// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	//A simpler and more efficient solution would be to use the Join function from the strings package
	fmt.Print(strings.Join(os.Args[1:], " "))
	//If we don’t care about format but just want to see the values
	fmt.Print(os.Args[1:])
}
