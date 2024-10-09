package main

import "fmt"

func main() {
	//  0-5 total 6 numbers detected which is even and odd
	for n := range 6 {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}

	// 100 numbers detected which are negative, 2 digit and more than 1 digit
	for num := range 100 {
		if num < 0 {
			fmt.Println(num, "is negative")
		} else if num > 10 {
			fmt.Println(num, "has 2 digit")
		} else {
			fmt.Println(num, "has more than 1 digit")
		}

	}

	for j := range 1000 {
		if (j+1)%2 == 0 {
			fmt.Println(j+1, "is even")
		} else {
			fmt.Println(j+1, "is odd")
		}
	}

}
