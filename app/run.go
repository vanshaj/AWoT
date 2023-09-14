package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/modelbase"
)

func Run() error {
	list := []string{"ec2", "s3"}
	model := modelbase.NewBaseListModel(
		modelbase.WithList(list...))
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		return err
	}
	return nil
}
