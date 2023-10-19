package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func titleScreen(width, height int) string {
	text := "CEMETERY ESCAPE"

	description := `You  must escape the cemetery.
Search tombstones 󰮢 to find the key.
Then head for the door 󰠚,
but watch out for ghosts 󰊠.

REQUIREMENTS:

- A Nerd Font (https://www.nerdfonts.com/)
- A dark color scheme terminal

MOVE    ARROW KEYS
PAUSE   P
QUIT    Q`
	descriptionStyle := lipgloss.NewStyle().
		PaddingTop(1).
		PaddingRight(0).
		PaddingBottom(1).
		PaddingLeft(0)

	screen := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.subtleOrange.Render(text),
		descriptionStyle.Render(description),
		"Good luck!\n\nPRESS ANY KEY TO START",
	)

	screen = lipgloss.PlaceHorizontal(width, lipgloss.Center, screen)
	screen = lipgloss.PlaceVertical(height, lipgloss.Center, screen)

	return screen
}

func winScreen(width, height int) string {
	text := "CONGRATULATIONS!\n  YOU ESCAPED!\n\n      󰊠 󰊠"

	email := styles.magenta.Render("tom@tomontheinternet.com")
	screen := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.subtleOrange.Render(text),
		"\nWhy not send me an email at\n"+email+"\nto let me know what\nyou thought of the game?\n\n\nPRESS Q TO EXIT",
	)

	screen = lipgloss.PlaceHorizontal(width, lipgloss.Center, screen)
	screen = lipgloss.PlaceVertical(height, lipgloss.Center, screen)

	return screen
}

func deathScreen(width, height int) string {
	text := `A GHOST GOT YOU`

	screen := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.subtleOrange.Render(text),
		"\nPRESS A TO TRY THE LEVEL AGAIN\n\nPRESS Q TO QUIT",
	)

	screen = lipgloss.PlaceHorizontal(width, lipgloss.Center, screen)
	screen = lipgloss.PlaceVertical(height, lipgloss.Center, screen)

	return screen
}

func pausedScreen(width, height int) string {
	text := "PAUSED"
	screen := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.subtleOrange.Render(text),
		"\nPRESS P TO CONTINUE PLAYING",
	)

	screen = lipgloss.PlaceHorizontal(width, lipgloss.Center, screen)
	screen = lipgloss.PlaceVertical(height, lipgloss.Center, screen)

	return screen
}

func tooZoomedIn(width, height int) string {
	text := `Your screen is too small!
Try zooming out a bit.
`

	screen := lipgloss.PlaceHorizontal(width, lipgloss.Center, text)
	screen = lipgloss.PlaceVertical(height, lipgloss.Center, screen)

	return screen
}

func help() {
	text := `
CEMETERY ESCAPE

    Cemetery Escape is a game that you can play in your terminal.
    It was created by Tom.

    Get in touch!

    tom@tomontheinternet.com
`
	fmt.Println(text)
}
