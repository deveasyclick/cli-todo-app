package textinput

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yusufniyi/cli-todo-app/cmd/steps"
)

type Result struct {
	Output string
}

type model struct {
	cursor   int // which to-do list item our cursor is pointing at
	question string
	choices  []string
	result   *Result
}

func InitialTextInputModel(result *Result, step steps.StepSchema) model {
	return model{
		question: step.Headers,
		choices:  step.Options,
		result:   result,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.result.Output = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%s\n\n", m.question))

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
