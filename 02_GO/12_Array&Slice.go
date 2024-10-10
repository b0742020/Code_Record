package main

import "fmt"

func main() {
	a := 1
	c := a << 1
	ptr := &c
	fmt.Println(*ptr)
	fmt.Println(c)
	var count [3]int64
	count[0] = 1
	count[1] = 2
	count[2] = 3

	for index, value := range count {
		fmt.Printf("Index [%d] : %d\n", index, value)
	}
}
