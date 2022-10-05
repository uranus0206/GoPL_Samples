// Echo4 prints its ciommand-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

}

func delta(old, new int) int {

	// p := new(int) // new has been used as variable, the new function is unavailable in this block

	return new - old
}

func newInt() *int {
	// method 1
	// return new(int)

	// method 2
	var dummy int
	return &dummy
}
