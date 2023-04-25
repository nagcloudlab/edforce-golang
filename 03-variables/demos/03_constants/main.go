package main

import "fmt"

func main() {
	const a = 42
	const b float32 = 3

	var i int = a
	fmt.Println(i)
	var f32 float32 = a
	fmt.Println(f32)
	f32 = b
	fmt.Println(f32)

	// var i2 int = b
	// fmt.Println(i2)

	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e
		f = iota
		g
		h = 10 * iota
	)

	fmt.Println(d, e, f, g, h)

}
