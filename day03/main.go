package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Claim struct {
	Id         string
	LeftOffset int // +x
	TopOffset  int // +y
	Width      int // x
	Length     int // y
}

type Coordinate struct {
	X int
	Y int
}

func (c Claim) GetCoordinates() []Coordinate {
	var res []Coordinate
	x := c.LeftOffset
	y := c.TopOffset

	for dx := range c.Width {
		for dy := range c.Length {
			c := Coordinate{
				X: x + dx,
				Y: y + dy,
			}
			res = append(res, c)

		}
	}

	return res

}

func GetClaims(data []byte) ([]Claim, error) {

	var res []Claim

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		strArr := strings.Split(string(entry), " ")

		id := strArr[0]

		// Offsets Parsing
		offsets := strArr[2]
		offsets = strings.TrimSuffix(offsets, ":")
		leftOffset := strings.Split(offsets, ",")[0]
		topOffset := strings.Split(offsets, ",")[1]

		leftOffsetInt, err := strconv.Atoi(leftOffset)
		if err != nil {
			return res, err
		}

		topOffsetInt, err := strconv.Atoi(topOffset)
		if err != nil {
			return res, err
		}

		// Dimensions Parsing
		dimensions := strArr[3]
		width := strings.Split(dimensions, "x")[0]
		widthInt, err := strconv.Atoi(width)
		if err != nil {
			return res, err
		}
		length := strings.Split(dimensions, "x")[1]
		lengthInt, err := strconv.Atoi(length)
		if err != nil {
			return res, err
		}

		c := Claim{
			Id:         id,
			LeftOffset: leftOffsetInt,
			TopOffset:  topOffsetInt,
			Width:      widthInt,
			Length:     lengthInt,
		}

		res = append(res, c)

	}

	return res, nil

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	claims, err := GetClaims(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(claims)

}
