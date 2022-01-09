package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"
	//a[2] = "228" // Error, invalid array index
	fmt.Println(a)
	fmt.Println(a[1])

	num := [...]int{10, 20, 30, 40}
	num[3] = 999
	//num[4] = 555 // Error, invalid array index

}
