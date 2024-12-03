package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func readInput() ([]int, []int) {
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

	var left []int
	var right []int
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())

		l, _ := strconv.Atoi(nums[0])
		left = append(left, l)

		r, _ := strconv.Atoi(nums[1])
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func part1() {
	left, right := readInput()

	l := sort.IntSlice(left)
	r := sort.IntSlice(right)

	l.Sort()
	r.Sort()

	distances := make([]int, len(l))
	for idx, lVal := range l {
		rVal := r[idx]

		diff := lVal - rVal
		if diff < 0 {
			distances[idx] = -diff
		} else {
			distances[idx] = diff
		}
	}

	result := 0
	for _, d := range distances {
		result += d
	}

	fmt.Println(result)
}

func part2() {
	left, right := readInput()

	similarity := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				similarity += l
			}
		}
	}

	fmt.Println(similarity)
}
