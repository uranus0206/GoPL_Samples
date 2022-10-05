// CF converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"uranus0206/tempconv"
)

func main() {
	// fmt.Println(tempconv.CToF(tempconv.BoilingC))
	argCnt := len(os.Args[1:])

	if argCnt > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			k := tempconv.Kelvin(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f),
				c, tempconv.CToF(c))

			fmt.Printf("%s = %s, %s = %s\n",
				k, tempconv.KToC(k),
				c, tempconv.CToK(c))

			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToK(f),
				k, tempconv.KToF(k))
		}

	} else {
		// No args
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			k := tempconv.Kelvin(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f),
				c, tempconv.CToF(c))

			fmt.Printf("%s = %s, %s = %s\n",
				k, tempconv.KToC(k),
				c, tempconv.CToK(c))

			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToK(f),
				k, tempconv.KToF(k))
		}

	}
}
