package main

import (
	"fmt"
	"strconv"
)

func main() {
	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x))

	// fmt.Println(strconv.FormatInt(int64(x), 2))
	// s := fmt.Sprintf("x=%b", x)
	// fmt.Println(s)

	x, err := strconv.Atoi("123"
	fmt.Println(x, err)
	y, err2 := strconv.ParseInt("123", 10, 64)
	fmt.Println(y, err2)

	var inputStr int
	x, err3 := fmt.Scanf("%d", &inputStr)
	fmt.Println(inputStr, x, err3)
}
