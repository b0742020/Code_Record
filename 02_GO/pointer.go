package main

import "fmt"

func value(x int) {
	x = 10
}

func reference(x *int) {
	*x = 40
}

//pointer   *取地址裡的值  &取變數的地址
func main() {
	// var a = 5
	// value(a)
	// fmt.Println("a =", a) // Output: a = 10
	// reference(&a)
	// fmt.Println("a =", a) // Output: a = 40
	var a = new(int)
	var b = new(string)
	fmt.Printf("value :%v type : %T pointer : %v\n", a, a, *a)
	fmt.Printf("value :%v type : %T pointer : %v\n", b, b, *b)
}
