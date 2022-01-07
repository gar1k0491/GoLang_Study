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
	defer fmt.Println("DONE 3!")
	defer fmt.Println("DONE 2!")
	defer fmt.Println("DONE 1!")
	v := 0
	for v < 47 {

		switch i := isTest(v); i {
		case 3:
			fmt.Println("v = 3")
		case 4:
			fmt.Println("v = 4")
		default:
			fmt.Println(fmt.Sprintf("unknown v, v = %d", v))

		}

		v++
	}

}

func isTest(v int) int {
	if v == 3 {
		return 3
	} else if v == 4 {
		return 4
	}
	return 5

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
