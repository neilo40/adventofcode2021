package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	day7_1()
	day7_2()
}

func day7_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	positionStrings := strings.Split(scanner.Text(), ",")

	positions := make([]int, 0, len(positionStrings))
	for _, p := range positionStrings {
		positionInt, _ := strconv.Atoi(p)
		positions = append(positions, positionInt)
	}

	minFuel := math.MaxInt
	for i := range positions {
		fuel := 0
		for _, p := range positions {
			fuel += int(math.Abs(float64(p - i)))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	log.Println("Minimum amount of fuel needed (human engine model) is", minFuel)
}

func day7_2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	positionStrings := strings.Split(scanner.Text(), ",")

	positions := make([]int, 0, len(positionStrings))
	for _, p := range positionStrings {
		positionInt, _ := strconv.Atoi(p)
		positions = append(positions, positionInt)
	}

	minFuel := math.MaxInt
	for i := range positions {
		fuel := 0
		for _, p := range positions {
			fuel += fuelUsed(p, i)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	log.Println("Minimum amount of fuel needed (crab engine model) is", minFuel)
}

func fuelUsed(start int, end int) int {
	distance := int(math.Abs(float64(start - end)))
	totalFuelUsed := 0
	for i := 1; i <= distance; i++ {
		totalFuelUsed += i
	}
	return totalFuelUsed
}
