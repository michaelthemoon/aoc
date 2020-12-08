package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func intValid(s string, min, max int) bool {
	if s == "" {
		return false
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if i < min {
		return false
	}

	if i > max {
		return false
	}

	return true
}

func ByrValid(s string) bool {
	return intValid(s, 1920, 2002)
}

func IyrValid(s string) bool {
	return intValid(s, 2010, 2020)
}

func EyrValid(s string) bool {
	return intValid(s, 2020, 2030)
}

func HgtValid(s string) bool {
	if s == "" {
		return false
	}

	if strings.Contains(s, "cm") {
		result := intValid(strings.Trim(s, "cm"), 150, 193)
		return result
	}

	return intValid(strings.Trim(s, "in"), 59, 76)
}

func HclValid(s string) bool {
	if s[0] != '#' {
		return false
	}

	if len(s) != 7 {
		return false
	}

	r, _ := regexp.MatchString("[a-f0-9]{6}", s)

	return r

}

var validEcl = map[string]struct{}{
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {},
}

func EclValid(s string) bool {
	_, ok := validEcl[s]

	return ok
}

func PidValid(s string) bool {
	if len(s) != 9 {
		return false
	}

	_, err := strconv.Atoi(string(s))

	return err == nil
}

var tokenFuncs = map[string]func(string) bool{
	"byr": ByrValid,
	"iyr": IyrValid,
	"eyr": EyrValid,
	"hgt": HgtValid,
	"hcl": HclValid,
	"ecl": EclValid,
	"pid": PidValid,
}

var validTokens = []string{
	"ecl",
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"pid",
}

var validCounts = map[string]int{
	"byr": 0,
	"iyr": 0,
	"eyr": 0,
	"hgt": 0,
	"hcl": 0,
	"ecl": 0,
	"pid": 0,
}

func validPassport(pp map[string]string) bool {
	if len(pp) < 7 {
		return false
	}

	for _, tk := range validTokens {
		s, ok := pp[tk]
		if !ok {
			return false
		}
		if !tokenFuncs[tk](s) {
			return false
		}

		validCounts[tk]++
	}

	return true
}

func main() {
	var rowCount int

	scanner := bufio.NewScanner(os.Stdin)

	var validCount int
	buf := ""
	for scanner.Scan() {
		rowCount++
		row := scanner.Text()
		if row != "" {
			buf = fmt.Sprintf("%s%s ", buf, row)
			continue
		}

		tokens := strings.Split(strings.Trim(buf, " "), " ")
		var pp = make(map[string]string)
		for _, token := range tokens {
			split := strings.Split(token, ":")
			pp[split[0]] = split[1]
		}

		if validPassport(pp) {
			validCount++
		}

		buf = ""
	}

	fmt.Println(validCount)
	fmt.Println(validCounts)
}
