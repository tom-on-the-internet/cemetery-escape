package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type position struct {
	x int
	y int
}

// model is the game model use with bubbletea.
type model struct {
	username       string
	levels         []level
	levelIdx       int
	playerPos      position
	termHeight     int
	termWidth      int
	playerCoolDown int
	windowTooSmall bool
	playerHasKey   bool
	gameWon        bool
	isGameOver     bool
	isPaused       bool
	hasStarted     bool
}

func (m model) level() level {
	return m.levels[m.levelIdx]
}

type level struct {
	ghostMap       map[position]*ghost
	tombstoneMap   map[position]*tombstone
	width          int
	height         int
	door           position
	playerStartPos position
}

type tombstone struct {
	hasKey  bool
	checked bool
}

type nothing struct{}

type tickMsg nothing

func main() {
	// flags? no.
	// any args, show help and quit
	if len(os.Args) > 1 {
		help()
		return
	}

	username := strings.TrimSpace(os.Getenv("USER"))
	if username == "" {
		username = "You"
	}

	levels := makeLevels()
	initialModel := model{
		username:  username,
		levels:    makeLevels(),
		levelIdx:  0,
		playerPos: levels[0].playerStartPos,
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	insanelyCleverFarewellMessage := "Ghoul-bye!"
	fmt.Print("\n" + insanelyCleverFarewellMessage + "\n")
}

func doTick() tea.Cmd {
	cmd := func(t time.Time) tea.Msg {
		return tickMsg(nothing{})
	}

	return tea.Tick(time.Millisecond*10, cmd)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.hasStarted {
			switch msg.String() {
			case "q", "ctrl+c", "esc":
				return m, tea.Quit
			}

			m.hasStarted = true

			return m, doTick()
		}

		if m.gameWon {
			switch msg.String() {
			case "q", "ctrl+c", "esc":
				return m, tea.Quit
			}

			return m, nil
		}

		if m.isGameOver {
			switch msg.String() {
			case "q", "ctrl+c", "esc":
				return m, tea.Quit
			case "a":
				m = m.restartLevel()

				return m, doTick()
			}

			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "up":
			if playerCanMove("up", m) {
				m.playerPos.y--
				m = afterPlayerMove(m)
			}

			return m, nil
		case "down":
			if playerCanMove("down", m) {
				m.playerPos.y++
				m = afterPlayerMove(m)
			}

			return m, nil
		case "left":
			if playerCanMove("left", m) {
				m.playerPos.x--
				m = afterPlayerMove(m)
			}

			return m, nil
		case "right":
			if playerCanMove("right", m) {
				m.playerPos.x++
				m = afterPlayerMove(m)
			}

		case "p":
			m.isPaused = !m.isPaused
			if !m.isPaused {
				return m, doTick()
			}

			return m, nil
		default:
			return m, nil
		}

	case tickMsg:
		if m.gameWon || m.isGameOver || m.isPaused || !m.hasStarted {
			return m, nil
		}

		m = onTick(m)

		return m, doTick()
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		m.windowTooSmall = msg.Height < minWindowHeight || msg.Width < minWindowWidth

		if m.hasStarted && m.windowTooSmall {
			m.isPaused = true
		}

		return m, nil
	}

	return m, nil
}

func (m model) restartLevel() model {
	m.isGameOver = false
	m.isPaused = false
	m.hasStarted = true
	m.levels = makeLevels()
	m.playerPos = m.level().playerStartPos

	return m
}

func playerCanMove(direction string, m model) bool {
	if m.playerCoolDown > 0 || m.gameWon || m.isGameOver || m.isPaused {
		return false
	}

	switch direction {
	case "up":
		m.playerPos.y--
	case "down":
		m.playerPos.y++
	case "left":
		m.playerPos.x--
	case "right":
		m.playerPos.x++
	}

	if t := m.level().tombstoneMap[m.playerPos]; t != nil {
		return false
	}

	playerInBounds := m.playerPos.y > 0 && m.playerPos.y < m.level().height-1 &&
		m.playerPos.x > 0 &&
		m.playerPos.x < m.level().width-1

	return playerInBounds
}

func afterPlayerMove(m model) model {
	m.playerCoolDown = 5

	isNextToDoor := isAdjacent(m.level().door.x, m.level().door.y, m.playerPos.x, m.playerPos.y)

	if m.playerHasKey && isNextToDoor {
		m.levelIdx++

		if m.levelIdx == len(m.levels) {
			m.gameWon = true
			return m
		}

		m.playerPos = m.level().playerStartPos
		m.playerHasKey = false

		return m
	}

	for stonePos, stone := range m.level().tombstoneMap {
		if isAdjacent(stonePos.x, stonePos.y, m.playerPos.x, m.playerPos.y) {
			stone.checked = true

			if stone.hasKey {
				m.playerHasKey = true
			}
		}
	}

	return m
}

func onTick(m model) model {
	// create a new map so we don't wipe out an existing ghost
	// when moving another ghost
	newGhostMap := map[position]*ghost{}

	for currPoint, g := range m.level().ghostMap {
		if g.cooldown > 0 {
			g.cooldown--
			newGhostMap[currPoint] = g

			continue
		}

		if len(g.path) == 0 {
			g.path = findPathForGhost(currPoint, *g, m)
			g.cooldown = 50

			if g.kind == "hunt" {
				g.cooldown = 0
			}

			newGhostMap[currPoint] = g

			continue
		}

		nextPoint := g.path[len(g.path)-1]
		g.path = g.path[:len(g.path)-1]

		if !ghostCanMove(nextPoint, g, newGhostMap, m) {
			g.path = []position{}
			newGhostMap[currPoint] = g

			continue
		}

		g.cooldown = 12
		if g.kind == "hunt" {
			g.cooldown = 8
		}

		newGhostMap[nextPoint] = g
	}

	m.playerCoolDown--

	m.levels[m.levelIdx].ghostMap = newGhostMap

	playerOnGhost := m.level().ghostMap[m.playerPos] != nil
	if playerOnGhost {
		m.isGameOver = true
	}

	return m
}
