package main

import (
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
}

func (d Day01) TestInput() string {
	return `3   4
4   3
2   5
1   3
3   9
3   3`
}

func (d Day01) SolveA(lines []string) string {
	list1 := []int{}
	list2 := []int{}
	// each row is 2 numbers
	for _, line := range lines {
		res := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(res[0])
		list1 = append(list1, num1)

		num2, _ := strconv.Atoi(res[1])
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	total := 0
	for i, num1 := range list1 {
		num2 := list2[i]
		dist := absInt(num1 - num2)
		total += dist
	}

	return strconv.Itoa(total)
}

func (d Day01) SolveB(lines []string) string {
	list1 := []int{}
	list2 := map[int]int{}
	// each row is 2 numbers
	for _, line := range lines {
		res := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(res[0])
		list1 = append(list1, num1)

		num2, _ := strconv.Atoi(res[1])
		count := list2[num2]
		list2[num2] = count + 1
	}

	total := 0
	for _, num := range list1 {
		total += num * list2[num]
	}
	return strconv.Itoa(total)
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
