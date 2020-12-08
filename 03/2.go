package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var hits, hits2, hits3, hits4, hits5 int
	rowNum := 1
	for scanner.Scan() {
		row := scanner.Text()
		if isHit(1, 1, rowNum, row) {
			hits++
		}
		if isHit(3, 1, rowNum, row) {
			hits2++
		}
		if isHit(5, 1, rowNum, row) {
			hits3++
		}
		if isHit(7, 1, rowNum, row) {
			hits4++
		}
		if isHit(1, 2, rowNum, row) {
			hits5++
		}
		rowNum++
	}
	fmt.Println(hits, hits2, hits3, hits4, hits5)
	fmt.Println(hits * hits2 * hits3 * hits4 * hits5)
}

func isHit(right, down, rowNum int, row string) bool {
	return rowNum%down == 0 && string(row[(right*rowNum/down)%31]) == "#"
}
