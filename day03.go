package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Day03 struct {
}

func (d Day03) TestInput() string {
	return `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
}



func (d Day03) SolveA(lines []string) string {
	var r = regexp.MustCompile(`mul\((\d*),(\d*)\)`)

	// get one continuous line
	mem := strings.Join(lines, "\n")
	matches := r.FindAllStringSubmatch(mem, -1)
	
	total := 0
	for _, match := range matches {
		prod := 1
		for i := 1; i < len(match); i++ {
			val,_ := strconv.Atoi(match[i])
			prod *= val
		}
		total += prod
	}
	return strconv.Itoa(total)
}

func (d Day03) SolveB(lines []string) string {
	var r = regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	var enabled = regexp.MustCompile(`do\(\)`)
	var disabled = regexp.MustCompile(`don't\(\)`)

	// get one continuous line
	mem := strings.Join(lines, "\n")
	matchIdxs := r.FindAllStringIndex(mem, -1)
	enabledIdxs := enabled.FindAllStringIndex(mem, -1)
	disabledIdxs := disabled.FindAllStringIndex(mem, -1)

	total := 0
	isEnabled := true
	matchIdx, enabledIdx, disabledIdx := 0, 0, 0

	for (matchIdx < len(matchIdxs)) {
		var en []int
		if enabledIdx >= len(enabledIdxs) {
			en = []int{-1}
		} else {
			en = enabledIdxs[enabledIdx]
		}

		var dis []int
		if disabledIdx >= len(disabledIdxs) {
			dis = []int{-1}
		} else {
			dis = disabledIdxs[disabledIdx]
		}

		m := matchIdxs[matchIdx]

		if en[0] != -1 && (dis[0] == -1 || en[0] < dis[0]) && en[0] < m[0] {
			isEnabled = true
			enabledIdx++
			continue
		}
		if dis[0] != -1 && (en[0] == -1 || dis[0] < en[0]) && dis[0] < m[0] {
			isEnabled = false
			disabledIdx++
			continue
		}

		matchIdx++

		if !isEnabled {
			continue
		}
		
		// got a match
		match := r.FindStringSubmatch(mem[m[0]:m[1]])
		prod := 1
		for i := 1; i < len(match); i++ {
			val,_ := strconv.Atoi(match[i])
			prod *= val
		}
		total += prod
	}
	
	return strconv.Itoa(total)
}
