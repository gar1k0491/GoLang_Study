//http://golang-book.ru/
package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	fmt.Println(f1())

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func f1() int {
	return f2()
}
func f2() int {
	return 1
}

//Answer1: fucn sum(summ []float64) float64
