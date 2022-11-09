package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct{ X, Y int }
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var dilbert Employee

	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	// employeeOfTheMonth.Position += " (proactive team player)"
	//equivalent to: (*employeeOfTheMonth).Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(employeeOfTheMonth.Position)
	fmt.Printf("%p\t%p\n", &employeeOfTheMonth, &(*employeeOfTheMonth))

	fmt.Println(Scale(Point{1, 2}, 5))
}

func Scale(p Point, factor int) Point {
	// return new pointer with factor
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	// pass with pointer to save memeory
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
