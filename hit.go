package main

import "fmt"

func Numerics(a int, b int, c string) int {

	if c == "a" {
		return a + b
	} else if c == "b" {
		return a - b
	} else if c == "c" {
		return a * b
	} else if c == "d" {
		return a / b
	}
	return 8989898
}

func main() {

	var x int
	var y int
	var z string
	var err error
	_, err = fmt.Scanln(&x)
	_, err = fmt.Scanln(&y)
	_, err = fmt.Scanln(&z)

	if err == nil {
		res := Numerics(x, y, z)
		fmt.Println("result: ", res)
	}

}
