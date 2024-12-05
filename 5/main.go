package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readInput()

	incorrectUpdates := part1(rules, updates)
	part2(rules, incorrectUpdates)
}

func part1(rules map[int][]int, updates [][]int) [][]int {
	incorrectUpdates := make([][]int, 0)
	middlePageNum := 0
	for _, update := range updates {
		correct := isCorrect(rules, update)

		if correct {
			n := update[len(update)/2]
			middlePageNum += n
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	fmt.Println(middlePageNum)
	return incorrectUpdates
}

func part2(rules map[int][]int, incorrectUpdates [][]int) {
	incorrectMiddlePageNum := 0
	for _, update := range incorrectUpdates {
		corrected := make([]int, 0, len(update))

		for _, x := range update {
			ys, _ := rules[x]

			if len(corrected) == 0 {
				corrected = []int{x}
				continue
			}

			moved := false
			for j, c := range corrected {
				if !checkRules(c, ys) {
					corrected = moveRight(j, corrected)
					corrected[j] = x
					moved = true
					break
				}
			}

			if !moved {
				corrected = append(corrected, x)
			}
		}

		n := corrected[len(corrected)/2]
		incorrectMiddlePageNum += n
	}

	fmt.Println(incorrectMiddlePageNum)
}

func isCorrect(rules map[int][]int, update []int) bool {
	for i, x := range update {
		ys, _ := rules[x]

		for j := 0; j < len(ys); j++ {
			for k := 0; k < i; k++ {
				if ys[j] == update[k] {
					return false
				}
			}
		}
	}

	return true
}

func moveRight(start int, slice []int) []int {
	newSlice := make([]int, len(slice)+1)

	for i := 0; i < start; i++ {
		newSlice[i] = slice[i]
	}

	for i := start; i < len(slice); i++ {
		newSlice[i+1] = slice[i]
	}

	return newSlice
}

func checkRules(current int, ys []int) bool {
	for _, y := range ys {
		if current == y {
			return false
		}
	}

	return true
}

func readInput() (map[int][]int, [][]int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	rules := make(map[int][]int)
	var updates [][]int
	readingRules := true
	for scanner.Scan() {
		if readingRules {

			line := strings.Split(scanner.Text(), "|")

			if len(line) != 2 {
				readingRules = false
				continue
			}

			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])

			rules[x] = append(rules[x], y)
		} else {
			line := strings.Split(scanner.Text(), ",")
			update := make([]int, len(line))

			for i, s := range line {
				update[i], _ = strconv.Atoi(s)
			}

			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, updates
}
