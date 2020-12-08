package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	requiredAmounts := make(map[int64]interface{})

	for scanner.Scan() {
		// Throw away error the file is fine
		value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		_, ok := requiredAmounts[value]

		if !ok {
			requiredAmounts[2020-value] = struct{}{}
			continue
		}

		fmt.Println(value * (2020 - value))
		break
	}
}
