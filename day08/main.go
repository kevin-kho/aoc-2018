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
	ChildrenIdx   []int // Part1: Metadata values, Part2: Child Indexes if not a leaf node
	Children      []*Node
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

func SolvePartOneAlt(intArr []int) int {
	var res int

	root := BuildTree(intArr)

	var dfs func(node *Node) int
	dfs = func(node *Node) int {

		// Exit condition: hit leaf node
		if len(node.Children) == 0 {
			return SumIntArr(node.ChildrenIdx)
		}

		// Add the Node
		childSum := SumIntArr(node.ChildrenIdx)

		// Add all the children's
		for _, child := range node.Children {
			childSum += dfs(child)
		}

		return childSum

	}

	res = dfs(root)

	return res

}

func BuildTree(intArr []int) *Node {
	var root *Node
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

			c := intArr[0:stack[len(stack)-1].MetadataCount]
			intArr = intArr[stack[len(stack)-1].MetadataCount:]
			stack[len(stack)-1].ChildrenIdx = c

			// Pop the used up Node
			if len(stack) == 1 {
				root = stack[len(stack)-1]
			}
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Decrement top of stack
			if len(stack) > 0 {
				stack[len(stack)-1].ChildCount -= 1
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, node)
			}
		}
	}
	return root

}

func SolvePartTwo(intArr []int) int {

	root := BuildTree(intArr)

	var dfs func(node *Node) int
	dfs = func(node *Node) int {

		// case: hit leaf node, sum the metadata
		if len(node.Children) == 0 {
			return SumIntArr(node.ChildrenIdx)
		}

		var childrenValues []int
		for _, childNode := range node.Children {
			childrenValues = append(childrenValues, dfs(childNode))
		}

		var total int
		// Indexes are 1-based
		// Add up the values only if index is within range
		for _, i := range node.ChildrenIdx {
			if i-1 < len(childrenValues) {
				total += childrenValues[i-1]
			}
		}

		return total

	}

	res := dfs(root)

	return res
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

	resAlt := SolvePartOneAlt(intArr)
	fmt.Println(resAlt)

	res2 := SolvePartTwo(intArr)
	fmt.Println(res2)

}
