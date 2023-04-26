package main

import (
	"fmt"
)

func main() {
	testScores := map[string]float64 {	
		"Harry": 87.3,
		"Hermione": 105,
		"Ronald": 63.5,
		"Neville": 27,
	}
	c := clone(testScores)

	fmt.Println(c)
}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}