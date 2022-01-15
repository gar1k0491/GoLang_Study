package main

import "fmt"

func main() {

	letters := []string{"a", "b", "c"}
	letters[0] = "Changed"
	letters = append(letters, "d_added")
	letters = append(letters, "e_added", "f_added")
	fmt.Println(letters)
	fmt.Println(fmt.Sprintf("letters len: %d", len(letters)))
	fmt.Println(fmt.Sprintf("letters cap: %d", cap(letters)))
	fmt.Println("-------------------------------------")

	createSlice := make([]string, 3)
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	fmt.Println(createSlice)
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	fmt.Println("-------------------------------------")
	//	createSlice[3] = "NEW_ELEMENT" // error
	createSlice = append(createSlice, "NEW_ELEMENT")
	fmt.Println(createSlice)
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
	fmt.Println("-------------------------------------")

	animalsArr := [4]string{
		"dog",
		"cat",
		"mouse",
		"rat",
	}

	fmt.Println(animalsArr)
	var a []string = animalsArr[0:2]
	fmt.Println(a)
	b := animalsArr[1:3]
	fmt.Println(b)
	fmt.Println("-------------------------------------")

	numArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := numArr[:5]
	fmt.Println(t)
	tt := numArr[5:]
	fmt.Println(tt)
	ttt := numArr[:]
	fmt.Println(ttt)

}
