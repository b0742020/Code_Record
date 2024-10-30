package main

func main() {
	/*
		break 表示跳出當前循環
	*/
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("Loop is over")

	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= 10; j++ {
	// 		if j == 3 {
	// 			break
	// 		}
	// 		fmt.Printf("i = %d j =%d\n", i, j)
	// 	}
	// }
	/////////////////////////////////////////////
	/*
		label 標出想break的循環
	*/
	// label1:
	// 	for i := 0; i < 10; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 5 {
	// 				break label1
	// 			}
	// 			fmt.Printf("i = %v\tj =%v \n", i, j)
	// 		}
	// 	}
	/*continue 表示跳過當前循環的下一個迭代*/
	/*
		上下等效
	*/
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 5 {
	// 				break
	// 			}
	// 			fmt.Printf("i = %v\tj = %v \n", i, j)
	// 		}
	// 	}
	// 	fmt.Println("///////////////")
	// label2:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 5 {
	// 				continue label2
	// 			}
	// 			fmt.Printf("i = %v\tj = %v \n", i, j)
	// 		}
	// 	}
	/////////////////////////////////////////////
	/*goto 使用方法*/
	// 	var x = 0
	// 	if x < 10 {
	// 		fmt.Println(x)
	// 		goto label3
	// 	}
	// label3:
	// 	fmt.Println("x is greater than or equal to 10")
}
