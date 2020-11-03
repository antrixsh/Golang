package main

import "fmt"

func main() {
	//var declares 1 or more variables

	var a = "initial" //implicit variable
	fmt.Println(a)

	// You can declare multiple variables at once

	var b, c int = 1, 2 //Explicit Variable
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are xero-values
	var e int
	fmt.Println(e)

	// The Syntax := is shorthand for declaring and initializing a variable.
	f := "apple"
	fmt.Println(f)
}
