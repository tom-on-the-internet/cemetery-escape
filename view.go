package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	viewportWidth   = 40
	viewportHeight  = 18
	minWindowWidth  = 42
	minWindowHeight = 20
)

var styles = struct{ yellow, grey, green, red, lightGrey, white, blue, subtleOrange, magenta lipgloss.Style }{
	yellow:       lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00")),
	grey:         lipgloss.NewStyle().Foreground(lipgloss.Color("#777777")),
	green:        lipgloss.NewStyle().Foreground(lipgloss.Color("#228822")),
	red:          lipgloss.NewStyle().Foreground(lipgloss.Color("#cc0022")),
	lightGrey:    lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb")),
	white:        lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")),
	blue:         lipgloss.NewStyle().Foreground(lipgloss.Color("#89cff0")),
	subtleOrange: lipgloss.NewStyle().Foreground(lipgloss.Color("#fadbc3")),
	magenta:      lipgloss.NewStyle().Foreground(lipgloss.Color("#BB00FF")),
}

func (m model) View() string {
	if m.windowTooSmall {
		return tooZoomedIn(m.termWidth, m.termHeight)
	}

	if !m.hasStarted {
		return titleScreen(m.termWidth, m.termHeight)
	}

	if m.isPaused {
		return pausedScreen(m.termWidth, m.termHeight)
	}

	if m.gameWon {
		return winScreen(m.termWidth, m.termHeight)
	}

	if m.isGameOver {
		return deathScreen(m.termWidth, m.termHeight)
	}

	xStart, xEnd, yStart, yEnd := getBounds(m)

	// use a string builder because string concatenation is too slow
	builder := strings.Builder{}
	builder.Grow(viewportWidth * viewportHeight)

	for rowIdx := yStart; rowIdx <= yEnd; rowIdx++ {
		for colIdx := xStart; colIdx <= xEnd; colIdx++ {
			currentPoint := position{x: colIdx, y: rowIdx}
			tile := getTile(m, currentPoint, xStart, xEnd, yStart, yEnd)
			builder.WriteString(tile)
		}
		builder.WriteString("\n")
	}

	gameMap := builder.String()

	screen := lipgloss.PlaceHorizontal(m.termWidth, lipgloss.Center, gameMap)

	if m.termWidth > 110 {
		screen = lipgloss.JoinVertical(
			lipgloss.Center,
			screen,
			"Hey "+m.username+"! Zoom in a bit! This is supposed to be scary.",
		)
	}

	screen = lipgloss.PlaceVertical(m.termHeight, lipgloss.Center, screen)

	return screen
}

// getTile determines what should be rendered in this position.
func getTile(m model, pos position, xStart, xEnd, yStart, yEnd int) string {
	if pos == m.playerPos {
		return styles.blue.Render("")
	}

	if t := m.level().tombstoneMap[pos]; t != nil {
		if t.checked && t.hasKey {
			return styles.yellow.Render("󰮢")
		}

		if t.checked {
			return styles.grey.Render("󰮢")
		}

		return styles.lightGrey.Render("󰮢")
	}

	if g := m.level().ghostMap[pos]; g != nil {
		if g.kind == "wander" {
			return styles.white.Render("󰊠")
		}

		if g.kind == "follow" {
			return styles.green.Render("󰊠")
		}

		if g.kind == "hunt" {
			return styles.red.Render("󰊠")
		}
	}

	if pos.y == m.level().door.y && pos.x == m.level().door.x {
		if m.playerHasKey {
			return styles.yellow.Render("󰠚")
		}

		return "󰠚"
	}

	// entrance marker
	if pos.x == 0 && m.level().playerStartPos.x == 1 &&
		pos.y == m.level().playerStartPos.y {
		return "┃"
	} else if pos.x == m.level().width-1 && m.level().playerStartPos.x == m.level().width-2 && pos.y == m.level().playerStartPos.y {
		return "┃"
	} else if pos.y == 0 && m.level().playerStartPos.y == 1 && pos.x == m.level().playerStartPos.x {
		return "━"
	} else if pos.y == m.level().height-1 && m.level().playerStartPos.y == m.level().height-2 && pos.x == m.level().playerStartPos.x {
		return "━"
	}

	if (pos.x == 0 || pos.x == m.level().width-1) && pos.y > 0 &&
		pos.y < m.level().height-1 {
		switch pos.y {
		case yStart:
			return ""
		case yEnd:
			return ""
		default:
			return "│"
		}
	}

	if (pos.y == 0 || pos.y == m.level().height-1) && pos.x > 0 &&
		pos.x < m.level().width-1 {
		switch pos.x {
		case xStart:
			return ""
		case xEnd:
			return ""
		default:
			return "─"
		}
	}

	if pos.y == 0 && pos.x == 0 {
		return "┌"
	}

	if pos.y == 0 && pos.x == m.level().width-1 {
		return "┐"
	}

	if pos.y == m.level().height-1 && pos.x == 0 {
		return "└"
	}

	if pos.y == m.level().height-1 && pos.x == m.level().width-1 {
		return "┘"
	}

	return " "
}

// getBounds determines where the viewport starts and stops.
func getBounds(m model) (int, int, int, int) {
	halfViewportWidth := viewportWidth / 2

	xStart := m.playerPos.x - halfViewportWidth
	if xStart < 0 {
		xStart = 0
	}

	xEnd := xStart + viewportWidth
	if xEnd > m.level().width-1 {
		xEnd = m.level().width - 1

		xStart = xEnd - viewportWidth
		if xStart < 0 {
			xStart = 0
		}
	}

	halfViewportHeight := viewportHeight / 2

	yStart := m.playerPos.y - halfViewportHeight
	if yStart < 0 {
		yStart = 0
	}

	yEnd := yStart + viewportHeight
	if yEnd > m.level().height-1 {
		yEnd = m.level().height - 1

		yStart = yEnd - viewportHeight
		if yStart < 0 {
			yStart = 0
		}
	}

	return xStart, xEnd, yStart, yEnd
}
