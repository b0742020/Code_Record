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
func test() bool {
	fmt.Println("Test......")
	return true
}

func main() {
	/*FormatInt
	first parameter is  int64 format
	second  parameter is what your base like binary octal decimal hexdecimal etc....
	*/
	// var i int = 10
	// str1 := strconv.FormatInt(int64(i), 10)
	// fmt.Println(str1)
	//////////////////////////////////////////////////////
	/*FormatFloat
	1st  要轉換的值
	2nd  轉換後的格式
	3rd  要轉換的精度
	4th  轉換的位數
	*/
	// var f float32 = 20.2333333
	// str2 := strconv.FormatFloat(float64(f), 'f', 6, 32)
	// fmt.Printf("Value : %v Type : %T\n", str2, str2)
	//////////////////////////////////////////////////////
	/*String convert to int
	 */
	// 	str := "123456"
	// 	fmt.Printf("Value : %v Type : %T\n", str, str)
	// 	num, _ := strconv.ParseInt(str, 10, 64)
	// 	fmt.Printf("Value : %v Type : %T\n", num, num)
	//////////////////////////////////////////////////////
	/*String convert to float
	 */
	// str := "123456.33333333"
	// num, err := strconv.ParseFloat(str, 64)
	// fmt.Printf("Value : %v Type : %T Value : %v\n", num, num, err)
	//////////////////////////////////////////////////////
	/*取餘數*/
	// fmt.Println(10 % -3)
	// fmt.Println(-10 % 3)
	//////////////////////////////////////////////////////
	/*邏輯與(前面是Flase後面就不會執行)邏輯或(前面是True後面就不會執行)短路*/
	// var a int = 20
	// if a > 10 && test() {
	// 	fmt.Println("Execute.....")
	// }

	// if a > 30 && test() {
	// 	fmt.Println("Execute.....")
	// }

	// if a > 30 || test() {
	// 	fmt.Println("Execute.....")
	// }
	//////////////////////////////////////////////////////
	/*值對換*/
	// 1st way
	// var a int = 10
	// var b int = 20
	// t := a
	// a = b
	// b = t
	//  fmt.Println(a, b)
	//2nd way
	// var a int = 10
	// var b int = 20
	// a, b = b, a
	// fmt.Println(a, b)
	//////////////////////////////////////////////////////
	/*比大小*/
	// var a = 34
	// var b = 23
	// var max int
	// if a > b {
	// 	max = a

	// } else {
	// 	max = b
	// }
	// fmt.Printf("Max value is : %v", max)
	//////////////////////////////////////////////////////
	/*For Loop*/
	// for i := 1 ; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// var sum int
	// var count int
	// for i := 1; i <= 100; i++ {
	// 	if i%9 == 0 {
	// 		sum += i
	// 		count++
	// 		fmt.Println(i)
	// 	}
	// }
	// fmt.Println(sum)
	// fmt.Println(count)

	/*雙重for loop flow
	1. i=0 print 6 顆星星
	2. i=1 print 6 顆星星
	3. i=2 print 6 顆星星
	4. i=3 print 6 顆星星
	5. i=4 print 6 顆星星
	6. i=5 print 6 顆星星
	*/
	// var row = 5
	// var col = 6
	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// var row = 5
	// for i := 1; i <= row; i++ {
	// 	for j := 1; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	//99乘法表
	// var row = 9
	// for i := 1; i <= row; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%v*%v=%v \t", i, j, i*j)
	// 	}
	// 	fmt.Println("")
	// }

	// for range loop
	// var str = "Hello Golang"
	// for k, v := range str {
	// 	fmt.Printf("Index : %v & Value : %c\n", k, v)
	// }

	// var arr = []string{"apple", "banana", "orange"}
	// for k, v := range arr {
	// 	fmt.Printf("Index : %v & Value : %v\n", k, v)
	// }
}
