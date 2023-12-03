package main

import (
	"fmt"

	utils "github.com/baspar/adventofcode2023/internal"
)

type DayImpl struct {
	Grid [][]byte
}

func (d *DayImpl) isNumber(x, y int) bool {
	return '0' <= d.Grid[y][x] && d.Grid[y][x] <= '9'
}

func (d *DayImpl) symbolNeighbors(x, y int) (symbolNeighbors []*byte) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// Don't check current cell
			if dx == 0 && dy == 0 {
				continue
			}

			i, j := x+dx, y+dy
			if i < 0 || j < 0 || j >= len(d.Grid) || i >= len(d.Grid[j]) || // Is out of bounce
				d.isNumber(i, j) || // Is a number
				d.Grid[j][i] == '.' { // Empty
				continue
			}

			symbolNeighbors = append(symbolNeighbors, &d.Grid[j][i])
		}
	}
	return
}

func (d *DayImpl) Init(lines []string) error {
	d.Grid = nil
	for _, line := range lines {
		d.Grid = append(d.Grid, []byte(line))
	}
	return nil
}

func (d *DayImpl) Part1() (string, error) {
	sum := 0
	for y, row := range d.Grid {
		accumulatedVal, hasAnySymbolNeighbor := 0, false
		for x, c := range row {
			if d.isNumber(x, y) {
				accumulatedVal = 10*accumulatedVal + int(c-'0')
				hasAnySymbolNeighbor = hasAnySymbolNeighbor || len(d.symbolNeighbors(x, y)) > 0
				continue
			}

			if hasAnySymbolNeighbor {
				sum += accumulatedVal
			}

			accumulatedVal, hasAnySymbolNeighbor = 0, false
		}

		if hasAnySymbolNeighbor {
			sum += accumulatedVal
		}
	}
	return fmt.Sprint(sum), nil
}

func (d *DayImpl) Part2() (string, error) {
	gears := make(map[*byte][]int)
	for y, row := range d.Grid {
		accumulatedVal, currentValNei := 0, make(map[*byte]bool)
		for x, c := range row {
			if d.isNumber(x, y) {
				accumulatedVal = 10*accumulatedVal + int(c-'0')
				for _, nei := range d.symbolNeighbors(x, y) {
					if *nei == '*' {
						currentValNei[nei] = true
					}
				}
				continue
			}

			for gear := range currentValNei {
				gears[gear] = append(gears[gear], accumulatedVal)
			}

			accumulatedVal, currentValNei = 0, make(map[*byte]bool)
		}

		for gear := range currentValNei {
			gears[gear] = append(gears[gear], accumulatedVal)
		}
	}
	sum := 0
	for _, parts := range gears {
		if len(parts) != 2 {
			continue
		}

		sum += parts[0] * parts[1]
	}
	return fmt.Sprint(sum), nil
}

func main() {
	utils.Run(&DayImpl{}, false)
}
