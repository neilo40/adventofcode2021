package main

import (
	"bufio"
	"log"
	"os"
)

type runeStack struct {
	stack []rune
}

func (rs *runeStack) push(r rune) {
	rs.stack = append(rs.stack, r)
}

func (rs *runeStack) pop() rune {
	r := rs.stack[len(rs.stack)-1]
	rs.stack = rs.stack[0:]
	return r
}

func main() {
	day10_1()
}

func day10_1() {
	f, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	matchingRunes := map[rune]rune{']': '[', ')': '(', '}': '{', '>': '<'}
	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	total := 0
	for scanner.Scan() {
		inChunk := map[rune]int{'[': 0, '(': 0, '{': 0, '<': 0}
		line := scanner.Text()
		for i, r := range line {
			switch {
			case r == '[' || r == '(' || r == '{' || r == '<':
				inChunk[r]++
			case r == ']' || r == ')' || r == '}' || r == '>':
				openingRune := matchingRunes[r]
				if inChunk[openingRune] > 0 {
					inChunk[openingRune]--
				} else {
					// unmatched closing rune
					log.Printf("Found unmatched %c in %s at position %d\n", r, line, i)
					total += scores[r]
				}
			}
		}
	}

	log.Println("Total score for corrupted lines is", total)
}
