package main

import "fmt"

//const constWidth = 320 // Constants u can`t change!

func main() {
	/*fmt.Println(constWidth)
	s2, s3, s4 := test()
	sum, i := test2()
	fmt.Println(test1())
	fmt.Println(s2 ,s3, s4)
	fmt.Println(sum)*/

	v := 0
	for v < 1000 {

		if v == 100 {
			fmt.Println("v is 100")
		}

		v += 10
	}
}

/*type s struct {

}
	a string
	b string
	c string

}

func test() (a, c, b string) {
	a = "Hello "
	c = "all "
	b = "World!"
	return a, c, b

}

func test1() string {
	return "empty"

}

func test2() (sum, i int) {
	sum = 0
	for i := 0; i < 10; i++ {
		sum += i

	}
	return
}
*/
