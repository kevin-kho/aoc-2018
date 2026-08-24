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

type GuardSleep struct {
	Guard                    int
	SleepMinuteFreq          map[int]int
	MostFreqSleptMinute      int
	MostFreqSleptMinuteCount int
}

func (g *GuardSleep) GetHighestSleptMinute() {
	maxMinute := -1
	maxCt := -1
	for minute, ct := range g.SleepMinuteFreq {
		if ct > maxCt {
			maxMinute = minute
			maxCt = ct
		}
	}

	g.MostFreqSleptMinute = maxMinute
	g.MostFreqSleptMinuteCount = maxCt
}

func CreateGuardSleep(guard int, sleepMinutes []int) GuardSleep {
	mp := make(map[int]int)
	for _, minute := range sleepMinutes {
		mp[minute]++
	}

	return GuardSleep{
		Guard:           guard,
		SleepMinuteFreq: mp,
	}
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

func SolvePartOne(entries []Entry) int {

	// Determines guard that sleeps the most
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

	// Determines guard's most slept minute
	guardMinuteFreq := make(map[int]int)
	currGuard = 0
	for _, entry := range entries {
		if entry.Action == "begins shift" {
			currGuard = entry.Guard
		}

		if entry.Action == "falls asleep" && currGuard == guard {
			sleepStart = entry.Timestamp
		}

		if entry.Action == "wakes up" && currGuard == guard {
			sleepStartMinute := sleepStart.Minute()
			sleepMinutes := entry.Timestamp.Sub(sleepStart).Minutes()
			guardMinuteFreq[sleepStartMinute]++ // Time 0
			for range int(sleepMinutes) - 1 {
				sleepStartMinute++
				sleepStartMinute = sleepStartMinute % 60
				guardMinuteFreq[sleepStartMinute]++
			}
		}
	}

	maxMinute := 0
	maxMinuteCt := 0
	for minute, ct := range guardMinuteFreq {
		if ct > maxMinuteCt {
			maxMinuteCt = ct
			maxMinute = minute
		}

	}

	return guard * maxMinute

}

func SolvePartTwo(entries []Entry, minute int) int {

	guardSleepMinutes := make(map[int][]int)
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
			sleepStartMinute := sleepStart.Minute()
			sleepMinutes := entry.Timestamp.Sub(sleepStart).Minutes()

			var minutes []int
			for i := sleepStartMinute; i < sleepStartMinute+int(sleepMinutes); i++ {
				minutes = append(minutes, i%60)
			}

			guardSleepMinutes[currGuard] = append(guardSleepMinutes[currGuard], minutes...)

		}

	}

	var gs []GuardSleep
	for guard, minutes := range guardSleepMinutes {
		g := CreateGuardSleep(guard, minutes)
		g.GetHighestSleptMinute()
		gs = append(gs, g)
	}

	var res int
	var maxCt int
	for _, guardSleep := range gs {
		if guardSleep.MostFreqSleptMinuteCount > maxCt {
			maxCt = guardSleep.MostFreqSleptMinuteCount
			res = guardSleep.Guard * guardSleep.MostFreqSleptMinute
		}

	}

	return res

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

	res2 := SolvePartTwo(entries, 45)
	fmt.Println(res2)

}
