package main

import "fmt"

type Pos struct {
	X int
	Y int
}

type PosMemo struct {
	Pos
	M int
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

func TotalizeMByM(p Pos, grid map[Pos]int, m int) int {
	var total int
	for x := p.X; x < p.X+m; x++ {
		for y := p.Y; y < p.Y+m; y++ {
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

// Too slow
func SolvePartTwo(grid map[Pos]int) Pos {
	var currMax int
	var res Pos
	for x := 1; x < 301; x++ {
		for y := 1; y < 301; y++ {
			p := Pos{X: x, Y: y}
			mMax := min(301-x, 301-y)

			for m := 1; m < mMax+1; m++ {
				val := TotalizeMByM(p, grid, m)
				if val > currMax {
					res = p
					currMax = val
				}

			}
		}
	}

	return res
}

// Correct but slow
func SolvePartTwoMemo(grid map[Pos]int) PosMemo {
	// Build memo base case M = 1
	memo := make(map[PosMemo]int)
	for p, v := range grid {
		memo[PosMemo{
			Pos: p,
			M:   1,
		}] = v
	}

	var recurse func(p PosMemo) int
	recurse = func(p PosMemo) int {
		// Exit condition: Out of bounds
		if !(1 <= p.X && p.X < 301) || !(1 <= p.Y && p.Y < 301) {
			return 0
		}
		// Exit condition: we already solved
		if val, ok := memo[p]; ok {
			return val
		}

		// Sum up current row
		var val int
		for x := p.X; x < p.X+p.M; x++ {
			val += grid[Pos{X: x, Y: p.Y}]
		}

		// Sum up current column
		for y := p.Y; y < p.Y+p.M; y++ {
			val += grid[Pos{X: p.X, Y: y}]
		}

		// Subtract the doubly counted point
		val -= grid[Pos{X: p.X, Y: p.Y}]

		// Recurively call the diagonal
		val += recurse(PosMemo{
			Pos: Pos{
				X: p.X + 1,
				Y: p.Y + 1,
			},
			M: p.M - 1,
		})

		// Memoize it and return value
		memo[p] = val
		return memo[p]
	}

	// Loop through and solve
	var currMax int
	var res PosMemo
	for x := 1; x < 301; x++ {
		for y := 1; y < 301; y++ {
			p := Pos{X: x, Y: y}
			mMax := min(301-x, 301-y)
			for m := 1; m < mMax+1; m++ {
				pm := PosMemo{Pos: p, M: m}
				val := recurse(pm)
				if val > currMax {
					res = pm
					currMax = val
				}
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

	// // TLE
	// res2 := SolvePartTwo(mp)
	// fmt.Println(res2)

	// Works but slow
	res2 := SolvePartTwoMemo(mp)
	fmt.Println(res2)

}
