package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Alpha struct {
	Upper byte
	Lower byte
}

func SolvePartOne(data []byte) int {
	var stack []byte

	for _, b := range data {

		if len(stack) > 0 && (stack[len(stack)-1]-b == 32 || b-stack[len(stack)-1] == 32) {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, b)

	}

	return len(stack)

}

func SolvePartTwo(data []byte) int {

	// Create alpha pairs to exclude
	var alphas []Alpha
	var i byte = 65
	alphas = append(alphas, Alpha{
		Upper: i,
		Lower: i + 32,
	})
	for range 25 {
		i++
		alphas = append(alphas, Alpha{
			Upper: i,
			Lower: i + 32,
		})
	}

	var filteredData [][]byte
	for _, alpha := range alphas {
		filteredData = append(filteredData, FilterData(data, alpha))
	}

	shortestPolymer := len(data)
	for _, d := range filteredData {
		shortestPolymer = min(shortestPolymer, SolvePartOne(d))
	}

	return shortestPolymer

}

func FilterData(data []byte, exclude Alpha) []byte {
	var res []byte
	for _, b := range data {
		if b == exclude.Lower || b == exclude.Upper {
			continue
		}
		res = append(res, b)
	}

	return res
}

func main() {

	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	res := SolvePartOne(data)
	fmt.Println(res)

	res2 := SolvePartTwo(data)
	fmt.Println(res2)

}
