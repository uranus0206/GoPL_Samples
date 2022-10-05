package main

import (
	"fmt"
	"time"
	"uranus0206/popcount"
)

func main() {
	start := time.Now()
	a := popcount.PopCount(255 * 255 * 255 * 255 * 255)
	elapsed := time.Since(start)
	fmt.Printf("elapsd: %v, count: %v\n", elapsed, a)

	start = time.Now()
	a = popcount.PopCountLoop(255 * 255 * 255 * 255 * 255)
	elapsed = time.Since(start)
	fmt.Printf("loop elapsd: %v, count: %v\n", elapsed, a)
}
