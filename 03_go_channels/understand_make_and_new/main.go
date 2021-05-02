package main

import "fmt"

func main() {
	var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	var v  []int = make([]int, 10) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	var p2 *[]int = new([]int)
	*p2 = make([]int, 10, 10)

	// Idiomatic:
	v2 := make([]int, 10)
	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(v)
	fmt.Println(v2)
}