package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	rows := readInput()
	part1(rows)
	part2(rows)
}

func part1(rows [][]byte) {
	xmasCounter := 0

	for i, row := range rows {
		for j, _ := range row {
			if checkDirection(i, j, 1, 0, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, 0, 1, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, 1, 1, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, 0, -1, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, -1, 0, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, -1, -1, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, 1, -1, rows) {
				xmasCounter++
			}
			if checkDirection(i, j, -1, 1, rows) {
				xmasCounter++
			}
		}
	}

	fmt.Println(xmasCounter)
}

func part2(rows [][]byte) {
	xmasCounter := 0
	for i, row := range rows {
		for j, _ := range row {
			if checkCross(i, j, rows) {
				xmasCounter++
			}
		}
	}

	fmt.Println(xmasCounter)
}

func checkCross(i int, j int, rows [][]byte) bool {
	if rows[i][j] != 'A' {
		return false
	}

	if !(indicesSafe(i-1, j-1, rows) && indicesSafe(i+1, j+1, rows) && indicesSafe(i-1, j+1, rows) && indicesSafe(i+1, j-1, rows)) {
		return false
	}

	if rows[i-1][j-1] == 'M' {
		if rows[i+1][j+1] != 'S' {
			return false
		}
	} else if rows[i-1][j-1] == 'S' {
		if rows[i+1][j+1] != 'M' {
			return false
		}
	} else {
		return false
	}

	if rows[i+1][j-1] == 'M' {
		if rows[i-1][j+1] != 'S' {
			return false
		}
	} else if rows[i+1][j-1] == 'S' {
		if rows[i-1][j+1] != 'M' {
			return false
		}
	} else {
		return false
	}

	return true
}

func checkDirection(i int, j int, iInc int, jInc int, rows [][]byte) bool {
	if rows[i][j] != 'X' {
		return false
	}

	i = i + iInc
	j = j + jInc

	if !(indicesSafe(i, j, rows) && rows[i][j] == 'M') {
		return false
	}

	i = i + iInc
	j = j + jInc

	if !(indicesSafe(i, j, rows) && rows[i][j] == 'A') {
		return false
	}

	i = i + iInc
	j = j + jInc

	if !(indicesSafe(i, j, rows) && rows[i][j] == 'S') {
		return false
	}

	return true
}

func indicesSafe(i int, j int, rows [][]byte) bool {
	if i < 0 || i >= len(rows) {
		return false
	}

	if j < 0 || j >= len(rows[i]) {
		return false
	}

	return true
}

func readInput() [][]byte {
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

	var out [][]byte
	for scanner.Scan() {
		out = append(out, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}
