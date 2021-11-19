package main

import (
	"fmt"
	"strings"
)

const FirstPlayer = 1
const SecondPlayer = 2

func main() {
	board := InitBoard()
	ShowBoard(board)
	player := FirstPlayer

	for {
		fmt.Printf("Player is %d now\n", player)

		x, y := InputMove(board)
		fmt.Printf("Valid move is (%d, %d)\n", x, y)

		if player == FirstPlayer {
			player = SecondPlayer
		} else {
			player = FirstPlayer
		}
	}
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

func isValidMove(x, y int) bool {
	if x < 1 || 3 < x {
		return false
	}
	if y < 1 || 3 < y {
		return false
	}
	return true
}

func InputMove(board [][]string) (int, int) {
	var x, y int
	for {
		fmt.Scanf("%d%d", &x, &y)
		if isValidMove(x, y) {
			break
		}
	}
	return x, y
}
