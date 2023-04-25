package main

import (
	"fmt"
)

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m)

	fmt.Println(m["coffee"])
	delete(m, "tea")
	fmt.Println(m)

	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	fmt.Println(m["tea"])
	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)

}
