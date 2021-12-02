package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2_1()
	day2_2()
}

func day2_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	horiz := 0
	depth := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horiz += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	log.Println("part 1 result:", horiz*depth)
}

func day2_2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	horiz := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horiz += amount
			depth += (aim * amount)
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	log.Println("part 2 result:", horiz*depth)
}
