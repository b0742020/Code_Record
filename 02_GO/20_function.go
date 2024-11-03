package main

import "fmt"

/*
	func functionName(parameter1 type1, parameter2 type2) return_type {
		// function body
	}
*/

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

//函數為可變參數，可變參數是指函數的參數數量不固定，Golang中可以通過在參數明後加...來表示。
func sumFn2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn3(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func calculate(x int, y ...int) (int, int) {
	sum := x
	for _, v := range y {
		sum += v
	}
	sum1 := x
	for _, v := range y {
		sum1 -= v
	}
	return sum, sum1
}

func calculate_1(x int, y ...int) (sum int, sub int) {
	temp := x
	for _, v := range y {
		temp += v
	}
	sum = temp
	temp1 := x
	for _, v := range y {
		temp1 -= v
	}
	sub = temp1
	return
}
func main() {
	sum := sumFn2(1, 2, 3, 4, 5)
	fmt.Println(sum)

	sum1 := sumFn(10, 20)
	fmt.Println(sum1)

	sum2 := sumFn3(10, 20, 30, 40, 50)
	fmt.Println(sum2)

	sum3, sum4 := calculate(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println(sum3, sum4)

	sum5, sum6 := calculate_1(10, 20, 30, 40, 50, 60, 70, 80, 90, 1000)
	fmt.Println(sum5, sum6)

	sum7, _ := calculate_1(10, 20, 30, 40, 50, 60, 70, 80, 90, 1000)
	fmt.Println(sum7)
}
