// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {

	// Versionn 1
	// var s, sep string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	// Version 2
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
