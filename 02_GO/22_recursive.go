package main

import "fmt"

func recur(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		recur(n)
	}
}

func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return n
	}
}

func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}

func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	recur(100)
	fmt.Println(fn2(100))
	fmt.Println(fn3(10))

	var fn = adder1()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	var fn1 = adder2()
	fmt.Println(fn1(10))
	fmt.Println(fn1(10))
	fmt.Println(fn1(10))
}
