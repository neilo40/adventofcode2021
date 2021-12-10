package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	day9_1()
}

func day9_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	caveMap := make([][]int, 0, 102)
	edgeRow := make([]int, 102)
	for i := range edgeRow {
		edgeRow[i] = 10
	}
	caveMap = append(caveMap, edgeRow)
	for scanner.Scan() {
		rowText := scanner.Text()
		row := make([]int, 0, len(rowText)+2)
		row = append(row, 10)
		for _, r := range rowText {
			row = append(row, int(r-'0'))
		}
		row = append(row, 10)
		caveMap = append(caveMap, row)
	}
	caveMap = append(caveMap, edgeRow)

	lowPoints := make([]int, 0)
	for i := 1; i < len(caveMap)-1; i++ {
		for j := 1; j < len(caveMap[0])-1; j++ {
			if caveMap[i][j] >= caveMap[i-1][j] {
				continue
			}
			if caveMap[i][j] >= caveMap[i+1][j] {
				continue
			}
			if caveMap[i][j] >= caveMap[i][j-1] {
				continue
			}
			if caveMap[i][j] >= caveMap[i][j+1] {
				continue
			}
			lowPoints = append(lowPoints, caveMap[i][j])
		}
	}

	riskLevelSum := 0
	for _, lp := range lowPoints {
		riskLevelSum++
		riskLevelSum += lp
	}

	log.Println("Sum of risk levels is", riskLevelSum)
}
