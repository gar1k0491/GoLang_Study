//http://golang-book.ru/
package main

import (
	"fmt"
)

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

	const t string = "Hello"
	//t = "Hell"
	fmt.Println(t)

	fmt.Println("Enter a number (x2): ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	//Answer1: var or :=
	//Answer2: 6
	s := 12
	s += 3
	fmt.Println(s)
	//Answer3: View area is {}
	//Answer4: var editable, const uneditable

	//Answer5:
	fmt.Println("Input temperature: ")
	var faren float64
	fmt.Scanf("%f", &faren)
	cels := (faren - 32) * 5 / 9
	fmt.Println(cels)

	//Answer6:
	fmt.Println("Input distance: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet / 3.281
	fmt.Println(meters)

}
