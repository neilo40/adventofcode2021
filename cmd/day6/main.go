package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	//day6_alt(80)
	day6(256)
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
	ageBuckets := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
	for _, a := range ages {
		age := int([]rune(a)[0] - '0') // convert single char string to int
		ageBuckets[age]++
	}

	totalFish := 0
	for k, v := range ageBuckets {
		log.Println("iterating fish with starting age", k)
		totalFish += v * iterate(k, iterations, 0)
	}

	log.Println("There are", totalFish, "lanternfish after", iterations, "days")
}

// iterate a single fish to completion
func iterate(startingAge int, iterations int, fishPop int) int {
	fish := make([]int, 0, int(math.Pow(2.0, float64(iterations/8))))
	fish = append(fish, startingAge)
	for i := 0; i < iterations; i++ {
		log.Println("Day", i, ", fish population:", len(fish))
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j]--
			}
		}
	}
	return len(fish)
}
