package main

import (
	"testing"

	utils "github.com/baspar/adventofcode2023/internal"
	"github.com/stretchr/testify/assert"
)

// var input = `1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`
var input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`


var inputs = map[string][2]string{
	input: {"142", ""},
}

// func TestPart1(t *testing.T) {
// 	assert := assert.New(t)
//
// 	for input, expectedRes := range inputs {
// 		d := &DayImpl{}
// 		d.Init(utils.SanitizeInput(input))
//
// 		res, err := d.Part1()
//
// 		assert.Equal(expectedRes[0], res)
// 		assert.Nil(err)
// 	}
// }

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
