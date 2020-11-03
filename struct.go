package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 25
	return &p
}

func main() {
	//This syntax creates a new struct
	fmt.Println(person{"ABC", 15})
	// we can name the fields when initializing a struct
	fmt.Println(person{name: "DEF", age: 16})
	// Omitted a field value
	fmt.Println(person{name: "GHI"})
}
