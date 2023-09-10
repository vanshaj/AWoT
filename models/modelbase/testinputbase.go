package modelbase

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type BaseTextInputModel struct {
	TextInput   textinput.Model
	ParentModel tea.Model
}

func (m BaseTextInputModel) Init() tea.Cmd {
	return nil
}

func (m BaseTextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		case "esc":
			return m.ParentModel, nil
		}
	}
	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m BaseTextInputModel) View() string {
	return fmt.Sprintf(
		"\n  %s\n\n%s",
		m.TextInput.View(),
		"(esc to go back)",
	) + "\n"
}
