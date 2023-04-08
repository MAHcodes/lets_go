package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://mah.codes"

type model struct {
	status int
	err    error
}

type errMsg struct{ err error }

type statusMsg int

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
	return checkServer
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
	}
  return m, nil
}

func (m model) View() string {
  if m.err != nil {
    return fmt.Sprintf("\nwell guess what? we have some serious problems: %s\n", m.err)
  }

  s := fmt.Sprintf("\nSending request to %s, please have a seat...\n", url)

  if m.status > 0 {
    s += fmt.Sprintf("%d: %s!\n", m.status, http.StatusText(m.status))
  }

  s += "\nPress ctrl+c to quit.\n"

  return s
}

func main () {
  _, err := tea.NewProgram(model{}).Run()
  if err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
}
