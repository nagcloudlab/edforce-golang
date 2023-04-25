package main

import "fmt"

func main() {
	s := "Hello, world!"

	p := &s

	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hello, Gophers!"

	fmt.Println(s, *p)

	p = new(string)

	fmt.Println(p, *p)

}
