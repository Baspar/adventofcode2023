package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2023/internal"
)

type ScratchCard struct {
	CardNumbers    map[int]bool
	WinningNumbers map[int]bool
}

func (c ScratchCard) NbWinningNumbers() (winningNumbers int) {
	for number := range c.CardNumbers {
		if c.WinningNumbers[number] {
			winningNumbers++
		}
	}

	return winningNumbers
}

type DayImpl struct {
	ScratchCards []ScratchCard
}

func (d *DayImpl) Init(lines []string) error {
	matchDigit := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")

		card := ScratchCard{make(map[int]bool), make(map[int]bool)}

		for _, number := range matchDigit.FindAllString(numbers[0], -1) {
			num, _ := strconv.Atoi(number)
			card.CardNumbers[num] = true
		}
		for _, number := range matchDigit.FindAllString(numbers[1], -1) {
			num, _ := strconv.Atoi(number)
			card.WinningNumbers[num] = true
		}

		d.ScratchCards = append(d.ScratchCards, card)
	}

	return nil
}
func (d *DayImpl) Part1() (string, error) {
	sum := 0
	for _, card := range d.ScratchCards {
		if score := card.NbWinningNumbers(); score > 0 {
			sum += 1 << (score - 1)
		}
	}
	return fmt.Sprint(sum), nil
}
func (d *DayImpl) Part2() (string, error) {
	totalNumberOfCards, numberOfCards := 0, make([]int, len(d.ScratchCards))
	for cardIndex, card := range d.ScratchCards {
		totalNumberOfCards += numberOfCards[cardIndex] + 1
		for cardCopyIndex := cardIndex + 1; cardCopyIndex < min(len(numberOfCards), cardIndex+card.NbWinningNumbers()+1); cardCopyIndex++ {
			numberOfCards[cardCopyIndex] += numberOfCards[cardIndex] + 1
		}
	}
	return fmt.Sprint(totalNumberOfCards), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
