package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 5, 7, 99, 8}
	arr1 := arr
	arr[0] = 11
	fmt.Println(arr)
	fmt.Println(arr1)
}
