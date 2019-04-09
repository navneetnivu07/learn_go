package main

import "fmt"

func foo() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("foo")
}

func main(){
	foo()
}