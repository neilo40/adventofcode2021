package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type number struct {
	marked bool
	value  int
}

type board struct {
	locations  [5][5]*number
	lastCalled int
	hasWon     bool
}

func main() {
	day4()
}

func day4() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	calledNumbers := strings.Split(scanner.Text(), ",")

	scanner.Scan() // pop the empty line before the first board

	boardLines := make([]string, 0, 5)
	boards := make([]*board, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, createBoard(boardLines))
			boardLines = make([]string, 0, 5)
		} else {
			boardLines = append(boardLines, line)
		}
	}
	boards = append(boards, createBoard(boardLines)) // save the last board which is not followed by a newline

	log.Println("Playing", len(boards), "boards")

	winCount := 1
	for _, n := range calledNumbers {
		num, _ := strconv.Atoi(n)
		playedBoards := 0
		for _, b := range boards {
			if b.hasWon {
				continue
			}
			playedBoards++
			b.markNumber(num)
			if b.isWinner() {
				log.Printf("Final score (#%d) is %d\n", winCount, b.finalScore())
				winCount++
			}
		}
		if playedBoards == 0 {
			return
		}
	}
}

func createBoard(lines []string) *board {
	locations := [5][5]*number{}
	for i, l := range lines {
		for j, n := range strings.Fields(l) {
			numValue, _ := strconv.Atoi(n)
			locations[i][j] = &number{false, numValue}
		}
	}
	return &board{locations, 0, false}
}

func (b *board) isWinner() bool {
	// rows
	markedNumbers := 0
	for _, r := range b.locations {
		for _, n := range r {
			if n.marked {
				markedNumbers++
			}
		}
		if markedNumbers == 5 {
			b.hasWon = true
			return true
		}
		markedNumbers = 0
	}

	//columns
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if b.locations[i][j].marked {
				markedNumbers++
			}
		}
		if markedNumbers == 5 {
			b.hasWon = true
			return true
		}
		markedNumbers = 0
	}

	return false
}

func (b *board) markNumber(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.locations[i][j].value == number {
				b.locations[i][j].marked = true
				b.lastCalled = number
				return
			}
		}
	}
}

func (b *board) finalScore() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.locations[i][j].marked {
				sum += b.locations[i][j].value
			}
		}
	}
	return sum * b.lastCalled
}
