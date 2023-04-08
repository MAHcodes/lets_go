package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput textinput.Model
	spinner   spinner.Model
	loading   bool
	status    int
	err       error
}

type errMsg struct{ err error }

type statusMsg int

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	ti := textinput.New()
	ti.Placeholder = "example.com"
	ti.Prompt = "Enter URL to check > https://"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		spinner:   s,
		textInput: ti,
	}
}

func checkServer(url string) tea.Cmd {
	return func() tea.Msg {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get("https://" + url)
		if err != nil {
			return errMsg{err}
		}
		return statusMsg(res.StatusCode)
	}
}

func (e errMsg) Error() string {
	return e.err.Error()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case statusMsg:
		fmt.Print("running")
		m.status = int(msg)
		m.loading = false
		return m, tea.Quit

	case errMsg:
		fmt.Print("running")
		m.err = msg
		m.loading = false
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEnter:
			m.loading = true
			m.textInput.Blur()
			cmd := checkServer(m.textInput.Value())
			return m, cmd

		case tea.KeyEscape:
			m.textInput.Focus()
			m.loading = false
		default:
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nwell guess what? we have some serious problems: %s\n", m.err)
	}

	s := "\n"
  quit := "\nPress ctrl+c to quit"

	if m.loading {
		s += fmt.Sprintf("%s Sending request to %s, please have a seat...\n", m.spinner.View(), m.textInput.Value())
	} else if m.status > 0 {
		s += fmt.Sprintf("%s -> %d: %s!\n", m.textInput.Value(), m.status, http.StatusText(m.status))
    quit = ""
	} else {
		s += fmt.Sprintf("%s\n", m.textInput.View())
	}

  s += quit

	if m.loading {
		s += " | Esc to cancel."
	}

	s += "\n"

	return s
}

func main() {
	_, err := tea.NewProgram(initialModel()).Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
