package main

import "fmt"

type Pos struct {
	X int
	Y int
}

func CreateGrid(serialNumber int) map[Pos]int {
	mp := make(map[Pos]int)
	for x := 1; x < 301; x++ {
		for y := 1; y < 301; y++ {
			fmt.Println(x, y)
		}
	}

	return mp

}

func main() {
	serialNumber := 8

	CreateGrid(serialNumber)

}
