package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var hits int
	x := 3
	for scanner.Scan() {
		row := scanner.Text()
		if string(row[x%31]) == "#" {
			hits++
		}

		x += 3
	}

	fmt.Println(hits)
}
