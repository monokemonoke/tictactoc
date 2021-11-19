package main

import (
	"fmt"
	"strings"
)

func main() {
	board := InitBoard()
	ShowBoard(board)
}

func InitBoard() [][]string {
	board := make([][]string, 3)
	for i := 0; i < 3; i++ {
		board[i] = []string{"  ", "  ", "  "}
	}
	return board
}

func ShowBoard(board [][]string) {
	fmt.Println("  |  |  |  ")

	for i, row := range board {
		fmt.Println("--+--+--+--")
		line := strings.Join(row, "|")
		line = fmt.Sprintf("%d |", i) + line
		fmt.Println(line)
	}
}
