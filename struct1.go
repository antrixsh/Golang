package main

import (
	"fmt"
	"reflect"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	var rect1 = rectangle{10.5, 25.10, "red"}
	fmt.Println(reflect.TypeOf(rect1))
	fmt.Println(reflect.ValueOf(rect1).Kind())

	rect2 := rectangle{length: 10.5, breadth: 25.10, color: "red"}
	fmt.Println(reflect.TypeOf(rect2))
	fmt.Println(reflect.ValueOf(rect2).Kind())

	rect3 := new(rectangle)
	fmt.Println(reflect.TypeOf(rect3))
	fmt.Println(reflect.ValueOf(rect3).Kind())

	var rect4 = &rectangle{}
	fmt.Println(reflect.TypeOf(rect4))
	fmt.Println(reflect.ValueOf(rect4).Kind())
}
