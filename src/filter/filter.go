package main

import "fmt"

func filter(array []int, fn func(e int) bool) []int {
	var res []int = make([]int, len(array))

	for i, v := range array {
		if fn(v) {
			res[i] = v
		}
	}

	return res
}

func mapFunc(array []int, fn func(int) int) []int {
	for i, v := range array {
		array[i] = fn(v)
	}
	return array
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*
		add := mapFunc(array, func(e int) int {
			return e + 2
		})

		square := mapFunc(array, func(e int) int {
			return e * e
		})

		filter(array, func(e int) bool {
			return (e % 2) == 0
		})

		//or
		filter(array, func(e int) bool {
			return e < 5
		})
	*/

	res := filter(mapFunc(array, func(e int) int {
		return e + 10
	}), func(e int) bool {
		return (e % 5) < 2
	})

	fmt.Println(res)
}
