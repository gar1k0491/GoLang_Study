package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	///////////////////////////////////////////
	u := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, v := range u {
		total += v
	}
	fmt.Println(total / float64(len(u)))

	///////////////////////////////////////////

}
