package main

func filter(array []int, function func(x int, res []int)) []int {
	result := []int{}

	for i := 0; i < len(array); i++ {
		function(array[i], result)
	}

	return result
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filter(array, func(x int, res []int) {
		if (x % 2) == 0 {
			res = append(res, x)
		}
	})

	//or
	filter(array, func(x int, res []int) {
		if x < 5 {
			res = append(res, x)
		}
	})
}
