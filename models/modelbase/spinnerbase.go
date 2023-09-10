package modelbase

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type BaseSpinnerModel struct {
	Spinner     spinner.Model
	ParentModel tea.Model
}

func (m BaseSpinnerModel) Init() tea.Cmd {
	return nil
}

func (m BaseSpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m BaseSpinnerModel) View() string {
	str := fmt.Sprintf("\n\n   %s Running your task......\n\n", m.Spinner.View())
	return str
}
