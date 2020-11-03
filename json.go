package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {
	json_string := `
	{
		"firstname": "Hemant",
		"lastname": "Kumar",
		"city": "pune"
	}`
	//In Golang Terminology call marshal the process of generating a JSON from string from a DS,
	//and unmarshal the act of parsing JSON to a DS.

	emp1 := new(Employee)
	json.Unmarshal([]byte(json_string), emp1)
	fmt.Println(emp1)
	emp2 := new(Employee)
	emp2.FirstName = "Govind"
	emp2.LastName = "Sharma"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)

}
