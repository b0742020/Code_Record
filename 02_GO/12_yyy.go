package main

import "fmt"

func main() {
	a := 1
	c := a << 1
	ptr := &c
	fmt.Println(*ptr)
	fmt.Println(c)
}
