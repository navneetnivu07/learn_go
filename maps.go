package main

import "fmt"

func main() {

	fmt.Println("Maps")

	Student := make(map[string]int)
	Student["nivu"] = 23
	Student["neet"] = 21

	fmt.Println(Student["nivu"])
	fmt.Println(len(Student))

	Student["nav"] = 22
	fmt.Println(len(Student))

	delete(Student, "neet")
	fmt.Println(Student)

	fmt.Println(len(Student))

	superhero := map[string]map[string]string{
		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	fmt.Println(superhero)
	fmt.Println(superhero["Batman"]["RealName"])
	fmt.Println(superhero["Batman"]["City"])

	if temp, hero := superhero["Batman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])

	}

}
