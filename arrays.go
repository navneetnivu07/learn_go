package main

import "fmt"

func main() {

	// arrays
	var favNum [2]int

	favNum[0] = 1
	favNum[1] = 5

	fmt.Println(favNum[0])

	favNum2 := [5]int{5, 4, 3, 2, 1}

	// iterator
	for _, value := range favNum2 {
		fmt.Println(value)
	}

	// iterator with index
	for i, value := range favNum2 {
		fmt.Println(value, i)
	}

	// array slice
	fmt.Println("original")
	fmt.Println(favNum2)

	sliced := favNum2[3:5] // [:2] and [3:]
	fmt.Println("sliced")
	fmt.Println(sliced)

	// create slice without defined set of values
	fmt.Println("sliceWithOutVal")
	sliceWithOutVal := make([]int, 5, 10) // we have room for 10 values
	fmt.Println(sliceWithOutVal)
	sliceWithOutVal[3] = 5
	fmt.Println(sliceWithOutVal)

	// copy slices
	fmt.Println("copy slices")
	numSlice2 := []int{5, 4, 3}
	copy(sliceWithOutVal, numSlice2) // overwrites values in position 0,1,2 in sliceWithOutVal
	fmt.Println(sliceWithOutVal)

	// append to array
	fmt.Println("append")
	sliceWithOutVal = append(sliceWithOutVal, 10, 11)
	fmt.Println(sliceWithOutVal)

	/*************************/

	// change array value using pointer
	var student [10]string

	student[0] = "Niu"
	student[1] = "b"
	student[2] = "c"

	fmt.Println(student[1])
	changeAryValPoint(&student[1])
	fmt.Println(student[1])

}

func changeAryValPoint(point *string) {
	fmt.Println(point)
	*point = "aa"
}
