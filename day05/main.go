package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(data []byte) int {
	var stack []byte

	for _, b := range data {

		if len(stack) > 1 && (stack[len(stack)-1]-b == 32 || b-stack[len(stack)-1] == 32) {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, b)

	}

	return len(stack)

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

}
