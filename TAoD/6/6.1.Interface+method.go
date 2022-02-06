package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2

}

type InterfaceHere interface {
	Sum() int
}

func main() {
	var a InterfaceHere
	sh := structHere{1, 2}
	os := otherStruct{3, 5}

	a = &sh
	fmt.Println(a.Sum())
	a = os
	fmt.Println(a.Sum())
	fmt.Println("-------------------------------------")

	fmt.Printf("(%v, %T)\n", a, a) // не изменяемый список, кортеж

}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B

}
