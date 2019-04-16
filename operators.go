package main

import "fmt"

func main() {
	fmt.Println("hello")

	val := fact(5)
	fmt.Println(val)

}

func fact(n int) int {
	if n >= 1 {
		return n * fact(n-1)
	} else {
		return 1
	}
}
