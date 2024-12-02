package main

import (
	"strconv"
	"strings"
)

type Day02 struct {
}

func (d Day02) TestInput() string {
	return `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
}

func (d Day02) SolveA(lines []string) string {
	total := 0
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		if d.isReportSafe(chunks, -1) {
			total += 1
		}
	}
	return strconv.Itoa(total)
}

func (d Day02) SolveB(lines []string) string {
	total := 0
	for _, line := range lines {
		chunks := strings.Split(line, " ")
		for i := -1; i < len(chunks); i++ {
			if d.isReportSafe(chunks, i) {
				total += 1
				break
			}
		}
	}
	return strconv.Itoa(total)
}

func (d Day02) isReportSafe(chunks []string, idxToSkip int) bool {
	last := -1
	lastdiff := 0

	for idx, numStr := range chunks {
		if idx == idxToSkip {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		if last == -1 {
			last = num
			continue
		}
		diff := num - last
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if lastdiff != 0 {
			if (diff < 0 && lastdiff > 0) || (diff > 0 && lastdiff < 0) {
				return false
			}
		}
		last = num
		lastdiff = diff
	}
	return true
}
