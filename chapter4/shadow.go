package main

import "fmt"

func main() {

	// 4-1
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// 4-2
	fmt.Println("-----")
	x2 := 10
	if x2 > 5 {
		x2, y := 5, 20
		fmt.Println(x2, y)
	}
	fmt.Println(x2)

	// 4-3
	//fmt.Println("-----")
	//x3 := 10
	//fmt.Println(x3)
	//fmt := "oops"
	//fmt.Println(fmt)
}
