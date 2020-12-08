package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var validTokens = map[string]struct{}{
	"byr": {},
	"iyr": {},
	"eyr": {},
	"hgt": {},
	"hcl": {},
	"ecl": {},
	"pid": {},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var validCount int
	buf := ""
	for scanner.Scan() {
		row := scanner.Text()
		if row != "" {
			buf = fmt.Sprintf("%s %s", buf, row)
			continue
		}

		allTokens := map[string]struct{}{}
		tokens := strings.Split(strings.Trim(buf, " "), " ")
		for _, token := range tokens {
			token := strings.Split(token, ":")[0]
			if token != "cid" {
				allTokens[token] = struct{}{}
			}
		}

		if reflect.DeepEqual(validTokens, allTokens) {
			validCount++
		}

		buf = row
	}

	fmt.Println(validCount)
}
