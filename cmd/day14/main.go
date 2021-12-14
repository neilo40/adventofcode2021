package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type rule struct {
	pair   string
	insert string
}

func main() {
	day14(10)
	day14(40) // too slow - needs rethink
}

func day14(iterations int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan() // empty line
	rules := make([]rule, 0, 100)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		rules = append(rules, rule{parts[0], parts[1]})
	}

	newPolymer := ""
	for i := 0; i < iterations; i++ {
		log.Println("Iteration", i)
		newPolymer = grow(polymer, rules)
		polymer = newPolymer
	}

	answer := analysePolymer(polymer)
	log.Println("Answer after", iterations, "iterations:", answer)
}

func grow(polymer string, rules []rule) string {
	newPolymer := ""
	for i := range polymer {
		if i == len(polymer)-1 {
			// special case - last element
			break
		}
		pair := fmt.Sprintf("%c%c", polymer[i], polymer[i+1])
		matched := false
		for _, r := range rules {
			if pair == r.pair {
				if i == 0 {
					newPolymer += fmt.Sprintf("%c%s%c", polymer[i], r.insert, polymer[i+1])
				} else {
					newPolymer += fmt.Sprintf("%s%c", r.insert, polymer[i+1])
				}
				matched = true
				break
			}
		}
		if !matched {
			if i == 0 {
				newPolymer += fmt.Sprintf("%c%c", polymer[i], polymer[i+1])
			} else {
				newPolymer += fmt.Sprintf("%c", polymer[i+1])
			}
		}
	}
	return newPolymer
}

func analysePolymer(polymer string) int {
	elementFreq := make(map[rune]int)
	for _, e := range polymer {
		_, ok := elementFreq[e]
		if ok {
			elementFreq[e]++
		} else {
			elementFreq[e] = 1
		}
	}

	max := 0
	min := math.MaxInt
	for _, v := range elementFreq {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}
