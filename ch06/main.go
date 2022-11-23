package main

import (
	"fmt"

	"ch06/geometry"
)

func main() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim)
	fmt.Println(geometry.PathDistance(perim))
	fmt.Println(perim.Distance())

	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := geometry.Point{1, 2}
	// &p.ScaleBy(2) // Wrong!!!
	(&p).ScaleBy(2)
	fmt.Println(p)

	q := geometry.Point{1, 2}
	qptr := &q
	qptr.ScaleBy(2)
	fmt.Println(q)

}
