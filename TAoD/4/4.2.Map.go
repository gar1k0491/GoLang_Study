package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{} // OR pointsMap := make(map[string]Point)
	otherPoint := make(map[int]Point)
	pointsMap["key1"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["key1"])
	fmt.Println("-------------------------------------")

	otherPoint[1] = Point{33, 44}
	fmt.Println(otherPoint)
	fmt.Println(otherPoint[1])
	fmt.Println("-------------------------------------")

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{}
		fmt.Println("oneMorePointsMap is nil")
	}
	oneMorePointsMap["a"] = Point{1511, 2511}
	fmt.Println(oneMorePointsMap["a"])
	fmt.Println("-------------------------------------")

	key := "s"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("Key=%s - exist in map \n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("Key=%s does't exist in map \n", key)
		fmt.Println(value)
	}

}
