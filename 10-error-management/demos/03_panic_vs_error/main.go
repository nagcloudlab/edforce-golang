package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide1(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result: ", result)

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("something went wrong: %v\n", msg)
		}
	}()
	divide2(10, 5)

	divide3(10, 5)

}

func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return l / r, nil

}

func divide2(l, r int) int {
	return l / r
}

func divide3(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
