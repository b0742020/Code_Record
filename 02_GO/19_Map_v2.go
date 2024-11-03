package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	/*
		我們想要在Slice中存放一系列用戶訊息，這時候我們可以用map類型的切片來存放用戶訊息。
	*/
	// var userinfo = make([]map[string]string, 4)
	// fmt.Println(userinfo)
	// if userinfo[0] == nil {
	// 	userinfo[0] = make(map[string]string)
	// 	userinfo[0]["username"] = "John"
	// 	userinfo[0]["age"] = "22"
	// 	userinfo[0]["gender"] = "male"
	// }

	// if userinfo[1] == nil {
	// 	userinfo[1] = make(map[string]string)
	// 	userinfo[1]["username"] = "Gibson"
	// 	userinfo[1]["age"] = "23"
	// 	userinfo[1]["gender"] = "male"
	// }

	// if userinfo[2] == nil {
	// 	userinfo[2] = make(map[string]string)
	// 	userinfo[2]["username"] = "Jane"
	// 	userinfo[2]["age"] = "23"
	// 	userinfo[2]["gender"] = "female"
	// }

	// if userinfo[3] == nil {
	// 	userinfo[3] = make(map[string]string)
	// 	userinfo[3]["username"] = "Jone"
	// 	userinfo[3]["age"] = "24"
	// 	userinfo[3]["gender"] = "female"
	// }

	// for _, v := range userinfo {
	// 	fmt.Println(v)
	// }

	// for _, v := range userinfo {
	// 	for key, value := range v {
	// 		fmt.Printf("key : %v value : %v\n", key, value)
	// 	}
	// }

	var userinfo1 = make(map[string][]string)
	userinfo1["username"] = []string{"John", "Gibson", "Jane"}
	userinfo1["age"] = []string{"22", "23", "23"}
	userinfo1["gender"] = []string{"male", "male", "female"}

	fmt.Println(userinfo1)
	for _, v := range userinfo1 {
		for _, v := range v {
			fmt.Println(v)
		}
	}

	//Map and Slice both of them are called by reference type
	/*
		值類型:改變變量副本值的時候，不會改變變量本身的值(基本數據類型、陣列)
		引用類型:改變變量副本值的時候，會改變變量本身的值(Slice、Maps)
	*/

	var userinfo2 = make(map[string]string)
	userinfo2["username"] = "John"
	userinfo2["age"] = "22"
	userinfo3 := userinfo2
	userinfo3["age"] = "23"
	fmt.Println(userinfo2)
	fmt.Println(userinfo3)

	//Map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 98

	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice)
	sort.Ints(keySlice)
	fmt.Println(keySlice)

	for _, v := range keySlice {
		fmt.Printf("key = %v, value = %v\n", v, map1[v])
	}

	//數一句話中出現幾個單字
	var str = "How do you do how do you do"
	var strSlice = strings.Split(str, " ") // 3個單字
	fmt.Println(strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
