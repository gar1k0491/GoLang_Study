package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	a := "hello world"
	b := 42
	fmt.Println(a)
	fmt.Println(b)
	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = "WTF"
	fmt.Println(a)

	g := &b
	*g = b / 2
	fmt.Println(b)
}
