package main

import "fmt"

func main() {
	fmt.Printf("Введиче число, которое хотите возвести в квадрат: ")
	var i int
	fmt.Scan(&i)
	square := i * i
	fmt.Println("Введенное вами число:", i)
	fmt.Println("Результат:", square)

}
