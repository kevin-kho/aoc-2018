package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Entry struct {
	Timestamp time.Time
	Guard     int
	Action    string
}

func ParseData(data []byte) ([]Entry, error) {
	var res []Entry

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStrArr := strings.Split(string(entry), " ")

		date := entryStrArr[0]
		date = strings.TrimPrefix(date, "[")

		hourMinute := entryStrArr[1]
		hourMinute = strings.TrimSuffix(hourMinute, "]")

		dateTime := fmt.Sprintf("%v %v", date, hourMinute)
		layout := "2006-01-02 15:04"
		t, err := time.Parse(layout, dateTime)
		if err != nil {
			return res, err
		}

		var guard int
		var action string

		// Case Guard and action
		if entryStrArr[2] == "Guard" {
			guard, err = strconv.Atoi(strings.TrimPrefix(entryStrArr[3], "#"))
			if err != nil {
				return res, err
			}

			action = strings.Join(entryStrArr[4:], " ")

			// Case: only action
		} else {
			action = strings.Join(entryStrArr[2:], " ")
		}

		e := Entry{
			Timestamp: t,
			Guard:     guard,
			Action:    action,
		}

		res = append(res, e)

	}

	slices.SortFunc(res, func(a, b Entry) int {
		return a.Timestamp.Compare(b.Timestamp)

	})

	return res, nil
}

func SolvePartOne(entries []Entry) float64 {

	guardSleep := make(map[int]float64)

	currGuard := 0
	var sleepStart time.Time
	for _, entry := range entries {
		if entry.Action == "begins shift" {
			currGuard = entry.Guard
		}

		if entry.Action == "falls asleep" {
			sleepStart = entry.Timestamp
		}

		if entry.Action == "wakes up" {
			sleepMinutes := entry.Timestamp.Sub(sleepStart).Minutes()
			guardSleep[currGuard] += sleepMinutes
		}

	}

	var guard int
	var maxSleep float64

	for g, sleepDuration := range guardSleep {
		if sleepDuration > maxSleep {
			maxSleep = sleepDuration
			guard = g
		}
	}

	return maxSleep * float64(guard)

}

func main() {
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	entries, err := ParseData(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(entries)
	fmt.Println(res)

}
