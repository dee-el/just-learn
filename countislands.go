package main

func CountIslands(grid [][]int) int {
	// constraints:
	// - lands is connected
	count := 0

	gr := Grid{
		Rows:        len(grid),
		Columns:     len(grid[0]),
		Coordinates: grid,
	}

	for i := 0; i < gr.Rows; i++ {
		for j := 0; j < gr.Columns; j++ {
			if gr.Coordinates[i][j] == 1 {
				count++
				findIslandArea(gr, Coordinate{X: i, Y: j})
			}
		}
	}

	return count
}

func findIslandArea(grid Grid, curr Coordinate) {
	if curr.X < 0 ||
		curr.Y < 0 ||
		curr.X >= grid.Rows ||
		curr.Y >= grid.Columns ||
		grid.Coordinates[curr.X][curr.Y] != 1 {
		return
	}

	// flipping the coord as already visited
	grid.Coordinates[curr.X][curr.Y] = -1

	for _, dir := range directions {
		findIslandArea(grid, Coordinate{X: curr.X + dir.X, Y: curr.Y + dir.Y})
	}
}
