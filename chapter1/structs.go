package main

import "fmt"

func main() {
	//Page 56/57
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}
	julia := person{
		"Julia",
		40,
		"cat",
	}
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(fred, bob, julia, beth)

	bob.name = "Bob"
	fmt.Println(bob.name)

}
