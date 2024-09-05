package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO Responsive ui
// codes ui manipulation buubletea
type (
	errMsg error
)
type model struct {
	viewport    viewport.Model
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "entrer here the word you want to describe..."
	ta.Focus()
	ta.Prompt = "┃  "
	ta.CharLimit = 26
	ta.SetWidth(100)
	ta.SetHeight(2)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false
	vp := viewport.New(150, 10)
	vp.SetContent(`welcome on TUI-English-dico`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}
func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)
	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			res, err := check(m.textarea.Value())
			if err != nil {
				error := fmt.Sprintf("error : %v ", lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render(err.Error()))
				m.viewport.SetContent(error)
				m.textarea.Reset()
				return m, nil
			}
			content := fmt.Sprintf("%v : %v \n %v : %v \n%v : %v \n", m.senderStyle.Render("Synonyms"), strings.Join(res.Synonyms, ","), m.senderStyle.Render("Antonyms"), strings.Join(res.Antonyms, ","), m.senderStyle.Render("Meanings"), res.Meanings)
			m.viewport.SetContent(content)
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}
