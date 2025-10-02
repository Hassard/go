package main

import "fmt"

func main() {
	// Page 48
	fmt.Println("Page 48")
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(string(b))

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2, s3, s4)

	//Page 49
	fmt.Println("Page 49")
	var ss string = "Hello, 🌞"
	var ss2 string = ss[5:8]
	var ss3 string = ss[:6]
	var ss4 string = ss[7:]
	fmt.Println(ss2, ss3, ss4, len(ss))

	var bs []byte = []byte(ss)
	var rs []rune = []rune(ss)
	fmt.Println(bs)
	fmt.Println(rs)
}
