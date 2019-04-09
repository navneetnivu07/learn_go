package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	name person
	lc   bool
}

type moneyPenny struct {
	person
	lc bool
}

// methods

//func (r receiver) identifier(parms) (return(s)) {code}

func (s secretAgent) speak() {
	fmt.Println("My name is ", s.name)
}

func (p person) speak() {
	fmt.Println("My person name is ", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I'm human")
}

func main() {
	sa1 := secretAgent{}
	sa1.name.first = "nivu"
	sa1.name.last = "mali"
	sa1.lc = true

	sa2 := moneyPenny{
		person: person{
			"pri",
			"pak",
		},
		lc: true,
	}

	p1 := person{
		"Dr.",
		"Nivu",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(sa2)
	bar(sa1)
	fmt.Println(p1)
	p1.speak()
	//bar(p1)
}
