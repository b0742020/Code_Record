package main

import (
	"fmt"
)

func foo() (int, int) {
	return 10, 20
}

// area function takes two integers as input and returns their product
func getUserinfo() (string, int) {

	return "Gibson", 24

}

func main() {
	// _ is a blank identifier
	x, _ := foo()
	fmt.Println(x)
	a := 1
	b := 2
	// swap variables
	a, b = b, a
	fmt.Printf("a = %d , b = %d\n", a, b)
	const LENGTH int = 5
	const WWIDTH int = 5
	var area int
	const q, w, e = 1, false, "str"
	area = LENGTH * WWIDTH
	fmt.Println("Area of rectangle is ", area)
	fmt.Println(q, w, e)

	var username, _ = getUserinfo()
	fmt.Println(username)

	num := 10
	fmt.Printf("num = %b\n", num)

	m1 := 8.2
	m2 := 3.5
	fmt.Println(m1 + m2)

	// var str = "123-456-789"
	// arr := strings.Split(str, "-")
	// str2 := strings.Join(arr, "*")
	// fmt.Println(str, arr, str2)

	// arr := []string{"php", "golang", "python"}
	// str3 := strings.Join(arr, "-")
	// fmt.Printf("%T\n", str3)
	// fmt.Println(str3)

	// str1 := "hello world"
	// str2 := "w"
	// num = strings.LastIndex(str1, str2)
	// fmt.Println(num)

	s := "hello world"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	//修改英文字串 用byte(ascii)
	s1 := "hello world"
	byteStr := []byte(s1)
	byteStr[4] = ' '
	fmt.Println(string(byteStr))

	//修改中文字串 用rune(utf8)
	s2 := "台灣"
	runeStr := []rune(s2)
	runeStr[0] = '臺'
	fmt.Println(string(runeStr))
}
