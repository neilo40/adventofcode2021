package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	day6(80)
	day6(256) // doesn't complete.  try recursive function?
}

func day6(iterations int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	ages := strings.Split(scanner.Text(), ",")

	fish := make([]int, 0, len(ages))
	for _, a := range ages {
		fish = append(fish, int([]rune(a)[0]-'0'))
	}

	for i := 0; i < iterations; i++ {
		log.Println("Day", i)
		newFish := make([]int, 0, len(fish)/2)
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 6
				newFish = append(newFish, 8)
			} else {
				fish[j]--
			}
		}
		fish = append(fish, newFish...)
	}

	log.Println("There are", len(fish), "lanternfish after", iterations, "days")
}
