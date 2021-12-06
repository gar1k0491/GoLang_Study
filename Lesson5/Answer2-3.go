//Answer2-3:
//http://golang-book.ru/
package main

import "fmt"

func main() {

	for a := 1; a <= 100; a++ {
		//		if a%3 == 0 {
		//			fmt.Println(a,"Fizz")
		//		} else if a%5 == 0 {
		//			fmt.Println(a,"Buzz")
		//		}
		switch {
		case a%3 == 0 && a%5 == 0:
			fmt.Println(a, "FizzBuzz") //or case a%15 == 0: fmt.Print...
		case a%3 == 0:
			fmt.Println(a, "Fizz")
		case a%5 == 0:
			fmt.Println(a, "Buzz")
		default:
			fmt.Println(a)
		}

	}
}
