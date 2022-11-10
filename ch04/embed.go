package main

import "fmt"

type Point2 struct {
	X, Y int
}
type Circle struct {
	Center Point2
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {

	w := Wheel{Circle{Point2{8, 8}, 5}, 20}
	// w := Wheel{
	// 	Circle: Circle{
	// 		Center: Point2{X: 8, Y: 8},
	// 		Radius: 5,
	// 	},
	// 	Spokes: 20,
	// }
	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}
