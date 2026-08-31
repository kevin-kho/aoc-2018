package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Node struct {
	ChildCount    int
	MetadataCount int
}

func CreateNode(childCount int, metadataCount int) Node {
	return Node{
		ChildCount:    childCount,
		MetadataCount: metadataCount,
	}
}

func GetIntArray(data []byte) ([]int, error) {
	var res []int

	for s := range strings.SplitSeq(string(data), " ") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}

	return res, nil

}

func CreateNodes(intArr []int) {

}

func SumIntArr(intArr []int) int {
	var res int
	for _, item := range intArr {
		res += item
	}
	return res

}

func SolvePartOne(intArr []int) int {
	// Totalize metadata
	var total int

	var stack []*Node

	for len(intArr) > 0 {
		childCt := intArr[0]
		metadataCt := intArr[1]
		intArr = intArr[2:]

		stack = append(stack, &Node{
			ChildCount:    childCt,
			MetadataCount: metadataCt,
		})

		for len(stack) > 0 && stack[len(stack)-1].ChildCount == 0 {
			// Increment Count
			for range stack[len(stack)-1].MetadataCount {
				total += intArr[0]
				intArr = intArr[1:]
			}

			// Pop the used up Node
			stack = stack[:len(stack)-1]

			// Decrement top of stack
			if len(stack) > 0 {
				stack[len(stack)-1].ChildCount -= 1
			}
		}
	}

	return total

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := GetIntArray(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intArr)
	fmt.Println(res)

}
