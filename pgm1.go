package main

import "fmt"

func main() {
	var n int
	var temp int = 1
	fmt.Println("Enter no of rows")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d", temp)
			temp++
		}
		fmt.Println(" ")
	}

}
