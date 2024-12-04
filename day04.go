package main

import (
	"strconv"
)

type Day04 struct {
}

func (d Day04) TestInput() string {
	return `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
}

func (d Day04) SolveA(lines []string) string {
	total := 0

	// horizontal
	for _, line := range lines {
		for x := 0; x < len(line)-3; x++ {
			chunk := line[x : x+4]
			if chunk == "XMAS" || chunk == "SAMX" {
				total += 1
			}
		}
	}

	// vertical
	for y := 0; y < len(lines)-3; y++ {
		for x := 0; x < len(lines[y]); x++ {
			chunk := string(lines[y][x]) + string(lines[y+1][x]) + string(lines[y+2][x]) + string(lines[y+3][x])
			if chunk == "XMAS" || chunk == "SAMX" {
				total += 1
			}
		}
	}

	// diagonal top right -> bottom left
	for y := 0; y < len(lines)-3; y++ {
		for x := 0; x < len(lines[y])-3; x++ {
			chunk := string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2]) + string(lines[y+3][x+3])
			if chunk == "XMAS" || chunk == "SAMX" {
				total += 1
			}
		}
	}

	// diagonal top left -> bottom right
	for y := 3; y < len(lines); y++ {
		for x := 0; x < len(lines[y])-3; x++ {
			chunk := string(lines[y][x]) + string(lines[y-1][x+1]) + string(lines[y-2][x+2]) + string(lines[y-3][x+3])
			if chunk == "XMAS" || chunk == "SAMX" {
				total += 1
			}
		}
	}

	return strconv.Itoa(total)
}

func (d Day04) SolveB(lines []string) string {
	total := 0

	// diagonal top left -> bottom right
	for y := 0; y < len(lines)-2; y++ {
		for x := 0; x < len(lines[y])-2; x++ {
			// center is A
			if lines[y+1][x+1] == 'A' {
				topL := lines[y][x]
				topR := lines[y][x+2]
				bottomL := lines[y+2][x]
				bottomR := lines[y+2][x+2]

				// top 2 are M
				if topL == 'M' && topR == 'M' && bottomL == 'S' && bottomR == 'S' {
					total += 1
					continue
				}

				// top 2 are S
				if topL == 'S' && topR == 'S' && bottomL == 'M' && bottomR == 'M' {
					total += 1
					continue
				}

				// left 2 are M
				if topL == 'M' && topR == 'S' && bottomL == 'M' && bottomR == 'S' {
					total += 1
					continue
				}

				// left 2 are S
				if topL == 'S' && topR == 'M' && bottomL == 'S' && bottomR == 'M' {
					total += 1
					continue
				}
			}
		}
	}

	return strconv.Itoa(total)
}
