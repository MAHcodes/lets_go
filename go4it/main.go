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
	triangle = "󰔶"
)

type Board struct {
	cols  uint8
	rows  uint8
	cells [][]uint8
}

type model struct {
	cursor   uint8
	turn     uint8
	winner   uint8
	quitting bool
	board    *Board
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
			MarginLeft(1).
			Background(boardBG)

	pieceStyle = lg.NewStyle().
			Background(boardBG).
			Padding(0, 1)
)

func initModel() model {
	b := &Board{rows: rows, cols: cols}
	b.cells = make([][]uint8, rows)
	for i := range b.cells {
		b.cells[i] = make([]uint8, cols)
		for j := range b.cells[i] {
			b.cells[i][j] = 0
		}
	}
	return model{
		turn:   1,
		cursor: 1,
		board:  b,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func getPiece(value uint8) (s string) {
	if value == 1 {
		s = pieceStyle.Foreground(red).Render(piece)
	} else if value == 2 {
		s = pieceStyle.Foreground(yellow).Render(piece)
	} else {
		s = pieceStyle.Foreground(black).Render(piece)
	}
	return
}

func (m model) View() (s string) {
	if m.quitting {
		return ""
	}

	if m.winner != 0 {
		return fmt.Sprintf("Player %d win", m.winner)
	}

	s += fmt.Sprint(m.winner, "\n")

	pieceToPlay := ""

	if m.turn == 1 {
		pieceToPlay = lg.NewStyle().Foreground(red).Render(piece)
	} else {
		pieceToPlay = lg.NewStyle().Foreground(yellow).Render(piece)
	}

	for i := uint8(1); i <= rows; i++ {
		if m.cursor == i {
			s += pieceToPlay
		} else {
			s += "   "
		}
	}
	s = lg.NewStyle().Padding(0, 2).MarginLeft(1).Render(s)
	s += "\n"

	var row string
	for j := range m.board.cells[0] {
		row = ""
		for i := range m.board.cells {
			row = lg.JoinHorizontal(lg.Center, row, getPiece(m.board.cells[i][j]))
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

		case " ", "enter":
			col := m.board.cells[m.cursor-1]
			if col[0] != 0 {
				return m, nil
			}

			if col[cols-1] == 0 {
				col[cols-1] = m.turn
			} else {
				for j := range col {
					if col[j] != 0 {
						col[j-1] = m.turn
						break
					}
				}
			}
			if m.turn == 1 {
				m.turn = 2
			} else {
				m.turn = 1
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
