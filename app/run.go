package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/servicemodel"
)

func Run() error {
	model := servicemodel.NewServiceModel()
	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		return err
	}
	return nil
}
