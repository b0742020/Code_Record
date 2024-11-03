package main

import (
	"fmt"
	"sort"
)

func bubble_sort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

func bubble_sort_reverse(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

func mapSort(map1 map[string]string) string {
	var sliceekey []string
	for key, _ := range map1 {
		sliceekey = append(sliceekey, key)
	}
	sort.Strings(sliceekey)
	var str string
	for _, key := range sliceekey {
		str += key + ":" + map1[key] + "\n"
	}
	return str
}

func main() {
	arr := []int{5, 2, 8, 3, 9, 1}
	sorted_arr := bubble_sort(arr)
	fmt.Println(sorted_arr)

	arr_reverse := []int{5, 2, 8, 3, 9, 1}
	sorted_arr_reverse := bubble_sort_reverse(arr_reverse)
	fmt.Println(sorted_arr_reverse)
}
