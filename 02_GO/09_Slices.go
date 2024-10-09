package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
}
