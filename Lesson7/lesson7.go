//http://golang-book.ru/
package main

import "fmt"

func average(ww []float64) float64 {
	total := 0.0
	for _, v := range ww {
		total += v
	}
	return total / float64(len(ww))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
