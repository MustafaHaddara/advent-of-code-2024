package main

import (
	"strconv"
	"strings"
)

type Day07 struct {
}

func (d Day07) TestInput() string {
	return `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
}

func (d Day07) parseTargetAndOperands(line string) (int, []int) {
	chunks := strings.Split(line, ": ")
	target, _ := strconv.Atoi(chunks[0])
	rest := strings.Split(chunks[1], " ")
	operands := make([]int, len(rest))
	for i, s := range rest {
		operands[i], _ = strconv.Atoi(s)
	}
	return target, operands
}

func (d Day07) isValid(target int, operands []int) bool {
	if len(operands) == 1 {
		return operands[0] == target
	}

	last := operands[len(operands)-1]
	first := operands[:len(operands)-1]

	// if target % first != 0 that means it can't be a factor
	return (target % last == 0 && d.isValid(target / last, first)) || d.isValid(target - last, first)
}

func (d Day07) isValidConcat(target int, operands []int) bool {
	if len(operands) == 1 {
		return operands[0] == target
	}

	last := operands[len(operands)-1]
	first := operands[:len(operands)-1]

	// if target % first != 0 that means it can't be a factor
	return (target % last == 0 && d.isValidConcat(target / last, first)) || 
		d.isValidConcat(target - last, first) || 
		d.isValidSuffix(target, last, first)
}

func (d Day07) isValidSuffix(target int, last int, first []int) bool {
	remaining := d.removeSuffix(target, last)
	if remaining == -1 {
		return false
	}
	return d.isValidConcat(remaining, first)
}

func (d Day07) removeSuffix(target int, suffix int) int {
	targetS := strconv.Itoa(target)
	suffixS := strconv.Itoa(suffix)

	res := strings.TrimSuffix(targetS, suffixS)
	if res == targetS {
		// suffix doesn't match
		return -1
	}
	resI,_ := strconv.Atoi(res)
	return resI
}

func (d Day07) SolveA(lines []string) string {
	total := 0

	for _, line := range lines {
		target, operands := d.parseTargetAndOperands(line)

		if d.isValid(target, operands) {
			total += target
		}
	}

	return strconv.Itoa(total)
}

func (d Day07) SolveB(lines []string) string {
	total := 0

	for _, line := range lines {
		target, operands := d.parseTargetAndOperands(line)

		if d.isValidConcat(target, operands) {
			total += target
		}
	}

	return strconv.Itoa(total)
}
