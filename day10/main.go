package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

type Point struct {
	Position Pos
	Velocity Pos
}

func (p *Point) Move() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

	fmt.Println(p.Position.X, p.Position.Y)
}

type Grid struct {
	XMax    int
	YMax    int
	Display [][]byte
}

func (g *Grid) ClearDisplay() {
	for y := range g.Display {
		for x := range y {
			g.Display[y][x] = '.'
		}
	}

}

func (g *Grid) UpdateDisplay(points []Point) {
	for _, point := range points {
		x := point.Position.X
		y := point.Position.Y
		if 0 <= x && x < g.XMax && 0 <= y && y < g.YMax {
			g.Display[y][x] = '#'
		}
	}
}

func (g Grid) WriteToFile() {
	var data []byte
	for _, row := range g.Display {
		data = append(data, row...)
		data = append(data, '\n')
	}

	err := os.WriteFile("output.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateGrid(xMax int, yMax int) Grid {

	var display [][]byte
	var row []byte
	for range xMax {
		row = append(row, '.')
	}

	for range yMax {
		display = append(display, slices.Clone(row))
	}

	return Grid{
		Display: display,
		XMax:    xMax,
		YMax:    yMax,
	}

}

func CreatePoint(vectors []string) (Point, error) {
	var res Point
	posVec := vectors[0]
	velVec := vectors[1]

	posVecArr := strings.Split(posVec, ",")
	velVecArr := strings.Split(velVec, ",")

	posX, err := strconv.Atoi(strings.TrimSpace(posVecArr[0]))
	if err != nil {
		return res, err
	}
	posY, err := strconv.Atoi(strings.TrimSpace(posVecArr[1]))

	if err != nil {
		return res, err
	}
	velX, err := strconv.Atoi(strings.TrimSpace(velVecArr[0]))
	if err != nil {
		return res, err
	}
	velY, err := strconv.Atoi(strings.TrimSpace(velVecArr[1]))
	if err != nil {
		return res, err
	}

	res = Point{Position: Pos{X: posX, Y: posY},
		Velocity: Pos{X: velX, Y: velY}}
	return res, nil

}

func CreatePoints(data []byte) ([]Point, error) {

	var res []Point

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		var vectors []string
		var curr string
		for _, char := range string(entry) {
			if char == '<' {
				curr = ""
				continue
			}
			if char == '>' {
				vectors = append(vectors, curr)
				curr = ""
				continue
			}
			curr += string(char)

		}
		pt, err := CreatePoint(vectors)
		if err != nil {
			return res, err
		}
		res = append(res, pt)
	}

	return res, nil

}

func SolvePartOne(points []Point) {

	grid := CreateGrid(100, 100)

	for range 3 {
		for _, p := range points {
			p.Move()
		}
	}

	grid.UpdateDisplay(points)

	grid.WriteToFile()

}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	points, err := CreatePoints(data)
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(points)

}
