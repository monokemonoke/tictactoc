package main

import (
	"fmt"
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

		board = Place(board, x, y, player)
		ShowBoard(board)

		if player == FirstPlayer {
			player = SecondPlayer
		} else {
			player = FirstPlayer
		}
	}
}

func InitBoard() [3][3]string {
	var board [3][3]string
	for i := 0; i < 3; i++ {
		board[i] = [3]string{"  ", "  ", "  "}
	}
	return board
}

func ShowBoard(board [3][3]string) {
	fmt.Println("  |  |  |  ")

	for i, row := range board {
		fmt.Println("--+--+--+--")
		line := fmt.Sprintf("%s|%s|%s", row[0], row[1], row[2])
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

func InputPlace(board [3][3]string) (int, int) {
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

func Place(board [3][3]string, x, y, player int) [3][3]string {
	var piece string
	if player == FirstPlayer {
		piece = " X"
	} else {
		piece = " O"
	}

	board[y-1][x-1] = piece
	return board
}
