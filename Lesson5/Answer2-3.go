//Answer2-3:
//http://golang-book.ru/
package main

import "fmt"

func main() {

	for a := 1; a <= 100; a++ {
		if a%3 == 0 {
			fmt.Println("Fizz")
		} else if a%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
