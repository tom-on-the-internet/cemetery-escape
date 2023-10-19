package main

import "math/rand"

type ghost struct {
	kind     string
	path     []position
	cooldown int
}

func ghostCanMove(p position, g *ghost, newGhostMap map[position]*ghost, m model) bool {
	if t := m.level().tombstoneMap[p]; t != nil {
		return false
	}

	if ghost := m.level().ghostMap[p]; ghost != nil {
		return false
	}

	if ghost := newGhostMap[p]; ghost != nil {
		return false
	}

	return true
}

func findPathForGhost(p position, g ghost, m model) []position {
	if g.kind == "hunt" && playerNearby(p, m, 20) {
		return huntPath(p, m)
	}

	if g.kind == "follow" && playerNearby(p, m, 20) {
		return followPath(p, m)
	}

	return wanderPath(p, m)
}

func wanderPath(p position, m model) []position {
	var xTarget, yTarget int

	for {
		xTarget = rand.Intn(10) - 5
		yTarget = rand.Intn(10) - 5

		if xTarget+p.x > 0 && xTarget+p.x < m.level().width-1 && yTarget+p.y > 0 &&
			yTarget+p.y < m.level().height-1 {
			break
		}
	}

	xTarget += p.x
	yTarget += p.y
	path := []position{}

	for {
		if xTarget != p.x && yTarget != p.y {
			if rand.Intn(2) == 0 {
				if xTarget > p.x {
					xTarget--
				} else {
					xTarget++
				}
			} else {
				if yTarget > p.y {
					yTarget--
				} else {
					yTarget++
				}
			}
		} else if xTarget != p.x {
			if xTarget > p.x {
				xTarget--
			} else {
				xTarget++
			}
		} else if yTarget != p.y {
			if yTarget > p.y {
				yTarget--
			} else {
				yTarget++
			}
		}

		if xTarget == p.x && yTarget == p.y {
			break
		}

		path = append(path, position{x: xTarget, y: yTarget})
	}

	return path
}

func playerNearby(p position, m model, dist int) bool {
	return abs(p.x-m.playerPos.x) <= dist && abs(p.y-m.playerPos.y) <= dist
}

func followPath(p position, m model) []position {
	xTarget := m.playerPos.x
	yTarget := m.playerPos.y
	path := []position{}

	for xTarget != p.x || yTarget != p.y {
		path = append(path, position{x: xTarget, y: yTarget})
		if xTarget != p.x && yTarget != p.y {
			if rand.Intn(2) == 0 {
				if xTarget > p.x {
					xTarget--
				} else {
					xTarget++
				}
			} else {
				if yTarget > p.y {
					yTarget--
				} else {
					yTarget++
				}
			}
		} else if xTarget != p.x {
			if xTarget > p.x {
				xTarget--
			} else {
				xTarget++
			}
		} else if yTarget != p.y {
			if yTarget > p.y {
				yTarget--
			} else {
				yTarget++
			}
		}
	}

	return path
}

func huntPath(start position, m model) []position {
	goal := m.playerPos

	frontier := priorityQueue{}
	frontier.put(start, 0)

	costSoFar := map[position]int{}
	costSoFar[start] = 0

	cameFrom := map[position]position{}
	cameFrom[start] = start

	for !frontier.empty() {
		current := frontier.get()

		if current == goal {
			break
		}

		neighbors := []position{}

		for _, pos := range getNeighbors(current) {
			if pos.x < 1 || pos.y < 1 || pos.x > m.level().width-2 || pos.y > m.level().height-2 {
				continue
			}

			if stone := m.level().tombstoneMap[pos]; stone != nil {
				continue
			}

			if ghost := m.level().ghostMap[pos]; ghost != nil {
				continue
			}

			neighbors = append(neighbors, pos)
		}

		for _, neighbor := range neighbors {
			newCost := costSoFar[current] + 1
			if cost, ok := costSoFar[neighbor]; ok && newCost >= cost {
				continue
			}

			costSoFar[neighbor] = newCost
			cameFrom[neighbor] = current

			frontier.put(neighbor, newCost+manhattanDistance(neighbor, goal))
		}
	}

	if _, ok := cameFrom[goal]; !ok {
		return []position{}
	}

	path := []position{}
	current := goal

	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}

	return path
}
