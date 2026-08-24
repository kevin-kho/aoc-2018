package main

import (
	"bytes"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

func SolvePartOne(coords []Pos) {

}

func CreateCoords(data []byte) ([]Pos, error) {
	var res []Pos

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryArr := bytes.Split(entry, []byte{',', ' '})
		x := string(entryArr[0])
		y := string(entryArr[1])

		xInt, err := strconv.Atoi(x)
		if err != nil {
			return res, err
		}

		yInt, err := strconv.Atoi(y)
		if err != nil {
			return res, err
		}

		res = append(res, Pos{X: xInt, Y: yInt})

	}

	return res, nil
}

func main() {
	data, err := common.ReadInput("inputExample.txt")
	if err != nil {
		log.Fatal()
	}
	data = common.TrimNewLineSuffix(data)

	coords, err := CreateCoords(data)
	if err != nil {
		log.Fatal(err)
	}

	SolvePartOne(coords)
}
