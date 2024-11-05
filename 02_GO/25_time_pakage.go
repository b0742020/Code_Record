package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, min, sec)
	str := t.Format("2006-01-02 15:04:05")
	fmt.Println(str)

	//生成時間戳
	timeObj := time.Now()
	unixtime := timeObj.Unix()
	fmt.Println(unixtime)
	timeObj2 := time.Unix(unixtime, 0)
	var str2 = timeObj2.Format("2006-01-02 15:04:05")
	fmt.Println(str2)
	//定時器
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		fmt.Println(t)
		n--
		if n == 0 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("aaa1")
	time.Sleep(time.Second)
	fmt.Println("aaa2")
	time.Sleep(time.Second)
	fmt.Println("aaa3")
	time.Sleep(time.Second)
	fmt.Println("aaa4")
	time.Sleep(time.Second)
}
