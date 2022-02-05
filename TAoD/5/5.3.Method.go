package main

import "fmt"

type Point struct {
	X, Y int
}

func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func movePointPtr(p *Point, x, y int) {
	p.X += x
	p.Y += y

}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(movePoint(p, 4, 2))
	fmt.Println(p)
	movePointPtr(&p, 4, 2)
	fmt.Println(p)

}
