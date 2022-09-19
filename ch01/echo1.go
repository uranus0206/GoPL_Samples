// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Versionn 1
	// var s, sep string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// // Version 2
	// s, sep := "", ""

	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// Version 3
	fmt.Println(strings.Join(os.Args[1:], " "))

}
