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
	part1()
	part2()
}

const minDiff = 1
const maxDiff = 3

func part1() {
	reports := readInput()

	nSafe := 0
	for _, report := range reports {
		diffs := make([]int, len(report)-1)
		prevVal := report[0]

		for i := 1; i < len(report); i++ {
			diffs[i-1] = prevVal - report[i]

			prevVal = report[i]
		}

		isPositive := diffs[0] > 0
		for i, diff := range diffs {
			if isPositive {
				if diff < minDiff || diff > maxDiff {
					fmt.Println(diffs, "not safe", i, isPositive)
					break
				}

				if i == len(diffs)-1 {
					nSafe++
					fmt.Println(diffs, "safe")
				}
			} else {
				if diff > -minDiff || diff < -maxDiff {
					fmt.Println(diffs, "not safe", i, isPositive)
					break
				}

				if i == len(diffs)-1 {
					nSafe++
					fmt.Println(diffs, "safe")
				}
			}
		}

	}
	fmt.Println(nSafe)
}

func isSafe(diffs []int) (bool, int) {
	isPositive := diffs[0] > 0
	for i, diff := range diffs {
		if isPositive {
			if diff < minDiff || diff > maxDiff {
				return false, i
			}
		} else {
			if diff > -minDiff || diff < -maxDiff {
				return false, i
			}
		}
	}

	return true, -1
}

func calcDiffs(report []int) []int {
	diffs := make([]int, len(report)-1)
	prevVal := report[0]

	for i := 1; i < len(report); i++ {
		diffs[i-1] = prevVal - report[i]

		prevVal = report[i]
	}

	return diffs
}

func dampen(report []int, wrongIdx int) bool {
	newReport := make([]int, 0)
	for i, val := range report {
		if i == wrongIdx {
			continue
		}
		newReport = append(newReport, val)
	}

	newDiffs := calcDiffs(newReport)
	safe, _ := isSafe(newDiffs)

	return safe
}

func part2() {
	reports := readInput()

	nSafe := 0
	for _, report := range reports {
		diffs := calcDiffs(report)
		safe, wrongIdx := isSafe(diffs)

		if safe {
			nSafe++
			continue
		} else {
			fmt.Println(report, diffs, "not safe", wrongIdx)
		}

		safe = dampen(report, wrongIdx) || dampen(report, wrongIdx-1) || dampen(report, wrongIdx+1)
		if safe {
			nSafe++
			continue
		}
	}

	fmt.Println(nSafe)
}

func readInput() [][]int {
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

	var out [][]int
	for scanner.Scan() {
		numsStr := strings.Fields(scanner.Text())

		nums := make([]int, len(numsStr))

		for i, s := range numsStr {
			nums[i], _ = strconv.Atoi(s)
		}

		out = append(out, nums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}
