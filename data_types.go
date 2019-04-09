package main

import "fmt"

func main() {
	var age int = 40

	var fa float64 = 1.66666

	randNum := 1

	fmt.Println(age, fa, randNum)

	var num = 1.000
	var n2 = 0.999

	fmt.Println(num - n2)

	const pi = 3.1444

	fmt.Println(pi)

	var (
		va = 2
		vb = 3
	)

	fmt.Println(va, vb)

	var myChar = 'a'
	var myStr = "double quotes"
	var myStr2 = `backslash quotes`

	fmt.Println(myChar, myStr, myStr2)

	var mybool = true

	// format boolean
	fmt.Printf("%t \n", mybool)

	// format float
	fmt.Printf("%.1f \n", pi)

	// get data type
	fmt.Printf("%T \n", pi)

	// int, binary, char
	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 100)
	fmt.Printf("%x \n", 100)
	fmt.Printf("%e \n", pi)
}
