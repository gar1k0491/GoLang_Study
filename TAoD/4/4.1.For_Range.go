package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c"}
	for i, l := range arr {
		fmt.Println(i)
		fmt.Println(l)
		fmt.Println("-------------------------------------")
	}
}
