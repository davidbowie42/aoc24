package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const reg1 = `mul\([\d]{1,3},[\d]{1,3}\)`
const reg2 = `mul\([\d]{1,3},[\d]{1,3}\)|do\(\)|don\'t\(\)`

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(content)
	part2(content)
}

func handleMul(m string) int {
	nums := strings.Replace(m, "mul(", "", 1)
	nums = strings.Replace(nums, ")", "", -1)

	n := strings.Split(nums, ",")

	n1, _ := strconv.Atoi(n[0])
	n2, _ := strconv.Atoi(n[1])

	return n1 * n2
}

func part1(content []byte) {
	r := regexp.MustCompile(reg1)

	matches := r.FindAllString(string(content), -1)

	sum := 0
	for _, m := range matches {
		sum += handleMul(m)
	}

	fmt.Println(sum)
}

func part2(content []byte) {
	r := regexp.MustCompile(reg2)

	matches := r.FindAllString(string(content), -1)

	sum := 0
	activated := true
	for _, m := range matches {
		if strings.HasPrefix(m, "don") {
			activated = false
		} else if strings.HasPrefix(m, "do") {
			activated = true
		} else if activated {
			sum += handleMul(m)
		}
	}

	fmt.Println(sum)
}
