package main

import "testing"

type isValidPlaceTestCase struct {
	x      int
	y      int
	expect bool
}

var isValidPlaceTasts = []isValidPlaceTestCase{
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

func TestIsValidMove(t *testing.T) {
	for _, test := range isValidPlaceTasts {
		actual := isValidPlace(test.x, test.y)
		if actual != test.expect {
			t.Fatalf("isValidMove(%d, %d) expects %t, but got %t\n", test.x, test.y, test.expect, actual)
		}
	}
}
