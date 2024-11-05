package main

import "fmt"

func f0() {
	fmt.Println("Start")
	defer func() {
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()

	fmt.Println("End")
}

//匿名返回
func f1() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

//命名返回
func f2() (a int) {
	defer func() {
		a++
	}()
	return a
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
