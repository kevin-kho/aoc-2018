package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func GetIntSlc(data []byte) ([]int, error) {
	var res []int
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		val, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}
		res = append(res, val)
	}
	return res, nil
}

func SolvePartOne(intSlc []int) int {
	var res int

	for _, i := range intSlc {
		res += i
	}

	return res

}

func SolvePartTwo(intSlc []int) int {

	var curr int
	seen := make(map[int]bool)
	seen[curr] = true

	i := 0
	for {
		i = i % len(intSlc)
		curr += intSlc[i]
		if seen[curr] {
			return curr
		}
		seen[curr] = true
		i++

	}

}

func main() {

	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intSlc, err := GetIntSlc(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intSlc)
	fmt.Println(res)

	res2 := SolvePartTwo(intSlc)
	fmt.Println(res2)

}
