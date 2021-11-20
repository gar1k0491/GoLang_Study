package main

import "fmt"

func main() {

	var p string
	p = "==============================="
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(p)

	/////////////////////////////////////////// ARRY

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
	fmt.Println(p)

	/////////////////////////////////////////// SLICE

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
	fmt.Println(p)

	/////////////////////////////////////////// MAP

	q := make(map[int]int)
	q[1] = 5
	fmt.Println(q[1])

	//elements := map[string]string{
	//	"H": "Hydrogen",
	//	"He": "Helium",
	//	"Li": "Lithium",
	//	"Be": "Beryllium",
	//	"B": "Boron",
	//	"C": "Carbon",
	//	"N": "Nitrogen",
	//	"O": "Oxygen",
	//	"F": "Fluorine",
	//	"Ne": "Neon",
	//}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if elements, ok := elements["H"]; ok {
		fmt.Println(elements["name"], elements["state"])
	}

	//fmt.Println(elements["B"])

	//name, ok := elements["F"]
	//fmt.Println(name, ok)

	//if name, ok := elements["F"]; ok {
	//	fmt.Println(name, ok)
	//}

	//Answer1:

}
