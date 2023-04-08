package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Get  key.Binding
	Help key.Binding
	Quit key.Binding
}

type Styles struct {
	BorderColor lipgloss.Color
	Box         lipgloss.Style
}

type model struct {
	textInput textinput.Model
	spinner   spinner.Model
	help      help.Model
	keys      keyMap
	loading   bool
	status    int
	err       error
	width     int
	height    int
	styles    *Styles
}

type errMsg struct{ err error }

type statusMsg int

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("205")
	s.Box = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		Padding(1).
		BorderStyle(lipgloss.RoundedBorder()).
		Width(70)
	return s
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help},
		{k.Get},
		{k.Quit},
	}
}

var keys = keyMap{
	Get: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("return", "Get url"),
	),
	Help: key.NewBinding(
		key.WithKeys(tea.KeyCtrlH.String()),
		key.WithHelp("ctrl+h", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyCtrlC.String(), tea.KeyEsc.String()),
		key.WithHelp("esc", "Quit"),
	),
}

func initialModel() *model {
	styles := DefaultStyles()

	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(styles.BorderColor)

	ti := textinput.New()
	ti.Placeholder = "example.com"
	label := lipgloss.NewStyle().Foreground(styles.BorderColor)
	ti.Prompt = label.Render("Enter URL  ") + "https://"
	ti.Focus()
	ti.CharLimit = 150

	return &model{
		spinner:   s,
		textInput: ti,
		help:      help.New(),
		keys:      keys,
		styles:    styles,
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

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case statusMsg:
		m.status = int(msg)
		m.loading = false

	case errMsg:
		m.err = msg
		m.loading = false

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Get):
			m.loading = true
			m.textInput.Blur()
			cmd := checkServer(m.textInput.Value())
			return m, cmd

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
		return fmt.Sprintf("\nwe have some serious problems: %s\n", m.err)
	}

	s := ""
	if m.loading {
		s += fmt.Sprintf("%s Sending request to %s, please have a seat...", m.spinner.View(), m.textInput.Value())
	} else if m.status > 0 {
		s += fmt.Sprintf("%s -> %d: %s!", m.textInput.Value(), m.status, http.StatusText(m.status))
	} else {
		s += fmt.Sprintf("%s", m.textInput.View())
	}

	box := m.styles.Box.Render(s)
	help := m.help.View(m.keys)

	app := lipgloss.JoinVertical(lipgloss.Center, box, help)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, app)
}

func main() {
	_, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
