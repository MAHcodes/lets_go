package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const url = "https://mah.codes"

type model struct {
	spinner spinner.Model
	status  int
	err     error
}

type errMsg struct{ err error }

type statusMsg int

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}

func (e errMsg) Error() string {
	return e.err.Error()
}

func (m model) Init() tea.Cmd {
  return m.spinner.Tick
	// return tea.Batch(checkServer, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nwell guess what? we have some serious problems: %s\n", m.err)
	}

  spinner := m.spinner.View() + " "
	if m.status > 0 {
    spinner = ""
  }


	s := fmt.Sprintf("\n%sSending request to %s, please have a seat...\n", spinner, url)

	if m.status > 0 {
		s += fmt.Sprintf("%d: %s!\n", m.status, http.StatusText(m.status))
	}

	s += "\nPress ctrl+c to quit.\n"

	return s
}

func main() {
	_, err := tea.NewProgram(initialModel()).Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
