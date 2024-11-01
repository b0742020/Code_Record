package main

import (
	"fmt"
)

func main() {

	//1.make create map type's variable
	var userinfo = make(map[string]string)

	userinfo["name"] = "John"
	userinfo["age"] = "30"
	userinfo["gender"] = "Male"

	fmt.Println(userinfo["name"])
	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["gender"])
	fmt.Println(userinfo)

	//2.2nd way to make create map type's variable
	var userinfo2 = map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}

	fmt.Println(userinfo2["name"])
	fmt.Println(userinfo2["age"])
	fmt.Println(userinfo2["gender"])
	fmt.Println(userinfo2)

	//3. 3rd way to make create map type's variable

	userinfo3 := map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}

	fmt.Println(userinfo3["name"])
	fmt.Println(userinfo3["age"])
	fmt.Println(userinfo3["gender"])
	fmt.Println(userinfo3)
	//2nd and 3rd are equal

	//4. iterate over map
	var userinfo4 = map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}
	for k, v := range userinfo4 {
		fmt.Printf("key : %v, value : %v\n", k, v)
	}

	//5. create and modify map's key and value
	var userinfo5 = map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}
	fmt.Println(userinfo5)
	userinfo5["name"] = "Gibson"
	fmt.Println(userinfo5)

	//6. search and get map's key and value
	var userinfo6 = map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}
	fmt.Println(userinfo6["age"])
	//v, ok := userinfo6["name"] //if key not exist, ok will be false
	v, ok := userinfo6["xxxx"]
	fmt.Println(v, ok)

	//7. delete map's key and value
	var userinfo7 = map[string]string{
		"name":   "John",
		"age":    "30",
		"gender": "Male",
	}
	delete(userinfo7, "age")
	fmt.Println(userinfo7)
}
