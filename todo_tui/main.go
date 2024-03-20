package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Todo struct {
	title    string
	selected bool
}

type model struct {
	cursor int
	todos  []Todo
}

func initModel() model {
	return model{
		todos: []Todo{{title: "Watch a movie"}, {title: "Buy food for tonight"}, {title: "Meet Elon Musk personally"}},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.todos[m.cursor].selected = !m.todos[m.cursor].selected
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "GET SHIT DONE!\n\n"

	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ""
		}

		selected := " "
		if m.todos[i].selected {
			selected = "*"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, todo.title)
	}

	s += "\nPress q to quit\n"

	return s
}

func main() {
	p := tea.NewProgram(initModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Are you using Arch? because the app is broken... %s", err)
		os.Exit(1)
	}
}
