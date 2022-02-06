package main

import (
	"fmt"
)

func main() {
	var a interface{}
	a = "jelly"
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n\n", a, a)
	a = 42
	fmt.Println(a)
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Println("-------------------------------------")

	var b interface{} = "hello"
	s := b.(string)
	fmt.Println(s)

	s, ok := b.(string)
	fmt.Println(s, ok)
	fmt.Println("-------------------------------------")

	f, ok := b.(float32)
	fmt.Println(f, ok)
	fmt.Println("-------------------------------------")

	switch b.(type) {
	case int:
		fmt.Println("b - is int")
	case float32:
		fmt.Println("b - is float32")
	case string:
		fmt.Println("b - is string")
	default:
		fmt.Printf("unknown type. %T\n", b)

	}

}
