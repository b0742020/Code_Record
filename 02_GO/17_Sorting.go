package main

import (
	"fmt"
	"sort"
)

func main() {
	// var nums = [...]int{1, 2, 6, 8, 7, 9, 70}
	// var target int = 0
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == 77 {
	// 			fmt.Printf("[%v %v]\n", nums[i], nums[j])
	// 			target = nums[i] + nums[j]
	// 		}
	// 	}

	// }
	// fmt.Printf("Target : %v\n", target)

	//selection sort(從小到大)
	var nums = []int{9, 8, 7, 1, 5, 4}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	fmt.Println(nums)

	//selection sort(從大到小)
	var nums_1 = []int{9, 8, 7, 1, 5, 4}
	for i := 0; i < len(nums_1); i++ {
		for j := i + 1; j < len(nums_1); j++ {
			if nums_1[i] < nums_1[j] {
				temp := nums_1[i]
				nums_1[i] = nums_1[j]
				nums_1[j] = temp
			}
		}
	}
	fmt.Println(nums_1)

	//bubble sort(從小到大)
	var nums_2 = []int{9, 8, 7, 1, 5, 4}
	for i := 0; i < len(nums_2); i++ {
		for j := 0; j < len(nums_2)-1-i; j++ {
			if nums_2[j] > nums_2[j+1] {
				temp := nums_2[j]
				nums_2[j] = nums_2[j+1]
				nums_2[j+1] = temp
			}
		}
	}
	fmt.Println(nums_2)

	//bubble sort(從大到小)
	var nums_3 = []int{9, 8, 7, 1, 5, 4}
	for i := 0; i < len(nums_3); i++ {
		for j := 0; j < len(nums_3)-1-i; j++ {
			if nums_3[j] < nums_3[j+1] {
				temp := nums_3[j]
				nums_3[j] = nums_3[j+1]
				nums_3[j+1] = temp
			}
		}
	}
	fmt.Println(nums_3)

	//Sort package(降序)
	intList := []int{9, 8, 7, 1, 5, 4}
	sort.Ints(intList)
	fmt.Println(intList)

	//Sort package(升序)
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)

}
