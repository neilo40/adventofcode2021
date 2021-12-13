package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

type coord struct {
	x int
	y int
}

func (c *coord) str() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type fold struct {
	direction string // x or y
	location  int
}

func coordFromString(coordString string) *coord {
	parts := strings.Split(coordString, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return &coord{x, y}
}

func foldFromString(foldString string) fold {
	fields := strings.Fields(foldString)
	foldDetails := strings.Split(fields[2], "=")
	location, _ := strconv.Atoi(foldDetails[1])
	return fold{foldDetails[0], location}
}

func main() {
	day13_1()
}

func day13_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// parse coordinates
	coords := make([]*coord, 0, 1125)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // gap between coords and fold instructions
		}
		coords = append(coords, coordFromString(line))
	}

	// parse fold instructions
	folds := make([]fold, 0, 12)
	for scanner.Scan() {
		folds = append(folds, foldFromString(scanner.Text()))
	}

	doFold(coords, folds[0])
	visibleDots := countUniqueCoords(coords)
	log.Println("There are", visibleDots, "visible dots after the first fold")
	for i := 1; i < len(folds); i++ {
		doFold(coords, folds[i])
	}
	print(coords)
}

func doFold(coords []*coord, f fold) {
	for _, c := range coords {
		// vertical fold
		if f.direction == "x" {
			// if coord is to the right of the fold location
			if c.x > f.location {
				// subtract 2 x distance from fold
				c.x = c.x - (2 * (c.x - f.location))
			}
		} else {
			// horizontal fold
			// if coord is below the fold location
			if c.y > f.location {
				// subtract 2 x distance from fold
				c.y = c.y - (2 * (c.y - f.location))
			}
		}
	}
}

func countUniqueCoords(coords []*coord) int {
	u := make(map[string]bool)
	for _, c := range coords {
		u[c.str()] = true
	}
	return len(u)
}

func print(coords []*coord) {
	dc := gg.NewContext(90, 20)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	for _, c := range coords {
		// doesn't render well at edge of canvas, shift down and right by 5
		// doesn't render well at 1:1 with points of radius 1, double coords to spread
		dc.DrawPoint((float64(c.x)*2)+5, (float64(c.y)*2)+5, 1.0)
		dc.Fill()
	}
	dc.SavePNG("result.png")
}
