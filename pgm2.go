package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var dob int
	var phone int
	var err error
	_, err = fmt.Scanln(&dob)
	_, err = fmt.Scanln(&phone)

	if err == nil {
		rand.Seed(time.Now().UnixNano())
		dob1 := []int{}
		phone1 := []int{}
		for dob != 0 {

			dobx := dob % 10
			dob = dob / 10
			dob1 = append(dob1, dobx)

		}

		for phone != 0 {

			phonex := phone % 10
			phone = phone / 10
			phone1 = append(phone1, phonex)

		}
		count1 := 0
		accNum := 0
		for count1 < 5 {
			accNum = accNum * 10
			accNum += dob1[rand.Intn(len(dob1)-1)]
			count1++
		}
		count2 := 0
		for count2 < 5 {
			accNum = accNum * 10
			accNum += phone1[rand.Intn(len(phone1)-1)]
			count2++
		}
		fmt.Println("Account No", accNum)
	}
}
