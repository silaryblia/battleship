package main

import "fmt"

func main() {
	var a, b, c int
	a = 5
	b = 4
	c = 6
	//	fmt.Println(a + b + c)
	fmt.Println(sum(a, b, c))
}

func sum(a, b, c int) int {
	result := a + b + c
	return result
}
