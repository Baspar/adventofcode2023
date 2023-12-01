package main

import (
	"fmt"
	"regexp"

	utils "github.com/baspar/adventofcode2023/internal"
)

type DayImpl struct {
	Lines []string
}

func (d *DayImpl) Init(lines []string) error {
	d.Lines = lines
	return nil
}
func valueOf(val []byte) int {
	switch string(val) {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}

	return int(val[0] - '0')
}

func (d *DayImpl) sumWith(r *regexp.Regexp) int {
	sum := 0
	for _, l := range d.Lines {
		var (
			line = []byte(l)
			startSearchAt = 0
			first, last []byte
		)
		for {
			matchLocation := r.FindIndex(line[startSearchAt:])
			if matchLocation == nil {
				break
			}

			from, to := matchLocation[0]+startSearchAt, matchLocation[1]+startSearchAt
			last = line[from:to]
			if first == nil {
				first = last
			}

			startSearchAt = from + 1
		}
		value := 10*valueOf(first) + valueOf(last)
		sum += value
	}
	return sum
}

func (d *DayImpl) Part1() (string, error) {
	sum := d.sumWith(regexp.MustCompile(`\d`))
	return fmt.Sprint(sum), nil
}
func (d *DayImpl) Part2() (string, error) {
	sum := d.sumWith(regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`))
	return fmt.Sprint(sum), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
