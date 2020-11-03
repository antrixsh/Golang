package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64
	color   string

	geometry struct {
		area      float64
		perimeter float64
	}
}

func main() {
	var rect rectangle
	rect.length = 10.0
	rect.breadth = 20.0
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}
