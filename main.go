package main

import (
	"fmt"
)

const (
	FirstPlayer  = 1 // X で打つプレイヤー
	SecondPlayer = 2 // O で打つプレイヤー
	Draw         = 3 // 引き分け
)

func main() {
	board := initBoard()
	showBoard(board)
	player := FirstPlayer
	pieceFor := map[int]string{FirstPlayer: "X", SecondPlayer: "O"}
	var winner int

	for {
		fmt.Printf(" %sを置くプレイヤーの番です\n", pieceFor[player])

		x, y := inputPlace(board)

		board = place(board, x, y, player)
		showBoard(board)
		if flag, p := isFinish(board); flag {
			winner = p
			break
		}

		if player == FirstPlayer {
			player = SecondPlayer
		} else {
			player = FirstPlayer
		}
	}
	switch winner {
	case FirstPlayer:
		fmt.Println(" Xを置くプレイヤーの勝ちです!")
	case SecondPlayer:
		fmt.Println(" Oを置くプレイヤーの勝ちです!")
	case Draw:
		fmt.Println("引き分けです!")
	default:
		fmt.Println("異常な値が渡されました")
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
	return board[y-1][x-1] == "  "
}

func inputPlace(board [3][3]string) (int, int) {
	var x, y int
	for {
		fmt.Scanf("%d%d", &x, &y)
		if !isValidRange(x, y) {
			fmt.Println("入力はそれぞれ1から3の間です")
			fmt.Println("例:) 右上に入力したい場合は\"3 1\"と入力")
			continue
		}
		fmt.Println(x, y)
		if isValidMove(board, x, y) {
			break
		} else {
			fmt.Printf("(x, y) = (%d, %d) には既に%sがおかれています!\n", x, y, board[y-1][x-1])
		}
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

func isFinish(board [3][3]string) (bool, int) {
	winPlayer := winner(board)
	if winPlayer == FirstPlayer || winPlayer == SecondPlayer {
		return true, winPlayer
	}
	for _, row := range board {
		for _, v := range row {
			if v == "  " {
				return false, Draw
			}
		}
	}
	return true, Draw
}

func winner(b [3][3]string) int {
	playerFor := map[string]int{" X": FirstPlayer, " O": SecondPlayer, "  ": -1 /*Invalid*/}
	for i := 0; i < 3; i++ {
		if b[i][0] != "  " && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return playerFor[b[i][0]]
		}
		if b[0][i] != "  " && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return playerFor[b[0][i]]
		}
	}
	if b[1][1] != "  " && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return playerFor[b[1][1]]
	}
	if b[1][1] != "  " && b[2][0] == b[1][1] && b[1][1] == b[0][2] {
		return playerFor[b[1][1]]
	}
	return Draw
}
