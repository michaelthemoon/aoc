package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []int64

	for scanner.Scan() {
		value, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		inputs = append(inputs, value)
	}

	for i := 0; i < len(inputs)-2; i++ {
		for j := 1; j < len(inputs)-1; j++ {
			for k := 2; k < len(inputs); k++ {
				if inputs[i]+inputs[j]+inputs[k] != 2020 {
					continue
				}

				fmt.Println(inputs[i] * inputs[j] * inputs[k])
				goto loopbreak
			}
		}
	}
loopbreak:
}
