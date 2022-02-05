package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func main() {
	pointsMap := map[string]int{
		"xx": 14424,
		"yy": 2124,
	}

	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)
	fmt.Println("-------------------------------------")

	for k, v := range pointsMap {
		fmt.Printf("%s is ", k)
		fmt.Println(v)
	}

}
