package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidityChecker(parseable string, isValid func(min, max int64, password, letter string) bool) bool {
	colonSplit := strings.Split(parseable, ":")
	spaceSplit := strings.Split(colonSplit[0], " ")
	hyphenSplit := strings.Split(spaceSplit[0], "-")
	min, _ := strconv.ParseInt(hyphenSplit[0], 10, 0)
	max, _ := strconv.ParseInt(hyphenSplit[1], 10, 0)

	return isValid(min, max, strings.Trim(colonSplit[1], " "), spaceSplit[1])
}

func IsValid(firstPosition, secondPosition int64, password, letter string) bool {
	var count int
	if string(password[firstPosition-1]) == letter {
		count++
	}
	if string(password[secondPosition-1]) == letter {
		count++
	}

	return count == 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var validCount int
	for scanner.Scan() {
		entry := scanner.Text()

		if ValidityChecker(entry, IsValid) {
			validCount++
		}
	}

	fmt.Println(validCount)
}
