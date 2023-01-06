package main

func MaxArea(grid [][]int) int {
	// constraints :
	// - start from top left

	gr := Grid{
		Rows:        len(grid),
		Columns:     len(grid[0]),
		Coordinates: grid,
	}

	maxed := 0

	for i := 0; i < gr.Rows; i++ {
		for j := 0; j < gr.Columns; j++ {

			if gr.Coordinates[i][j] == 1 {
				area := islandMeasurement(gr, Coordinate{X: i, Y: j})
				if maxed < area {
					maxed = area
				}
			}
		}
	}

	return maxed
}

func islandMeasurement(grid Grid, curr Coordinate) int {
	// flipping the coord as already visited
	grid.Coordinates[curr.X][curr.Y] = -1
	area := 1

	for _, dir := range directions {
		coord := Coordinate{curr.X + dir.X, curr.Y + dir.Y}

		if coord.X >= 0 &&
			coord.Y >= 0 &&
			coord.X < grid.Rows &&
			coord.Y < grid.Columns &&
			grid.Coordinates[coord.X][coord.Y] == 1 {
			area += islandMeasurement(grid, coord)
		}
	}

	return area
}
