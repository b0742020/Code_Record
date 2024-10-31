package main

func main() {
	// 	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 	for k, v := range arr {
	// 		fmt.Printf("arr[%d] = %d\n", k, v)
	// 	}
	// 	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 	for k, v := range arr1 {
	// 		fmt.Printf("arr1[%d] = %d\n", k, v)
	// 	}
	// 	var arr2 []int
	// 	var arr3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 	fmt.Println(arr2, arr3)
	// 	fmt.Println(arr2 == nil, arr3 == nil)
	// a := [5]int{55, 56, 57, 58, 59}
	// fmt.Println(a)
	// b := a[1:4]
	// fmt.Println(b)
	// c := a[2:]
	// fmt.Println(c)
	// d := a[:3]
	// fmt.Println(d)
	/*Slice 長度與容量*/
	//長度: Slice的所包含的元素個數
	//容量: Slice的所包含的第一個元素到底層陣列元素的個數
	// s := []int{2, 3, 88, 4, 9, 77, 86, 98}
	// fmt.Printf("長度 : %v 容量 : %v\n", len(s), cap(s))
	// a := s[2:]
	// fmt.Printf("長度 : %v 容量 : %v\n", len(a), cap(a))
	// b := s[1:3]
	// fmt.Printf("長度 : %v 容量 : %v\n", len(b), cap(b))
	// c := s[:3]
	// fmt.Printf("長度 : %v 容量 : %v\n", len(c), cap(c))

	/*make() create a Slice make([]Type, len, cap)*/

	//1. 在Slice 新增DATA 用append()
	// var slice_A []int
	// fmt.Printf("len : %v , cap : %v\n", len(slice_A), cap(slice_A))
	// slice_A = append(slice_A, 22, 33, 55, 77, 88, 11)
	// var slice_B []int
	// slice_B = append(slice_B, 100, 23, 55, 88, 99)
	// fmt.Printf("%v %v %v \n", slice_B, len(slice_B), cap(slice_B))

	//2. 在append()合併Slice
	// slice_A = append(slice_A, slice_B...)
	// fmt.Println(slice_A)

	//3. Slice的擴容策略
	// var slice_C []int
	// for i := 1; i <= 10; i++ {
	// 	slice_C = append(slice_C, i)
	// 	fmt.Printf("%v %v %v \n", slice_C, len(slice_C), cap(slice_C))
	// }

	//4.copy() 複製Slice
	// slice_A := []int{1, 2, 3, 5}
	// slice_B := make([]int, 4, 4)
	// copy(slice_B, slice_A)
	// fmt.Println(slice_A)
	// fmt.Println(slice_B)
	// slice_B[0] = 111
	// fmt.Println(slice_A)
	// fmt.Println(slice_B)

	//5.透過append() 刪除元素
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// a = append(a[:3], a[4:]...)
	// fmt.Println(a)
}
