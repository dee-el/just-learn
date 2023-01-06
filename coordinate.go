package main

type Coordinate struct {
	X int
	Y int
}

// BOTTOM UP LEFT RIGHT
var directions = []Coordinate{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}
