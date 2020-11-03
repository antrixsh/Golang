package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "Adit"
	s[1] = "Abhishek"
	s[2] = "Hitesh"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "Keerthi")
	s = append(s, "Kushagra", "Pratibha")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	fmt.Println("cpy:", c)
	copy(c, s)
	fmt.Println("cpys", c)

	l := s[1:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	m := []string{"a", "b", "c"}
	fmt.Println("dcl:", m)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}
