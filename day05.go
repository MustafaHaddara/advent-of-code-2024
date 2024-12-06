package main

import (
	"strconv"
	"strings"
)

type Day05 struct {
}

func (d Day05) TestInput() string {
	return `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
}

func (d Day05) parseDeps(lines []string) map[int][]int {
	var deps map[int][]int = make(map[int][]int)

	for i := 0; lines[i] != ""; i++ {
		chunks := strings.Split(lines[i], "|")
		val1, _ := strconv.Atoi(chunks[0])
		val2, _ := strconv.Atoi(chunks[1])

		if deps[val1] == nil {
			deps[val1] = []int{}
		}
		deps[val1] = append(deps[val1], val2)
	}

	return deps
}

func (d Day05) isValid(vals []int, deps map[int][]int) bool {
	seen := make(map[int]bool)
	for _, val := range vals {
		seen[val] = true

		d := deps[val]
		for _, expected := range d {
			if seen[expected] {
				// if we've already seen a value that must be after us
				// this array is invalid
				return false
			}
		}
	}
	return true
}

func (d Day05) SolveA(lines []string) string {
	total := 0

	deps := d.parseDeps(lines)
	i := 0
	for ; lines[i] != ""; i++ {
	}
	i++

	for ; i < len(lines); i++ {
		spec := lines[i]
		specs := strings.Split(spec, ",")
		ints := make([]int, len(specs))

		for i, c := range specs {
			val, _ := strconv.Atoi(c)
			ints[i] = val
		}

		if d.isValid(ints, deps) {
			// find the middle
			mid := (len(ints) / 2)
			total += ints[mid]
		}

	}

	return strconv.Itoa(total)
}

func (d Day05) findValidLast(vals map[int]bool, deps map[int][]int, seen map[int]bool) int {
outer:
	for val, _ := range vals {
		if seen[val] {
			// we've used this one already
			continue
		}
		dep := deps[val]
		for _, d := range dep {
			if !seen[d] && vals[d] {
				// can't be val
				continue outer
			}
		}

		// we get here, there's nothing left we need to be before
		// we can put this one at the end
		return val
	}
	return -1
}

func (d Day05) fix(vals []int, deps map[int][]int) []int {
	// our array, in set form
	set := make(map[int]bool)
	for _, val := range vals {
		set[val] = true
	}

	// get the one that can be last
	seen := make(map[int]bool)
	res := make([]int, len(vals))

	for i := len(vals) - 1; i >= 0; i-- {
		last := d.findValidLast(set, deps, seen)
		res[i] = last
		seen[last] = true
	}

	return res
}

func (d Day05) SolveB(lines []string) string {
	total := 0

	deps := d.parseDeps(lines)
	i := 0
	for ; lines[i] != ""; i++ {
	}
	i++

	for ; i < len(lines); i++ {
		spec := lines[i]
		specs := strings.Split(spec, ",")
		ints := make([]int, len(specs))

		for i, c := range specs {
			val, _ := strconv.Atoi(c)
			ints[i] = val
		}

		if !d.isValid(ints, deps) {
			// fix
			fixed := d.fix(ints, deps)

			// find the middle
			mid := (len(fixed) / 2)
			total += fixed[mid]
		}

	}

	return strconv.Itoa(total)
}
