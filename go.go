package main

import "fmt"

func main() {
	fmt.Println("hello")

	var student [10]string

	student[0] = "Niu"
	student[1] = "b"
	student[2] = "c"

	temp := [3]int{1, 2, 3}

	changeAryValPoint(&student[1])

	for i := 0; i < len(student); i++ {
		fmt.Println(student[i])

	}

	for _, value := range temp {
		fmt.Println(value)
	}
}

func changeAryValPoint(point *string) {
	fmt.Println(point)
	*point = "aa"
}
