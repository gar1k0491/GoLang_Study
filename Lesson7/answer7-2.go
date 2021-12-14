package main

import "fmt"

func main() {
	fmt.Println("Enter to get half of number:")

	var input float64
	fmt.Scanf("%d", &input)

	var output float64 = input / 2
	for output {
		if output%2 == 0 {
			fmt.Println(output, "True")
		} else {
			fmt.Println(output, "False")
		}

	}

	fmt.Println(output)
}
