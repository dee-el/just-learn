package main

func exists(coordinates []Coordinate, search Coordinate) bool {
	for _, coord := range coordinates {
		if coord.X == search.X && coord.Y == search.Y {
			return true
		}
	}

	return false
}
