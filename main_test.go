package main

import (
	"testing"
)

type isValidRangeTestCase struct {
	x      int
	y      int
	expect bool
}

var isValidRangeTests = []isValidRangeTestCase{
	{0, 0, false},
	{1, 0, false},
	{0, 1, false},
	{1, 1, true},
	{3, 3, true},
	{4, 3, false},
	{3, 4, false},
	{4, 4, false},
	{-1, -1, false},
}

func TestIsValidRange(t *testing.T) {
	for _, test := range isValidRangeTests {
		actual := isValidRange(test.x, test.y)
		if actual != test.expect {
			t.Fatalf("isValidRange(%d, %d) expects %t, but got %t\n", test.x, test.y, test.expect, actual)
		}
	}
}

type TestCaseIsValidMove struct {
	board  [3][3]string
	x      int
	y      int
	expect bool
}

func TestIsValidMove(t *testing.T) {
	var isValidMoveTests = []TestCaseIsValidMove{
		{
			board:  [3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
			x:      1,
			y:      1,
			expect: true,
		},
		{
			board:  [3][3]string{{" O", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
			x:      1,
			y:      1,
			expect: false,
		},
		{
			board:  [3][3]string{{"  ", "  ", "  "}, {" O", "  ", "  "}, {"  ", "  ", "  "}},
			x:      1,
			y:      2,
			expect: false,
		},
		{
			board:  [3][3]string{{" X", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
			x:      1,
			y:      1,
			expect: false,
		},
	}

	for i, tt := range isValidMoveTests {
		actual := isValidMove(tt.board, tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("CASE%d: expect %t but actual %t", i, tt.expect, actual)
		}
	}
}

type PlaceTestCase struct {
	board  [3][3]string
	x      int
	y      int
	player int
	expect [3][3]string
}

var placeTests = []PlaceTestCase{
	{
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
		1, 1, FirstPlayer,
		[3][3]string{{" X", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
		1, 1, SecondPlayer,
		[3][3]string{{" O", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
		2, 1, FirstPlayer,
		[3][3]string{{"  ", " X", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
		1, 2, FirstPlayer,
		[3][3]string{{"  ", "  ", "  "}, {" X", "  ", "  "}, {"  ", "  ", "  "}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", "  "}},
		3, 3, FirstPlayer,
		[3][3]string{{"  ", "  ", "  "}, {"  ", "  ", "  "}, {"  ", "  ", " X"}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {" O", " O", " O"}, {"  ", "  ", "  "}},
		2, 1, FirstPlayer,
		[3][3]string{{"  ", " X", "  "}, {" O", " O", " O"}, {"  ", "  ", "  "}},
	},
	{
		[3][3]string{{"  ", "  ", "  "}, {" O", " O", " O"}, {"  ", "  ", "  "}},
		1, 2, FirstPlayer,
		[3][3]string{{"  ", "  ", "  "}, {" X", " O", " O"}, {"  ", "  ", "  "}},
	},
}

func TestPlace(t *testing.T) {
	for index, test := range placeTests {
		actual := place(test.board, test.x, test.y, test.player)

		if len(actual) != 3 {
			t.Fatalf("len(actual) should be 3, but actual %d\n", len(actual))
		}
		for i, row := range actual {
			if len(row) != 3 {
				t.Fatalf("len(actual[%d]) should be 3, but actual %d\n", i, len(row))
			}
			if i >= 3 {
				t.Fatalf("The board Place() returns has board[3] row.")
			}
		}

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				if actual[y][x] != test.expect[y][x] {
					t.Logf("testCase%d\n", index)
					t.Logf("board[%d][%d] should be %s, but actual %s", y, x, test.expect[y][x], actual[y][x])
					t.Fatalf("when Place(%+v,%d,%d,%d)\n", test.board, test.y, test.x, test.player)
				}
			}
		}
	}
}
