package main

import (
	"fmt"
	"math"
)

func main() {
	// var u uint8 = 255
	// fmt.Println(u, u+1, u*u)

	// var ii int8 = 127
	// fmt.Println(ii, ii+1, ii*ii)

	// var x uint8 = 1<<1 | 1<<5
	// var y uint8 = 1<<1 | 1<<2

	// fmt.Printf("%08b\n", x)
	// fmt.Printf("%08b\n", y)

	// fmt.Printf("%08b\n", x&y)
	// fmt.Printf("%08b\n", x|y)
	// fmt.Printf("%08b\n", x^y)
	// fmt.Printf("%08b\n", x&^y)

	// for i := uint(0); i < 8; i++ {
	// 	if x&(1<<i) != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// fmt.Printf("%08b\n", x<<1)
	// fmt.Printf("%08b\n", x>>1)

	// medals := []string{"gold", "silver", "bronze"}
	// for i := len(medals) - 1; i >= 0; i-- {
	// 	fmt.Println(medals[i])
	// }

	// f := 1e100
	// iif := int(f)
	// fmt.Printf("%d\n", iif)

	// o := 0666
	// fmt.Printf("%d %[1]o %#[1]o\n", o)

	// ascii := 'a'
	// unicode := '國'
	// newline := '\n'
	// fmt.Printf("%d %[1]c %[1]q\n", ascii)
	// fmt.Printf("%d %[1]c %[1]q\n", unicode)
	// fmt.Printf("%d %[1]q\n", newline)

	var ff float32 = 16777216
	fmt.Println(ff == ff+1)

	const e = 2.71828
	fmt.Printf("%f", e)

	const Avogadro = 6.02214129e23
	const Plank = 6.62606957e-34
	fmt.Printf("%f %[1]e\n", Avogadro)
	fmt.Printf("%f %[1]e\n", Plank)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
