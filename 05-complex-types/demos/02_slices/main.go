package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "Chai Tea"
	fmt.Println(s)

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)

	s = slices.Delete(s, 1, 2)
	fmt.Println(s)

	s2 := s

	s[2] = "Chai Latte"

	fmt.Println(s, s2)
}
