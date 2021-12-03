package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func main() {
	day3_1()
	day3_2()
}

func day3_1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	zeroCount := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	oneCount := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for scanner.Scan() {
		for i, v := range scanner.Text() {
			switch v {
			case '0':
				zeroCount[i]++
			case '1':
				oneCount[i]++
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0
	for i := 0; i < 12; i++ {
		// string has index 0 at the left, but that should be the power of
		// 11 when converting from bin to int
		pow := 11 - i
		if oneCount[i] > zeroCount[i] {
			gammaRate += int(math.Pow(2, float64(pow)))
		} else {
			epsilonRate += int(math.Pow(2, float64(pow)))
		}
	}

	log.Println("Power consumption is", gammaRate*epsilonRate)
}

func day3_2() {
	numbers := inputToArray()
	for {
		for i := 0; i < 12; i++ {
			mc := mostCommon(numbers, i)
			numbers = filterNumbers(numbers, i, mc)
			if len(numbers) == 1 {
				break
			}
		}
		if len(numbers) == 1 {
			break
		}
	}
	oxygenGeneratorRating := convertToInt(numbers[0])

	numbers = inputToArray()
	for {
		for i := 0; i < 12; i++ {
			mc := leastCommon(numbers, i)
			numbers = filterNumbers(numbers, i, mc)
			if len(numbers) == 1 {
				break
			}
		}
		if len(numbers) == 1 {
			break
		}
	}
	co2ScrubberRating := convertToInt(numbers[0])

	log.Println("Life support rating is", oxygenGeneratorRating*co2ScrubberRating)
}

func inputToArray() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	numbers := make([]string, 0, 1000) // length of input.txt
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	return numbers
}

func mostCommon(numbers []string, position int) int {
	zeroCount, oneCount := bitCounts(numbers, position)
	if oneCount >= zeroCount {
		return 1
	} else {
		return 0
	}
}

func leastCommon(numbers []string, position int) int {
	zeroCount, oneCount := bitCounts(numbers, position)
	if oneCount < zeroCount {
		return 1
	} else {
		return 0
	}
}

func bitCounts(numbers []string, position int) (int, int) {
	zeroCount := 0
	oneCount := 0
	for _, v := range numbers {
		switch v[position] {
		case '0':
			zeroCount++
		case '1':
			oneCount++
		}
	}
	return zeroCount, oneCount
}

func filterNumbers(numbers []string, position int, keepValue int) []string {
	keptNumbers := make([]string, 0)
	for _, number := range numbers {
		// convert run to int by casting and subtracting the int value of '0'
		// works for ascii characters 0 to 9
		if int(number[position]-'0') == keepValue {
			keptNumbers = append(keptNumbers, number)
		}
	}
	return keptNumbers
}

func convertToInt(number string) int {
	v := 0
	for i, r := range number {
		// string has index 0 at the left, but that should be the power of
		// 11 when converting from bin to int
		pow := 11 - i
		n := int(r - '0')
		v += int(math.Pow(2, float64(pow))) * n
	}
	return v
}
