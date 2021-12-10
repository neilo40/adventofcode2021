package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

/* 7 seg arrangement
wireNames [7]rune
rune default value is 0

 0000
6    1
6    1
 2222
5    3
5    3
 4444

1 uses 2 segments 1,3
7 uses 3 segments 0,1,3
4 uses 4 segments 1,2,3,6
8 uses 7 segments 0,1,2,3,4,5,6

5 uses 5 segments 0,2,3,4,6
2 uses 5 segments 0,1,2,4,5
3 uses 5 segments 0,1,2,3,4

0 uses 6 segments 0,1,3,4,5,6
6 uses 6 segments 0,2,3,4,5,6
9 uses 6 segments 0,1,2,3,4,6

The only numbers that are ambiguous are 2,3,5 (5 segs) and 0,6,9 (6 segs)

Wow, this is a horrible and overly-complicated implementation...
*/

type segment struct {
	wireNames  [7]rune
	digitRunes [][]rune
}

func (s *segment) setWireNames(line string) {
	var one, four, seven string
	var maybeZeroSixNine []string
	var maybeTwoThreeFive []string

	for _, f := range strings.Fields(line) {
		switch len(f) {
		case 2:
			one = f
		case 3:
			seven = f
		case 4:
			four = f
		case 6:
			maybeZeroSixNine = append(maybeZeroSixNine, f)
		case 5:
			maybeTwoThreeFive = append(maybeTwoThreeFive, f)
		}
	}
	s.setSeg0Wire(one, seven)
	s.setSeg4Wire(four, maybeZeroSixNine)
	s.setSeg2Wire(one, maybeTwoThreeFive)
	s.setSeg5Wire(four, maybeTwoThreeFive)
	s.setSeg1Wire(one, maybeZeroSixNine)
	s.setSeg3Wire(maybeTwoThreeFive)
	s.setSeg6Wire(four)

	for i, r := range s.wireNames {
		fmt.Printf("(%d:%c), ", i, r)
	}

	// what runes comprise each digit
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[3], s.wireNames[4], s.wireNames[5], s.wireNames[6]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[1], s.wireNames[3]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[2], s.wireNames[4], s.wireNames[5]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[2], s.wireNames[3], s.wireNames[4]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[1], s.wireNames[2], s.wireNames[3], s.wireNames[6]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[2], s.wireNames[3], s.wireNames[4], s.wireNames[6]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[2], s.wireNames[3], s.wireNames[4], s.wireNames[5], s.wireNames[6]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[3]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[2], s.wireNames[3], s.wireNames[4], s.wireNames[5], s.wireNames[6]})
	s.digitRunes = append(s.digitRunes, []rune{s.wireNames[0], s.wireNames[1], s.wireNames[2], s.wireNames[3], s.wireNames[4], s.wireNames[6]})
}

func (s *segment) setSeg0Wire(one string, seven string) {
	//segment 0 is the remaining wire from 7 when the two wires from 1 are removed
	for _, r := range seven {
		if strings.ContainsRune(one, r) {
			continue
		} else {
			s.wireNames[0] = r
			return
		}
	}
}

func (s *segment) setSeg4Wire(four string, maybeZeroSixNine []string) {
	//segment 4 is the remaining wire when all four wires from 4 are removed and 0
	unmatched := ""
	for _, digit := range maybeZeroSixNine {
		for _, r := range digit {
			// remove four from digit
			if strings.ContainsRune(four, r) {
				continue
			}
			// remove s.wireNames[0] from digit
			if r == s.wireNames[0] {
				continue
			}
			unmatched = fmt.Sprintf("%s%c", unmatched, r)
		}
		if len(unmatched) == 1 {
			s.wireNames[4] = rune(unmatched[0])
			return
		}
		unmatched = ""
	}
}

func (s *segment) setSeg2Wire(one string, maybeTwoThreeFive []string) {
	// remove 0, 4, and wires in one from maybe235
	// if one remaining, it's 2
	unmatched := ""
	for _, digit := range maybeTwoThreeFive {
		for _, r := range digit {
			// remove one from digit
			if strings.ContainsRune(one, r) {
				continue
			}
			// remove s.wireNames[0] and [4] from digit
			if r == s.wireNames[0] || r == s.wireNames[4] {
				continue
			}
			unmatched = fmt.Sprintf("%s%c", unmatched, r)
		}
		if len(unmatched) == 1 {
			s.wireNames[2] = rune(unmatched[0])
			return
		}
		unmatched = ""
	}
}

func (s *segment) setSeg5Wire(four string, maybeTwoThreeFive []string) {
	// remove 0, 4 and wires in four from maybe235
	// if one remaining, it's 5
	unmatched := ""
	for _, digit := range maybeTwoThreeFive {
		for _, r := range digit {
			// remove four from digit
			if strings.ContainsRune(four, r) {
				continue
			}
			// remove s.wireNames[0] and [4] from digit
			if r == s.wireNames[0] || r == s.wireNames[4] {
				continue
			}
			unmatched = fmt.Sprintf("%s%c", unmatched, r)
		}
		if len(unmatched) == 1 {
			s.wireNames[5] = rune(unmatched[0])
			return
		}
		unmatched = ""
	}
}

func (s *segment) setSeg1Wire(one string, maybeZeroSixNine []string) {
	// take maybe069 away from one.  if one wire remaining, it's 1
	unmatched := ""
	for _, digit := range maybeZeroSixNine {
		for _, r := range one {
			// remove digit from one
			if strings.ContainsRune(digit, r) {
				continue
			}
			unmatched = fmt.Sprintf("%s%c", unmatched, r)
		}
		if len(unmatched) == 1 {
			s.wireNames[1] = rune(unmatched[0])
			return
		}
		unmatched = ""
	}
}

func (s *segment) setSeg3Wire(maybeTwoThreeFive []string) {
	// take 0,1,2,4,5 from maybe235, if one wire remaining it's 3
	unmatched := ""
	for _, digit := range maybeTwoThreeFive {
		for _, r := range digit {
			// remove s.wireNames[0] 1,2,4,5 from digit
			if r == s.wireNames[0] || r == s.wireNames[1] || r == s.wireNames[2] || r == s.wireNames[4] || r == s.wireNames[5] {
				continue
			}
			unmatched = fmt.Sprintf("%s%c", unmatched, r)
		}
		if len(unmatched) == 1 {
			s.wireNames[3] = rune(unmatched[0])
			return
		}
		unmatched = ""
	}
}

func (s *segment) setSeg6Wire(four string) {
	// remove 1,2,3 from four.  what remains is 6
	for _, r := range four {
		// remove s.wireNames[1] 2,3 from four
		if r == s.wireNames[1] || r == s.wireNames[2] || r == s.wireNames[3] {
			continue
		}
		s.wireNames[6] = r
	}
}

func (s *segment) digitForString(digit string) int {
	for i, drs := range s.digitRunes {
		if len(digit) != len(drs) {
			continue
		}
		foundDigit := true
		for _, r := range digit {
			foundRune := false
			for _, dr := range drs {
				if r == dr {
					foundRune = true
					break
				}
			}
			if !foundRune {
				foundDigit = false
				break
			}
		}
		if foundDigit {
			log.Println("digit", digit, "is number", i)
			return i
		}
	}
	return -1
}

func (s *segment) getOutput(code string) int {
	segs := strings.Fields(code)
	value := 0
	for i := 0; i < len(segs); i++ {
		value += (int(math.Pow10(i)) * s.digitForString(segs[len(segs)-i-1]))
	}
	log.Println("code", code, "is", value)
	return value
}

func main() {
	day8_1()
	day8_2()
}

func day8_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	digitCount := 0
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), "|")
		litSegments := strings.Fields(lineParts[1])
		for _, seg := range litSegments {
			if len(seg) == 2 || len(seg) == 3 || len(seg) == 4 || len(seg) == 7 {
				digitCount++
			}
		}
	}

	log.Println("There were", digitCount, "occurrences of digits 1,4,7, or 8")
}

func day8_2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), "|")
		s := segment{}
		s.setWireNames(lineParts[0])
		sum += s.getOutput(lineParts[1])
	}

	log.Println("total of all output values:", sum)
}
