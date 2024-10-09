package main

import "fmt"

func main() {
	var x int = 10
	var y float32 = 3.14
	var z bool = true
	var name string = "John"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(name)
	f := "Gibson"
	// f is a local variable and cannot be accessed outside the function
	fmt.Println(f)
}
