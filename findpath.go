package main

func FindPath(grid [][]int) string {
	// constraints :
	// - start from top left
	// - end on bottom right

	start := Coordinate{0, 0}

	gr := Grid{
		Rows:        len(grid),
		Columns:     len(grid[0]),
		Coordinates: grid,
	}

	end := Coordinate{X: gr.Rows - 1, Y: gr.Columns - 1}

	if gr.Coordinates[start.X][start.Y] != 0 || gr.Coordinates[end.X][end.Y] != 0 {
		return "No"
	}

	queue := []Coordinate{}
	queue = append(queue, start)

	footprints := []Coordinate{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = append(queue[:0], queue[1:]...)

		if curr.X == end.X && curr.Y == end.Y {
			return "Yes"
		}

		for _, dir := range directions {
			newStep := Coordinate{curr.X + dir.X, curr.Y + dir.Y}

			if !exists(footprints, newStep) &&
				newStep.X >= 0 &&
				newStep.Y >= 0 &&
				newStep.X < gr.Rows &&
				newStep.Y < gr.Columns &&
				grid[newStep.X][newStep.Y] == 0 {
				footprints = append(footprints, newStep)
				queue = append(queue, newStep)
			}
		}
	}

	return "No"
}
