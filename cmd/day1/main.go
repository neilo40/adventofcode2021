package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	day1_1()
	day1_2()
}

func day1_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	increaseCount := 0
	previousDepth := 0
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		if previousDepth > 0 {
			if depth > previousDepth {
				increaseCount++
			}
		}
		previousDepth = depth
	}

	log.Println("Count of individual increases:", increaseCount)
}

func day1_2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var measurements []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, depth)
	}

	increaseCount := 0
	previousDepth := 0
	for i := range measurements {
		// exit if our last measurement will be out of bounds
		if i+2 > len(measurements)-1 {
			break
		}

		depth := measurements[i] + measurements[i+1] + measurements[i+2]
		if previousDepth > 0 {
			if depth > previousDepth {
				increaseCount++
			}
		}
		previousDepth = depth
	}

	log.Println("Count of group increases:", increaseCount)
}
