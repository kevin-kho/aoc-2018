package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Entry struct {
	Id       string
	RuneFreq map[rune]int
}

func (e Entry) ContainsCount(ct int) bool {
	for _, val := range e.RuneFreq {
		if val == ct {
			return true
		}
	}
	return false
}

func CreateEntry(str string) Entry {

	mp := make(map[rune]int)
	for _, r := range str {
		mp[r]++
	}

	return Entry{
		Id:       str,
		RuneFreq: mp,
	}
}

func GetEntries(data []byte) []Entry {
	var res []Entry

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		e := CreateEntry(string(entry))
		res = append(res, e)
	}

	return res
}

func StringsDifferByOne(str0 string, str1 string) bool {
	i := 0
	var diffSeen bool
	for i < len(str0) {

		if str0[i] != str1[i] {
			if diffSeen {
				return false
			}
			diffSeen = true
		}

		i++

	}

	return diffSeen

}

func SolvePartOne(entries []Entry) int {
	var twoCount int
	var threeCount int
	for _, entry := range entries {

		if entry.ContainsCount(2) {
			twoCount++
		}

		if entry.ContainsCount(3) {
			threeCount++
		}

	}

	return twoCount * threeCount
}

func SolvePartTwo(entries []Entry) string {
	var res []byte

	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			diffByOne := StringsDifferByOne(entries[i].Id, entries[j].Id)
			if diffByOne {
				for z := range entries[i].Id {
					if entries[i].Id[z] == entries[j].Id[z] {
						res = append(res, entries[i].Id[z])
					}
				}
			}
		}
	}

	return string(res)
}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	entries := GetEntries(data)

	res := SolvePartOne(entries)
	fmt.Println(res)

	res2 := SolvePartTwo(entries)
	fmt.Println(res2)

}
