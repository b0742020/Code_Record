package main

import "fmt"

func main() {
	z := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for x := 0; x < len(z); x++ {
		for y := 0; y < len(z[x]); y++ {
			fmt.Print(z[x][y], " ")
		}
		fmt.Println()
	}
}
