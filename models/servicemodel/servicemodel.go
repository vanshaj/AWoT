package servicemodel

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models"
	"github.com/vanshaj/awot/models/modelbase"
)

type ServiceModel struct {
	modelbase.BaseListModel
}

func NewServiceModel() *ServiceModel {
	items := []list.Item{
		modelbase.Item("ec2"),
		modelbase.Item("s3"),
	}
	return &ServiceModel{
		modelbase.BaseListModel{
			List:        list.New(items, modelbase.ItemDelegate{}, modelbase.DefaultWidth, modelbase.ListHeight),
			ParentModel: nil,
		},
	}
}

func (m ServiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.List.SelectedItem().(modelbase.Item)
			if ok {
				m.Choice = string(i)
			}
			return models.GetModels(m, m.Choice), nil
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
