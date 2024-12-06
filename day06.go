package main

import (
	"slices"
	"strconv"
)

type Day06 struct {
}

func (d Day06) TestInput() string {
	return `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
}

type position struct {
	x int
	y int
}

func (d Day06) findInitialPosition(lines []string) position {
	for y, line := range lines {
		for x, c := range line {
			if c == rune('^') {
				return position{x, y}
			}
		}
	}
	return position{-1, -1}
}

func (d Day06) isOutOfBounds(guard position, lines []string) bool {
	return guard.x < 0 || guard.x >= len(lines[0]) || guard.y < 0 || guard.y >= len(lines)
}

func (d Day06) SolveA(lines []string) string {
	positions := make(map[position]bool)

	guard := d.findInitialPosition(lines)
	positions[guard] = true

	UP := position{0, -1}
	DOWN := position{0, 1}
	LEFT := position{-1, 0}
	RIGHT := position{1, 0}

	direction := UP
	for {
		next := position{guard.x + direction.x, guard.y + direction.y}
		if d.isOutOfBounds(next, lines) {
			break
		}
		if lines[next.y][next.x] != '#' {
			guard = next
			positions[guard] = true

		} else {
			if direction == UP {
				direction = RIGHT
			} else if direction == RIGHT {
				direction = DOWN
			} else if direction == DOWN {
				direction = LEFT
			} else if direction == LEFT {
				direction = UP
			}
		}
	}

	return strconv.Itoa(len(positions))
}

func (d Day06) makesLoop(guard position, lines []string, newObstacle position) bool {
	UP := position{0, -1}
	DOWN := position{0, 1}
	LEFT := position{-1, 0}
	RIGHT := position{1, 0}
	direction := UP

	seen := make(map[position][]position)

	for {
		if slices.Contains(seen[guard], direction) {
			return true
		}
		s := seen[guard]
		s = append(s, direction)
		seen[guard] = s

		next := position{guard.x + direction.x, guard.y + direction.y}
		if d.isOutOfBounds(next, lines) {
			return false
		}

		if lines[next.y][next.x] != '#' && !(next.x == newObstacle.x && next.y == newObstacle.y) {
			guard = next
		} else {
			if direction == UP {
				direction = RIGHT
			} else if direction == RIGHT {
				direction = DOWN
			} else if direction == DOWN {
				direction = LEFT
			} else if direction == LEFT {
				direction = UP
			}
		}
	}
}

func (d Day06) SolveB(lines []string) string {
	total := 0

	guard := d.findInitialPosition(lines)

	for y, line := range lines {
		for x, _ := range line {
			newObstacle := position{x, y}
			if guard == newObstacle {
				continue
			}
			if d.makesLoop(guard, lines, newObstacle) {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}
