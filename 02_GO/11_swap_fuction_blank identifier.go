package main

import "fmt"

func foo() (int, int) {
	return 10, 20
}

// area function takes two integers as input and returns their product
func area(x, y int) int {

	return x * y

}
func main() {
	// _ is a blank identifier
	x, _ := foo()
	fmt.Println(x)
	a := 1
	b := 2
	c := area(a, b)
	// swap variables
	a, b = b, a
	fmt.Printf("a = %d , b = %d\n", a, b)
	fmt.Println("c =", c)
	const LENGTH int = 5
	const WWIDTH int = 5
	var area int
	const q, w, e = 1, false, "str"
	area = LENGTH * WWIDTH
	fmt.Println("Area of rectangle is ", area)
	fmt.Println(q, w, e)
}
