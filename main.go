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

		x, y := InputPlace(board)

		Place(board, x, y, player)

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

func isValidPlace(x, y int) bool {
	if x < 1 || 3 < x {
		return false
	}
	if y < 1 || 3 < y {
		return false
	}
	return true
}

func InputPlace(board [][]string) (int, int) {
	var x, y int
	for {
		fmt.Scanf("%d%d", &x, &y)
		if isValidPlace(x, y) {
			break
		}
		fmt.Println("入力はそれぞれ1から3の間です")
	}
	return x, y
}

func Place(board [][]string, x, y, player int) {
	var piece string
	if player == FirstPlayer {
		piece = " X"
	} else {
		piece = " O"
	}

	board[y][x] = piece
}
