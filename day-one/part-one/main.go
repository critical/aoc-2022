package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	calories := []int{}
	tempCalSum := 0

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			calories = append(calories, tempCalSum)
			tempCalSum = 0
			continue
		}

		tempCalSum += num
	}

	sort.Ints(calories)

	log.Println(calories[len(calories)-1])
}
