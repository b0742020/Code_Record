package main

import "fmt"

//struct 大寫完Public 小寫為Private

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) PersonInfo() {
	fmt.Printf("Name : %s, Age : %d, Gender : %s\n", p.Name, p.Age, p.Gender)
}

func (p *Person) SetInfo(name string, age int, gender string) {
	p.Name = name
	p.Age = age
	p.Gender = gender
}

func main() {

	//1.
	var p1 Person
	p1.Name = "John"
	p1.Age = 30
	p1.Gender = "Male"
	fmt.Printf("value %v type :%T\n", p1, p1)
	fmt.Printf("value %#v type :%T\n", p1, p1)

	//2.
	var p2 = new(Person)
	p2.Name = "Jane"
	p2.Age = 25
	p2.Gender = "Female"
	(*p2).Gender = "Male"
	fmt.Printf("value %#v type :%T\n", p2, p2)

	//3.
	var p3 = &Person{}
	p3.Name = "Tom"
	p3.Age = 40
	p3.Gender = "Male"
	fmt.Printf("value %#v type :%T\n", p3, p3)

	//4.
	var p4 = Person{
		Name:   "Jack",
		Age:    35,
		Gender: "Male",
	}
	fmt.Printf("value %#v type :%T\n", p4, p4)

	p4.PersonInfo()
	p4.SetInfo("Jerry", 20, "Male")
	p4.PersonInfo()
}
