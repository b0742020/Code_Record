package main

import "fmt"

func main() {
	/*called by value*/
	// var arr = [...]int{1, 2, 5, 7, 99, 8}
	// arr1 := arr
	// arr[0] = 11
	// fmt.Println(arr)
	// fmt.Println(arr1)
	/*called by reference*/
	// 	var arr2 = []int{1, 2, 5, 7, 99, 8}
	// 	arr3 := arr2
	// 	arr2[0] = 11
	// 	fmt.Println(arr2)
	// 	fmt.Println(arr3)
	/*2D array */
	var arr = [...][2]string{{"apple", "banana"},
		{"orange", "pear"},
		{"peach", "pineapple"}}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}

	}
}
