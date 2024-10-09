package main

import (
	"fmt"
)

func sayHello(x string) {
	fmt.Println(x)
}

func sum() {
	var result int = 0
	for i := 1; i < 11; i++ {
		result += i
	}
	fmt.Println(result)
}

func main() {
	var x string = ""
	var stockcode = 940
	var enddate = "2021-12-31"
	var url = "Code = %d & endDate= %s"
	fmt.Printf(url, stockcode, enddate)
	fmt.Print("\n")
	fmt.Print("Enter your name: ")
	fmt.Scan(&x)
	fmt.Println(len(x))
	sayHello(x)
	sum()
}
