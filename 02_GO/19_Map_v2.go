package main

import "fmt"

func main() {
	// var userinfo = make([]map[string]string, 3, 3)
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

	// for _, v := range userinfo {
	// 	for key, value := range v {
	// 		fmt.Printf("key : %v value : %v\n", key, value)
	// 	}
	// }

	// var userinfo1 = make(map[string][]string)
	// userinfo1["username"] = []string{"John", "Gibson", "Jane"}
	// userinfo1["age"] = []string{"22", "23", "23"}
	// userinfo1["gender"] = []string{"male", "male", "female"}

	// fmt.Println(userinfo1)
	// for _, v := range userinfo1 {
	// 	for _, v := range v {
	// 		fmt.Println(v)
	// 	}
	// }

	//Map and Slice both of them are call b reference type
	var userinfo2 = make(map[string]string)
	userinfo2["username"] = "John"
	userinfo2["age"] = "22"
	userinfo3 := userinfo2
	userinfo3["age"] = "23"
	fmt.Println(userinfo2)
	fmt.Println(userinfo3)


	23 重看
}