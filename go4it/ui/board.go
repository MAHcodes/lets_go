package ui

import (
	"fmt"

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

func Init() *Board {
	b := &Board{rows: rows, cols: cols}
	b.cells = make([][]string, rows)
	for i := range b.cells {
		b.cells[i] = make([]string, cols)
		for j := range b.cells[i] {
			cell := pieceStyle.Render(empty)
			b.cells[i][j] = cell
		}
	}
	return b
}

func (b *Board) Print() {
	var col string
	for i := range b.cells {
		col = ""
		for j := range b.cells[i] {
			col = lg.JoinHorizontal(lg.Center, col, b.cells[i][j])
		}
		col = colStyle.Render(col)
		fmt.Println(col)
	}
	t := lg.NewStyle().Foreground(boardBG).Render(triangle)
	fmt.Print(" " + t)
	end := lg.PlaceHorizontal(lg.Width(col)-3, lg.Right, t)
	fmt.Println(end)
}

func BoardView() {
	b := Init()
	b.Print()
}
