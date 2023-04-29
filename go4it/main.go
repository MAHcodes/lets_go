package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

const (
	cols     = 6
	rows     = 7
	piece    = ""
	empty    = ""
	triangle = "󰔶"
)

type Board struct {
	cols  int
	rows  int
	cells [][]string
}

type model struct {
	cursor      int
	player1Turn bool
	quitting    bool
	board       *Board
}

var (
	blue    = lg.Color("4")
	yellow  = lg.Color("3")
	red     = lg.Color("1")
	black   = lg.Color("0")
	boardBG = blue
	p1Clr   = red
	p2Clr   = yellow

	colStyle = lg.NewStyle().
			Padding(0, 1).
			Background(boardBG)

	pieceStyle = lg.NewStyle().
			Background(boardBG).
			Padding(0, 1)
)

func initModel() model {
	b := &Board{rows: rows, cols: cols}
	b.cells = make([][]string, rows)
	for i := range b.cells {
		b.cells[i] = make([]string, cols)
		for j := range b.cells[i] {
			cell := pieceStyle.Foreground(black).Render(piece)
			b.cells[i][j] = cell
		}
	}
	return model{
		cursor: 1,
		board:  b,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() (s string) {
	if m.quitting {
		return ""
	}

	pieceToPlay := ""

	if m.player1Turn {
		pieceToPlay = lg.NewStyle().Foreground(red).Render(piece)
	} else {
		pieceToPlay = lg.NewStyle().Foreground(yellow).Render(piece)
	}

	s += fmt.Sprint(m.cursor, "\n")

	for i := 1; i <= rows; i++ {
		if m.cursor == i {
			s += pieceToPlay
		} else {
			s += "   "
		}
	}
	s = lg.NewStyle().Padding(0, 2).Render(s)
	s += "\n"

	var row string
	for j := range m.board.cells[0] {
		row = ""
		for i := range m.board.cells {
			row = lg.JoinHorizontal(lg.Center, row, m.board.cells[i][j])
		}
		row = colStyle.Render(row)
		s += row + "\n"
	}
	t := lg.NewStyle().Foreground(boardBG).Render(triangle)
	s += " " + t
	end := lg.PlaceHorizontal(lg.Width(row)-3, lg.Right, t)
	s += end + "\n"

	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "l", "right":
			if m.cursor >= rows {
				m.cursor = 1
			} else {
				m.cursor++
			}
		case "h", "left":
			if m.cursor <= 1 {
				m.cursor = rows
			} else {
				m.cursor--
			}
		}
	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
