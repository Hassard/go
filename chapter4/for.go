package main

import "fmt"

func main() {
	// 4-8
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----")

	// 4-9
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	fmt.Println("-----")
	// 4-10
	//for {
	//	fmt.Println("Hello")
	//}

	fmt.Println("-----")
	// 4-12
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

}
