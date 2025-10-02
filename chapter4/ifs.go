package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 4-5
	//n := rand.Intn(10)
	//if n == 0 {
	//	fmt.Println("That's too low")
	//} else if n > 5 {
	//	fmt.Println("That's too big:", n)
	//} else {
	//	fmt.Println("That's a good number:", n)
	//}

	// 4-6
	fmt.Println("-----")
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	fmt.Println("asdasd", n)
}
