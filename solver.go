package main

import (
	"fmt"
)

const (
	Empty = iota
	Red
	Blue
)

func not(b byte) byte {
	if b == Red {
		return Blue
	} else if b == Blue {
		return Red
	} else {
		return Empty
	}
}

type Board [][]byte

func newPuzzle(puzzle []string) Board {
	size := len(puzzle)
	b := make([][]byte, size)
	for y := range puzzle {
		b[y] = make([]byte, size)
		for x := range puzzle[y] {
			if puzzle[y][x] == 'R' {
				b[y][x] = Red
			} else if puzzle[y][x] == 'B' {
				b[y][x] = Blue
			} else {
				b[y][x] = Empty
			}
		}
	}
	return b
}

func (b Board) solved() bool {
	for _, row := range b[:] {
		for _, square := range row {
			if square == Empty {
				return false
			}
		}
	}
	return true
}

func (b Board) fillAdj() {
	// horizontal
	for y := 0; y < len(b); y++ {
		for x := 1; x < len(b)-1; x++ {
			if b[y][x-1] == b[y][x] && b[y][x+1] == Empty {
				b[y][x+1] = not(b[y][x])
			} else if b[y][x-1] == b[y][x+1] && b[y][x] == Empty {
				b[y][x] = not(b[y][x-1])
			} else if b[y][x] == b[y][x+1] && b[y][x-1] == Empty {
				b[y][x-1] = not(b[y][x])
			}
		}
	}
	// vertical
	for y := 1; y < len(b)-1; y++ {
		for x := 0; x < len(b); x++ {
			if b[y-1][x] == b[y][x] && b[y+1][x] == Empty {
				b[y+1][x] = not(b[y][x])
			} else if b[y-1][x] == b[y+1][x] && b[y][x] == Empty {
				b[y][x] = not(b[y-1][x])
			} else if b[y][x] == b[y+1][x] && b[y-1][x] == Empty {
				b[y-1][x] = not(b[y][x])
			}
		}
	}
}

func (b Board) fillLine() {
	// horizontal
	for y := range b {
		var blue, red int
		var empty []int
		for x := range b {
			if b[y][x] == Red {
				red++
			} else if b[y][x] == Blue {
				blue++
			} else {
				empty = append(empty, x)
			}
		}
		if red == len(b)/2 {
			for _, i := range empty {
				b[y][i] = Blue
			}
		} else if blue == len(b)/2 {
			for _, i := range empty {
				b[y][i] = Red
			}
		}
	}
	// vertical
	for x := range b {
		var blue, red int
		var empty []int
		for y := range b {
			if b[y][x] == Red {
				red++
			} else if b[y][x] == Blue {
				blue++
			} else {
				empty = append(empty, y)
			}
		}
		if red == len(b)/2 {
			for _, i := range empty {
				b[i][x] = Blue
			}
		} else if blue == len(b)/2 {
			for _, i := range empty {
				b[i][x] = Red
			}
		}
	}
}

func (b Board) fillDups() {
	// horizontal
	for y1 := range b {
		for y2 := range b {
			same := 0
			for x := range b {
				if b[y1][x] == b[y2][x] && b[y2][x] != Empty {
					same++
				}
			}
			if same == len(b)-2 {
				for x := range b {
					if b[y1][x] == Empty {
						b[y1][x] = not(b[y2][x])
					}
				}
			}
		}
	}
	// vertical
	for x1 := range b {
		for x2 := range b {
			same := 0
			for y := range b {
				if b[y][x1] == b[y][x2] && b[y][x2] != Empty {
					same++
				}
			}
			if same == len(b)-2 {
				for y := range b {
					if b[y][x1] == Empty {
						b[y][x1] = not(b[y][x2])
					}
				}
			}
		}
	}
}

func (b Board) solve() {
	for !b.solved() {
		b.fillAdj()
		b.fillLine()
		b.fillDups()
	}
}

func (b Board) print() {
	for _, row := range b {
		for _, square := range row {
			if square == Red {
				fmt.Print("R ")
			} else if square == Blue {
				fmt.Print("B ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	puzzle := newPuzzle([]string{
		"__RR__R__R",
		"B___B__B__",
		"_____B___B",
		"__R__B____",
		"B___B__R_B",
		"___R_B__RR",
		"_RR_______",
		"____B_____",
		"_BB___R__R",
		"_BR___R___",
	})
	puzzle.solve()
	puzzle.print()
}
