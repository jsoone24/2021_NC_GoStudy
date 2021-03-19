package main

import "fmt"

func fiboFac() func() int {
	first, second := 1, 1

	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	fibo := fiboFac()

	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}
}
