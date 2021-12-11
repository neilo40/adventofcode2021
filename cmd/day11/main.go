package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type octopus struct {
	energy          int
	flashedThisStep bool
}

func (o *octopus) print() {
	fmt.Printf("%d", o.energy)
}

func main() {
	day11()
}

func day11() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	grid := [10][10]*octopus{}
	for i := 0; i < 10; i++ {
		scanner.Scan()
		row := scanner.Text()
		for j := 0; j < 10; j++ {
			grid[i][j] = &octopus{int(row[j] - '0'), false}
		}
	}

	flashes := 0
	stepCount := 1
	for {
		//printGrid(&grid)
		flashesThisStep := step(&grid)
		flashes += flashesThisStep
		if stepCount == 100 {
			log.Println("After 100 steps, there were", flashes, "flashes")
		}
		if flashesThisStep == 100 {
			log.Println("All octopuses flashed simultaneously at step", stepCount)
			break
		}
		stepCount++
	}
}

func step(grid *[10][10]*octopus) int {
	flashes := 0
	// add 1 to every octopus
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j].energy++
		}
	}

	for {
		// any that have energy >= 9 and not flashed this step, flash
		flashed := false
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if grid[i][j].energy > 9 && !grid[i][j].flashedThisStep {
					// increase all neighbouring octopuses energy
					for x := i - 1; x <= i+1; x++ {
						for y := j - 1; y <= j+1; y++ {
							if x < 0 || x > 9 || y < 0 || y > 9 {
								continue // out of grid bounds
							}
							grid[x][y].energy++
						}
					}
					grid[i][j].flashedThisStep = true
					flashes++
					flashed = true
				}
			}
		}
		// repeat until no more flashes
		if !flashed {
			break
		}
	}

	// reset for next step
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if grid[i][j].flashedThisStep {
				grid[i][j].energy = 0
			}
			grid[i][j].flashedThisStep = false
		}
	}

	return flashes
}

func printGrid(grid *[10][10]*octopus) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			grid[i][j].print()
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
