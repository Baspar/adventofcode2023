package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/baspar/adventofcode2023/internal"
)

type Color string

func ColorFromString(color string) Color {
	switch color {
	case string(Red):
		return Red
	case string(Green):
		return Green
	case string(Blue):
		return Blue
	}
	panic(fmt.Sprintf("Color '%s' unrecognized", color))
}

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

type CubeSet map[Color]int

type Game struct {
	Id       int
	CubeSets []CubeSet
}
type DayImpl struct {
	Games []Game
}

func (d *DayImpl) Init(lines []string) error {
	rGame := regexp.MustCompile(`^Game (\d+): (.*)$`)
	for _, line := range lines {
		matches := rGame.FindStringSubmatch(line)
		if matches == nil {
			return fmt.Errorf("Couldn't extract game id from line '%s'", line)
		}

		gameId, _ := strconv.Atoi(matches[1])
		game := Game{Id: gameId}
		for _, _cubeSet := range strings.Split(matches[2], "; ") {
			cubeSet := make(CubeSet)
			for _, _cubes := range strings.Split(_cubeSet, ", ") {
				nbCubes, _ := strconv.Atoi(strings.Split(_cubes, " ")[0])
				colorCubes := ColorFromString(strings.Split(_cubes, " ")[1])
				cubeSet[colorCubes] += nbCubes
			}
			game.CubeSets = append(game.CubeSets, cubeSet)
		}
		d.Games = append(d.Games, game)
	}
	return nil
}
func (d *DayImpl) Part1() (string, error) {
	sum := 0

GameLoop:
	for _, game := range d.Games {
		for _, cubeSet := range game.CubeSets {
			if cubeSet[Red] > 12 {
				continue GameLoop
			}
			if cubeSet[Green] > 13 {
				continue GameLoop
			}
			if cubeSet[Blue] > 14 {
				continue GameLoop
			}
		}
		sum += game.Id
	}
	return fmt.Sprintf("%d", sum), nil
}
func (d *DayImpl) Part2() (string, error) {
	sum := 0
	for _, game := range d.Games {
		maxCubes := make(CubeSet)
		for _, cubeSet := range game.CubeSets {
			maxCubes[Red] = max(maxCubes[Red], cubeSet[Red])
			maxCubes[Green] = max(maxCubes[Green], cubeSet[Green])
			maxCubes[Blue] = max(maxCubes[Blue], cubeSet[Blue])
		}
		sum += maxCubes[Red]*maxCubes[Green]*maxCubes[Blue]
	}
	return fmt.Sprintf("%d", sum), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
