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
	cursor int
	board  *Board
}

var (
	blue    = lg.Color("4")
	yellow  = lg.Color("3")
	red     = lg.Color("1")
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
			cell := pieceStyle.Render(empty)
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
	var col string
	for i := range m.board.cells {
		col = ""
		for j := range m.board.cells[i] {
			col = lg.JoinHorizontal(lg.Center, col, m.board.cells[i][j])
		}
		col = colStyle.Render(col)
		s += col + "\n"
	}
	t := lg.NewStyle().Foreground(boardBG).Render(triangle)
	s += " " + t
	end := lg.PlaceHorizontal(lg.Width(col)-3, lg.Right, t)
	s += end + "\n"

	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
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
