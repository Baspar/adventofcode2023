package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	utils "github.com/baspar/adventofcode2023/internal"
)

const (
	SEED_TO_SOIL = iota
	SOIL_TO_FERTILIZER
	FERTILIZER_TO_WATER
	WATER_TO_LIGHT
	LIGHT_TO_TEMPERATURE
	TEMPERATURE_TO_HUMIDITY
	HUMIDITY_TO_LOCATION
)

type Translation struct {
	From int
	To   int
	Jump int
}

func (t Translation) Map(source int) (bool, int) {
	if t.From <= source && source < t.To {
		return true, source + t.Jump
	}
	return false, source
}

type DayImpl struct {
	Seeds          []int
	TranslationMap map[int][]Translation
}

func (d *DayImpl) Init(lines []string) error {
	matchDigit := regexp.MustCompile(`\d+`)
	for _, seed := range matchDigit.FindAllString(lines[0], -1) {
		s, _ := strconv.Atoi(seed)
		d.Seeds = append(d.Seeds, s)
	}

	d.TranslationMap = make(map[int][]Translation)
	translationIndex := SEED_TO_SOIL
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			i += 1
			translationIndex++
			continue
		}
		var sourceFrom, targetFrom, n int
		fmt.Sscanf(lines[i], "%d %d %d", &targetFrom, &sourceFrom, &n)
		d.TranslationMap[translationIndex] = append(d.TranslationMap[translationIndex], Translation{
			From: sourceFrom,
			To:   sourceFrom + n,
			Jump: targetFrom - sourceFrom,
		})
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	var minLocation = math.MaxInt
	for _, seed := range d.Seeds {
		for layer := SEED_TO_SOIL; layer <= HUMIDITY_TO_LOCATION; layer++ {
			for _, translation := range d.TranslationMap[layer] {
				if valid, nextVal := translation.Map(seed); valid {
					seed = nextVal
					break
				}
			}
		}

		minLocation = min(minLocation, seed)
	}
	return fmt.Sprint(minLocation), nil
}
func (d *DayImpl) Part2() (string, error) {
	return "", nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
