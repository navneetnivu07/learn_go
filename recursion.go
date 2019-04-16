package main

import "fmt"

func main() {
	fmt.Println("hello")

	val := factorial(5)
	fmt.Println(val)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}
