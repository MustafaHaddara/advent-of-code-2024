package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solver interface {
	SolveA([]string) string
	SolveB([]string) string
	TestInput() string
}

var solvers = []Solver{}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 3 {
		fmt.Println("Usage: <bin> <day> <part> test?")
		os.Exit(1)
	}
	test := len(os.Args) > 3 && strings.ToLower(os.Args[3]) == "test"

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day > len(solvers) {
		fmt.Println("invalid day num given")
		os.Exit(1)
	}
	solver := solvers[day-1]

	var rows []string
	if test {
		rows = strings.Split(solver.TestInput(), "\n")
	} else {
		cwd, _ := os.Getwd()
		fi, err := os.ReadFile(fmt.Sprintf("%s/inputs/day%02d.txt", cwd, day))
		if err != nil {
			panic(err)
		}
		contents := string(fi[:])
		rows = strings.Split(contents, "\n")
	}

	a := strings.ToLower(os.Args[2]) == "a"

	if a {
		fmt.Println(solvers[0].SolveA(rows))
	} else {
		fmt.Println(solvers[0].SolveB(rows))
	}
}
