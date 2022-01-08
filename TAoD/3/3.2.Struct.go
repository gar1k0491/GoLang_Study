package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {

	p1 := Point{
		X: 33,
		Y: 440,
		S: "WoW!",
	}
	p1.method()

	structs()

}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "A",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.Y = 123
	fmt.Println(p1)

	p2 := Point{
		X: 450,
	}
	fmt.Println(p2)

	g := &p1
	fmt.Println(*g)
	fmt.Println(&g)
	fmt.Println(g)

}
