package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func main() {
	day5_1()
	day5_2()
}

func day5_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	lines := make([]line, 0)
	max := 0 // use this to size the grid array
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		values := [4]int{}
		for i := 0; i < 4; i++ {
			values[i], _ = strconv.Atoi(matches[0][i+1]) // matches[0][0] is the original string
			if values[i] > max {
				max = values[i]
			}
		}

		if values[0] == values[2] || values[1] == values[3] {
			// horizontal or vertical line
			startPoint := point{values[0], values[1]}
			endPoint := point{values[2], values[3]}
			lines = append(lines, line{startPoint, endPoint})
		}
	}

	max++
	grid := make([][]int, max)
	for i := range grid {
		grid[i] = make([]int, max)
	}

	for _, l := range lines {
		if l.start.x == l.end.x {
			//vertical
			if l.start.y > l.end.y {
				for y := l.start.y; y >= l.end.y; y-- {
					grid[l.start.x][y]++
				}
			} else {
				for y := l.start.y; y <= l.end.y; y++ {
					grid[l.start.x][y]++
				}
			}
		} else {
			//horizontal
			if l.start.x > l.end.x {
				for x := l.start.x; x >= l.end.x; x-- {
					grid[x][l.start.y]++
				}
			} else {
				for x := l.start.x; x <= l.end.x; x++ {
					grid[x][l.start.y]++
				}
			}
		}
	}

	overlapCount := 0
	for _, r := range grid {
		for _, p := range r {
			if p > 1 {
				overlapCount++
			}
		}
	}

	log.Println("Count of overlapping points (horiz and vert):", overlapCount)
}

func day5_2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	lines := make([]line, 0)
	max := 0 // use this to size the grid array
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		values := [4]int{}
		for i := 0; i < 4; i++ {
			values[i], _ = strconv.Atoi(matches[0][i+1]) // matches[0][0] is the original string
			if values[i] > max {
				max = values[i]
			}
		}

		if values[0] == values[2] || values[1] == values[3] || // horiz or vert
			math.Abs(float64(values[0]-values[2])) == math.Abs(float64(values[1]-values[3])) { // diagonal
			startPoint := point{values[0], values[1]}
			endPoint := point{values[2], values[3]}
			lines = append(lines, line{startPoint, endPoint})
		}
	}

	max++
	grid := make([][]int, max)
	for i := range grid {
		grid[i] = make([]int, max)
	}

	for _, l := range lines {
		if l.start.x == l.end.x {
			//vertical
			if l.start.y > l.end.y {
				for y := l.start.y; y >= l.end.y; y-- {
					grid[l.start.x][y]++
				}
			} else {
				for y := l.start.y; y <= l.end.y; y++ {
					grid[l.start.x][y]++
				}
			}
		} else if l.start.y == l.end.y {
			//horizontal
			if l.start.x > l.end.x {
				for x := l.start.x; x >= l.end.x; x-- {
					grid[x][l.start.y]++
				}
			} else {
				for x := l.start.x; x <= l.end.x; x++ {
					grid[x][l.start.y]++
				}
			}
		} else {
			//diagonal
			y := l.start.y
			if l.start.x > l.end.x {
				for x := l.start.x; x >= l.end.x; x-- {
					if l.start.y > l.end.y {
						grid[x][y]++
						y--
					} else {
						grid[x][y]++
						y++
					}
				}
			} else {
				for x := l.start.x; x <= l.end.x; x++ {
					if l.start.y > l.end.y {
						grid[x][y]++
						y--
					} else {
						grid[x][y]++
						y++
					}
				}
			}
		}
	}

	overlapCount := 0
	for _, r := range grid {
		for _, p := range r {
			if p > 1 {
				overlapCount++
			}
		}
	}

	log.Println("Count of overlapping points (horiz, vert, and diagonal):", overlapCount)
}
