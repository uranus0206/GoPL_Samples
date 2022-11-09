package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Printf("values[:0]=%v\n", values[:0])
	return appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {
	slice := []int{1, 9, 2, 8, 3, 7, 6, 4, 5}
	fmt.Println(sort(slice))

}
