package main

import "fmt"

func main() {

	fmt.Println("Maps")

	Student := make(map[string]int)
	Student["nivu"] = 23
	Student["shree"] = 21

	fmt.Println(Student["nivu"])
	fmt.Println(len(Student))

	Student["vish"] = 22
	fmt.Println(len(Student))

	delete(Student, "shree")
	fmt.Println(Student)

	fmt.Println(len(Student))

superhero:
	map[string]map[string]string{}

}
