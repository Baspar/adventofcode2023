package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2023/internal"
	"github.com/stretchr/testify/assert"
)

var input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var inputs = map[string][2]string{
	input: {"4361", "467835"},
}

func TestPart1(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		d := &DayImpl{}
		d.Init(utils.SanitizeInput(input))

		res, err := d.Part1()

		assert.Equal(expectedRes[0], res)
		assert.Nil(err)
	}
}

func TestPart2(t *testing.T) {
	assert := assert.New(t)

	for input, expectedRes := range inputs {
		d := &DayImpl{}
		d.Init(utils.SanitizeInput(input))

		res, err := d.Part2()

		assert.Equal(expectedRes[1], res)
		assert.Nil(err)
	}
}
