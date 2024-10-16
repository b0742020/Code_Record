package main

import "fmt"

func main() {

	count := []int64{1, 2, 3, 4, 5}
	subcount := count[0:2]
	fmt.Println(subcount)

	langs := []string{"GO", "Java", "Python", "C++"}
	for index, e := range langs {
		fmt.Println(index, e)
	}

	slice := make([]int, 3)
	for i := 0; i < len(slice); i++ {
		y := i + 1
		slice[i] = y * 2
	}
	fmt.Println(slice)
}
