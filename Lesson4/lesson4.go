package main

import "fmt"

var x string = "Hello World MTHFCKR"

func main() {
	fmt.Println(x)

	var y string
	y = "first"
	fmt.Println(y)
	y = "second"
	fmt.Println(y)

	var z string
	z = "first "
	fmt.Println(z)
	z = z + "second "
	fmt.Println(z)
	z += "third"
	fmt.Println(z)

	var q string
	var w string
	q = "hello"
	w = "hello"
	fmt.Println(q == w)

	var e = 5
	fmt.Println(e)

	var r = "Hello GO!"
	fmt.Println(r)

	nameMy := "Igor"
	fmt.Println("My name is", nameMy)

}

func f() {
	fmt.Println(x)

}
