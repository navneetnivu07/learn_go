package main

import "fmt"

func main() {

	// arrays
	var favNum [2]int

	favNum[0] = 1
	favNum[1] = 5

	fmt.Println(favNum[0])

	favNum2 := [5]int{5, 4, 3, 2, 1}

	for i, value := range favNum2 {
		fmt.Println(value, i)
	}

	for _, value := range favNum2 {
		fmt.Println(value)
	}

	numSlice := favNum2[3:5] // [:2] and [3:]

	fmt.Println("numSlice")

	for _, value := range numSlice {
		fmt.Println(value)
	}

	// create slice without defined set of values
	fmt.Println("numSlice3")
	numSlice3 := make([]int, 5, 10) // we have room for 10 values

	numSlice3[3] = 5

	for _, value := range numSlice3 {
		fmt.Println(value)
	}

	// copy slices
	fmt.Println("copy slices")

	numSlice2 := []int{5, 4, 3, 2, 1}

	copy(numSlice3, numSlice2)

	for _, value := range numSlice3 {
		fmt.Println(value)
	}

	// append to array
	fmt.Println("append")
	numSlice3 = append(numSlice3, 10, 11)

	fmt.Println(numSlice3[5])

	/*************************/

	var student [10]string

	student[0] = "Niu"
	student[1] = "b"
	student[2] = "c"

	temp := [3]int{1, 2, 3}

	changeAryValPoint(&student[1])

	for i := 0; i < len(student); i++ {
		fmt.Println(student[i])
	}

	// iterator
	for _, value := range temp {
		fmt.Println(value)
	}

	// iterator with index
	for i, value := range temp {
		fmt.Println(i, value)
	}
}

func changeAryValPoint(point *string) {
	fmt.Println(point)
	*point = "aa"
}
