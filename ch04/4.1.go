package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%v %v %[2]T\n", i, v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	symbol := [...]string{USD: "$", EUR: "c", GBP: "L", RMB: "Y"}
	fmt.Println(RMB, symbol[RMB])

	// sha256
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	months := [...]string{
		1: "Jan", 2: "Feb", 3: "Mar",
		4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep",
		10: "Oct", 11: "Nov", 12: "Dec",
	} // index 0 will be empty string

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	fmt.Printf("Q2 %x\n, month7 %x\n, summer7 %x\n, endlessSummer7 %x\n", &Q2[2], &months[6], &summer[0], &endlessSummer[0])
	// variables share the same slice address!!!

	a1 := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a1[:])
	fmt.Println(a1)

	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)

	var x1, y1 []int
	for i := 0; i < 10; i++ {
		y1 = appendInt(x1, i)
		fmt.Printf("%d cap=%d\t%v\t%x\n", i, cap(y1), y1, &y1[0])
		x1 = y1
	}

	s1 := []int{5, 6, 7, 8, 9}
	fmt.Printf("s1 address: %p\n", &s1)
	fmt.Println(removeInt(s1, 2))
	fmt.Println(removeInt2(s1, 2))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// append func
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// slice remove
// Notice the return will point the same address of original slice
func removeInt(slice []int, i int) []int {
	fmt.Printf("slice: %p\n", slice)
	copy(slice[i:], slice[i+1:])
	fmt.Println("After copy ", slice)
	return slice[:len(slice)-1]
}

func removeInt2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	fmt.Println("2 After copy ", slice)
	return slice[:len(slice)-1]
}
