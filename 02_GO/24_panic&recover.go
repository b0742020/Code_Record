package main

import (
	"errors"
	"fmt"
)

func fn1() {
	fmt.Println("fn1 called")
}

func fn2() {
	panic("panic in fn2")
}

func fn3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("ERROR")
}

func fn4(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error : ", err)
		}
	}()
	return a / b
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("file not found")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error will be sent to me by email")
		}
	}()

	err := readFile("XXXX.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("continue...")
}
