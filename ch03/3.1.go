package main

import (
	"fmt"
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

	// var ff float32 = 16777216
	// fmt.Println(ff == ff+1)

	// const e = 2.71828
	// fmt.Printf("%f", e)

	// const Avogadro = 6.02214129e23
	// const Plank = 6.62606957e-34
	// fmt.Printf("%f %[1]e\n", Avogadro)
	// fmt.Printf("%f %[1]e\n", Plank)

	// for x := 0; x < 8; x++ {
	// 	fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	// }

	// var z float64
	// fmt.Println(z, -z, 1/z, -1/z, z/z)

	// nan := math.NaN()
	// fmt.Println(nan == nan, nan < nan, nan > nan)

	// s := "hello, world"
	// fmt.Println(len(s))
	// fmt.Println(s[0], s[7])

	// // c := s[len(s)]
	// // fmt.Println(c)

	// fmt.Println(s[0:5])

	// s := "left foot"
	// t := s
	// s += ", right foot"
	// fmt.Println(s)
	// fmt.Println(t)

	// rawStr := `This is a raw string.\n\n\n

	// Usage: '\t\t\t'
	// `
	// fmt.Println(rawStr)

	// unifiStr := "hello, 世"
	// fmt.Println(unifiStr)
	// fmt.Println(Contains(unifiStr, "世"))

	// s1 := "Hello, 世界"
	// fmt.Println(len(s1))
	// fmt.Println(utf8.RuneCountInString(s1))

	// for i := 0; i < len(s1); {
	// 	r, size := utf8.DecodeRuneInString(s1[i:])
	// 	fmt.Printf("%d\t%c\n", i, r)
	// 	i += size
	// }

	// for i, r := range "Hello, 世界" {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	// }

	ch1 := "這是中文測試"
	fmt.Printf("% x\n", ch1)
	chr := []rune(ch1)
	fmt.Printf("%x\n", chr)

	fmt.Println(string(1234567))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
