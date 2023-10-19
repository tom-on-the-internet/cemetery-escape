package main

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func isAdjacent(x1, y1, x2, y2 int) bool {
	dx := abs(x1 - x2)
	dy := abs(y1 - y2)

	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

func getNeighbors(pos position) []position {
	return []position{
		{x: pos.x - 1, y: pos.y},
		{x: pos.x + 1, y: pos.y},
		{x: pos.x, y: pos.y - 1},
		{x: pos.x, y: pos.y + 1},
	}
}

func manhattanDistance(a, b position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}
