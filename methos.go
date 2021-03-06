package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("=================================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "-----------------------------------------"
}

func main() {
	e := Employee{
		FirstName: "Abhishek",
		LastName:  "Sharma",
		Email:     "abhi@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000,
				HRA:   6000,
				TA:    3000,
			},
			Salary{
				Basic: 17000,
				HRA:   7000,
				TA:    4000,
			},
		},
	}
	fmt.Println(e.EmpInfo())
}
