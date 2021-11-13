package main

import "fmt"

func main() {
	//Srings
	fmt.Println("1 + 1 =", 1+1.0)
	fmt.Println(len("Thirty five symbol length of string"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World\n")
	//Boolean types
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Answer1: BIN type
	//Answer2: 255
	//Answer3: 1364067664
	fmt.Println("32132 × 42452 =", 32132*42452)
	//Answer4: String is a line of symbols. For find length use operation "len"
	//Answer5: false || false || true = true
	fmt.Println((true && false) || (false && true) || !(false && false))

}
