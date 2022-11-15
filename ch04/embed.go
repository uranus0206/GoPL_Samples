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

type Cellphone struct {
	ID   int
	name string
}
type Iphone struct {
	Cellphone
}
type HTC struct {
	Cellphone
}

func main() {

	// w := Wheel{Circle{Point2{8, 8}, 5}, 20}
	// // w := Wheel{
	// // 	Circle: Circle{
	// // 		Center: Point2{X: 8, Y: 8},
	// // 		Radius: 5,
	// // 	},
	// // 	Spokes: 20,
	// // }
	// fmt.Printf("%#v\n", w)

	// w.X = 42
	// fmt.Printf("%#v\n", w)

	c1 := Circle{
		Center: Point2{
			X: 1,
			Y: 1,
		},
		Radius: 2,
	}
	c1.X = 2
	c1.Y = 3

	newIphone := Iphone{
		Cellphone{
			ID:   1,
			name: "iphone",
		},
	}
	fmt.Printf("%v", newIphone)
}
