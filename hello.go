package main

import "fmt"

func main() {
	fmt.Println("hello")

	// pointers
	pval := 1
	fmt.Println(pval)
	fmt.Println(&pval)

	changeValue(&pval)

	fmt.Println(pval)

}

func changeValue(point *int) {
	*point = 2
}
