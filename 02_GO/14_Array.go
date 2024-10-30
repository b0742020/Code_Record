package main

func main() {
	/*array宣告以及初始化*/

	/*初始化法一*/
	// arr1[0] = 1
	// arr1[1] = 1
	// arr1[2] = 1
	// arr1[3] = 1
	// arr1[4] = 1
	// fmt.Println("arr1 :", arr1)

	/*初始化法二*/
	// var arr1 = [5]int{1, 1, 1, 1, 1}
	// fmt.Println("arr1 :", arr1)

	/*初始化法三*/
	/*任意長度但非動態改變*/
	// var arr1 = [...]int{1, 1, 1, 1, 1, 2, 3, 4, 5}
	// fmt.Println("arr1 :", arr1)
	// fmt.Println(len(arr1))

	/*初始化法四*/

	// arr := [...]int{0: 1, 1: 10, 2: 40, 9: 100}
	// fmt.Println(arr)

	/*循環遍歷 For loop & For range loop*/

	// arr := [...]string{"apple", "banana", "orange", "grape", "pear"}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("Index: %d Value: %s\n", i, arr[i])
	// }
	// for k, v := range arr {
	// 	fmt.Println(k, v)
	// }
	// arr1 := [...]int{1, 2, 3, 4, 5}
	// var sum int
	// for i, v := range arr1 {
	// 	arr1[i] = v * 2
	// 	sum += arr1[i]
	// }
	// fmt.Println(sum)
	// fmt.Println(sum / len(arr1))

	/*從陣列中找最大值*/
	// var arr = [...]int{1, -8, 9, 10, 89, 55, 66}
	// max := arr[0]
	// index := 0
	// for i := 0; i < len(arr); i++ {
	// 	if max < arr[i] {
	// 		max = arr[i]
	// 		index = i
	// 	}
	// }
	// fmt.Println("Max value is:", max, "Index is:", index)

	/*從陣列中找到合為8的兩個元素並標示出來*/
	// var arr = [...]int{1, 3, 5, 7, 8}
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i]+arr[j] == 8 {
	// 			fmt.Printf("(%v,%v)\n", i, j)
	// 		}
	// 	}
	// }

}
