package main

import "fmt"

func main() {
	i := 1

	// for loop one way
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// for loop two way

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

}
