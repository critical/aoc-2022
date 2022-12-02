package main

import (
	"bufio"
	"log"
	"os"
)

// If you're reading this... I'm not proud of this code.
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	partOneOutcomes := map[string]int{
		"A X": 3 + 1,
		"A Y": 6 + 2,
		"A Z": 0 + 3,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 6 + 1,
		"C Y": 0 + 2,
		"C Z": 3 + 3,
	}

	partTwoOutcomes := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}

	partOneTotalScore := 0
	partTwoTotalScore := 0

	for scanner.Scan() {
		partOneTotalScore += partOneOutcomes[scanner.Text()]
		partTwoTotalScore += partTwoOutcomes[scanner.Text()]
	}

	log.Println("Part 1 score:", partOneTotalScore)
	log.Println("Part 2 score:", partTwoTotalScore)
}
