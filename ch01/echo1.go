// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// Exercise 1.1
	// Versionn 1
	// var s, sep string

	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// // Version 2
	// s, sep := "", ""

	// for _, arg := range os.Args[0:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// Version 3
	// fmt.Println(strings.Join(os.Args[0:], " "))

	// Exercise 1.2
	// Versionn 1
	// for i := 0; i < len(os.Args); i++ {
	// 	s := "index: " + strconv.Itoa(i) + ", arguments: " + os.Args[i]
	// 	fmt.Println(s)
	// }

	// // Version 2
	// for idx, arg := range os.Args[0:] {
	// 	fmt.Println("index: " + strconv.Itoa(idx) + ", arg: " + arg)
	// }

	// Exercise 1.3

	// Versionn 1
	var s, sep string

	start1 := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%d elapsed nano v1\n", time.Since(start1).Nanoseconds())
	// // Version 2
	s1, sep1 := "", ""

	start1 = time.Now()
	for _, arg := range os.Args[1:] {
		s1 += sep1 + arg
		sep1 = " "
	}
	fmt.Println(s1)
	fmt.Printf("%d elapsed nanao v2\n", time.Since(start1).Nanoseconds())

	// Version 3
	start1 = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%d elapsed nano v3\n", time.Since(start1).Nanoseconds())
}
