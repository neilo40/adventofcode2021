package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
)

type runeStack struct {
	stack []rune
}

func (rs *runeStack) push(r rune) {
	rs.stack = append(rs.stack, r)
}

func (rs *runeStack) pop() (rune, error) {
	if len(rs.stack) == 0 {
		return ' ', errors.New("pop from empty stack")
	}
	r := rs.stack[len(rs.stack)-1]
	rs.stack = rs.stack[0 : len(rs.stack)-1]
	return r, nil
}

func main() {
	day10()
}

func day10() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	matchingRunes := map[rune]rune{']': '[', ')': '(', '}': '{', '>': '<'}
	corruptScores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	incompleteScores := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

	totalCorrupt := 0
	totalIncomplete := make([]int, 0)
	for scanner.Scan() {
		openingRuneStack := runeStack{}
		line := scanner.Text()
		corrupted := false
		for _, r := range line {
			switch {
			case r == '[' || r == '(' || r == '{' || r == '<':
				openingRuneStack.push(r)
			case r == ']' || r == ')' || r == '}' || r == '>':
				openingRune := matchingRunes[r]
				mustMatch, _ := openingRuneStack.pop()
				if openingRune != mustMatch {
					// unmatched closing rune
					totalCorrupt += corruptScores[r]
					corrupted = true
				}
			}
		}
		if !corrupted { //incomplete
			score := 0
			for {
				r, err := openingRuneStack.pop()
				if err != nil {
					break
				}
				score = score * 5
				score += incompleteScores[r]
			}
			totalIncomplete = append(totalIncomplete, score)
		}
	}

	// sort incomplete scores and find median
	sort.Ints(totalIncomplete)

	log.Println("Total score for corrupted lines is", totalCorrupt)
	log.Println("Total score for incomplete lines is", totalIncomplete[len(totalIncomplete)/2])
}
