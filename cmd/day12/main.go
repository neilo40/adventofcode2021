package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type cave struct {
	small      bool
	name       string
	neighbours []*cave
}

func main() {
	caves := day12_1()
	day12_2(caves)
}

func day12_1() map[string]*cave {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	caves := map[string]*cave{}
	for scanner.Scan() {
		cavePairs := strings.Split(scanner.Text(), "-")
		// create the caves if they don't exist
		for i, c := range cavePairs {
			_, ok := caves[cavePairs[i]]
			if !ok {
				lower := strings.ToLower(c)
				small := (lower == c)
				caves[c] = &cave{small, c, []*cave{}}
			}
		}
		// join the caves together
		caves[cavePairs[0]].neighbours = append(caves[cavePairs[0]].neighbours, caves[cavePairs[1]])
		caves[cavePairs[1]].neighbours = append(caves[cavePairs[1]].neighbours, caves[cavePairs[0]])
	}

	totalPaths := walk([]*cave{}, caves["start"], 0)
	log.Println("There are", totalPaths, "unique paths through the caves in part 1")

	return caves
}

func walk(visited []*cave, currentCave *cave, uniquePaths int) int {
	if currentCave.name == "end" {
		return uniquePaths + 1
	}

	paths := uniquePaths
	for _, c := range currentCave.neighbours {
		lower := strings.ToLower(c.name)
		isLower := (lower == c.name)
		// keep going if cave is uppercase, or is lowercase and not been visited yet
		if !isLower || (isLower && !contains(visited, c)) {
			paths += uniquePaths + walk(append(visited, currentCave), c, uniquePaths)
		}
	}
	return paths
}

func contains(caves []*cave, c *cave) bool {
	for _, cav := range caves {
		if cav.name == c.name {
			return true
		}
	}
	return false
}

func day12_2(caves map[string]*cave) {
	totalPaths := walk2([]*cave{}, caves["start"], 0)
	log.Println("There are", totalPaths, "unique paths through the caves in part 2")
}

func walk2(visited []*cave, currentCave *cave, uniquePaths int) int {
	if currentCave.name == "end" {
		return uniquePaths + 1
	}

	paths := uniquePaths
	for _, c := range currentCave.neighbours {
		if shouldContinue(visited, currentCave, c) {
			paths += uniquePaths + walk2(append(visited, currentCave), c, uniquePaths)
		}
	}
	return paths
}

func shouldContinue(visited []*cave, currentCave *cave, neighbour *cave) bool {
	// do not visit start twice
	if neighbour.name == "start" {
		return false
	}

	// if uppercase, go for it
	nLower := strings.ToLower(neighbour.name)
	isNLower := (nLower == neighbour.name)
	if !isNLower {
		return true
	}

	visitCount := map[string]int{currentCave.name: 1}
	for _, vc := range visited {
		lower := strings.ToLower(vc.name)
		isLower := (lower == vc.name)
		if isLower {
			_, ok := visitCount[vc.name]
			if ok {
				visitCount[vc.name]++
			} else {
				visitCount[vc.name] = 1
			}
		}
	}

	_, ok := visitCount[neighbour.name]
	if ok {
		for _, v := range visitCount {
			if v > 1 {
				// lowercase but a cave has been visited twice
				return false
			}
		}
		// lowercase and no lowercase cave has been visited twice
		return true
	}

	// lowercase and not been visited - go for it
	return true
}
