package main

import (
	"fmt"
)

const (
	FirstPlayer  = 1
	SecondPlayer = 2
)

func main() {
	board := initBoard()
	showBoard(board)
	player := FirstPlayer
	pieceFor := map[int]string{FirstPlayer: "X", SecondPlayer: "O"}

	for {
		fmt.Printf(" %sを置くプレイヤーの番です\n", pieceFor[player])

		x, y := inputPlace(board)

		board = place(board, x, y, player)
		showBoard(board)

		if player == FirstPlayer {
			player = SecondPlayer
		} else {
			player = FirstPlayer
		}
	}
}

func initBoard() [3][3]string {
	var board [3][3]string
	for i := 0; i < 3; i++ {
		board[i] = [3]string{"  ", "  ", "  "}
	}
	return board
}

func showBoard(board [3][3]string) {
	fmt.Println("  | 1| 2| 3")

	for i, row := range board {
		fmt.Println("--+--+--+--")
		line := fmt.Sprintf("%s|%s|%s", row[0], row[1], row[2])
		line = fmt.Sprintf("%d |", i+1) + line
		fmt.Println(line)
	}
}

func isValidRange(x, y int) bool {
	if x < 1 || 3 < x {
		return false
	}
	if y < 1 || 3 < y {
		return false
	}
	return true
}

func isValidMove(board [3][3]string, x, y int) bool {
	return board[y-1][x-1] == " "
}

func inputPlace(board [3][3]string) (int, int) {
	var x, y int
	for {
		fmt.Scanf("%d%d", &x, &y)
		if isValidRange(x, y) {
			break
		}
		fmt.Println("入力はそれぞれ1から3の間です")
		fmt.Println("例:) 右上に入力したい場合は\"3 1\"と入力")
	}
	return x, y
}

func place(board [3][3]string, x, y, player int) [3][3]string {
	var piece string
	if player == FirstPlayer {
		piece = " X"
	} else {
		piece = " O"
	}

	board[y-1][x-1] = piece
	return board
}
