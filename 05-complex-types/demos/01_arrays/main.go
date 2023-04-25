package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "Chai Tea"

	fmt.Println(arr)

	arr2 := arr

	arr2[2] = "Chai Latte"

	fmt.Println(arr, arr2)
}
