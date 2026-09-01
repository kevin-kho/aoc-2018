package main

import "fmt"

type Pos struct {
	X int
	Y int
}

func GetHundredDigit(num int) int {
	if num < 100 {
		return 0
	}

	num = num / 100
	digit := num % 10

	return digit

}

func CreateGrid(serialNumber int) map[Pos]int {
	mp := make(map[Pos]int)
	for x := 1; x < 301; x++ {
		for y := 1; y < 301; y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += serialNumber
			powerLevel *= rackId

			digit := GetHundredDigit(powerLevel)
			digit -= 5

			mp[Pos{
				X: x,
				Y: y,
			}] = digit
		}
	}

	return mp

}

func TotalizeThreeByThree(p Pos, grid map[Pos]int) int {
	var total int
	for x := p.X; x < p.X+3; x++ {
		for y := p.Y; y < p.Y+3; y++ {
			total += grid[Pos{
				X: x,
				Y: y,
			}]

		}
	}

	return total

}

func SolvePartOne(grid map[Pos]int) Pos {
	var currMax int
	var res Pos

	// Top left value
	for x := 1; x < 299; x++ {
		for y := 1; y < 299; y++ {

			p := Pos{X: x, Y: y}

			val := TotalizeThreeByThree(p, grid)

			if val > currMax {
				res = p
				currMax = val
			}
		}
	}

	return res

}

func main() {
	serialNumber := 8868

	mp := CreateGrid(serialNumber)

	res := SolvePartOne(mp)
	fmt.Println(res)

}
