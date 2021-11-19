//http://golang-book.ru/
package main

import "fmt"

//func main() {
//	i := 1
//	for i <= 10 {
//		fmt.Println(i)
//		i += 1
//
//
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
			switch i {
			case 0:
				fmt.Println("Zero")
			case 1:
				fmt.Println("One")
			case 2:
				fmt.Println("Two")
			case 3:
				fmt.Println("Three")
			case 4:
				fmt.Println("Four")
			case 5:
				fmt.Println("Five")
			case 6:
				fmt.Println("Six")
			default:
				fmt.Println("Unknown Number")
			}
		} else {
			fmt.Println(i, "odd")
			switch i {
			case 0:
				fmt.Println("Zero")
			case 1:
				fmt.Println("One")
			case 2:
				fmt.Println("Two")
			case 3:
				fmt.Println("Three")
			case 4:
				fmt.Println("Four")
			case 5:
				fmt.Println("Five")
			case 6:
				fmt.Println("Six")
			default:
				fmt.Println("Unknown Number")
			}
		}
	}
	//Answer1: output - "Small)
	//Answer2: go to file Answer2-3.go
	//Answer3: go to file Answer2-3.go
}
