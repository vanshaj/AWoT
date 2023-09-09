package ec2base

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/modelbase"
)

type EC2Model struct {
	modelbase.BaseListModel
}

func NewEC2Model(m tea.Model) *EC2Model {
	items := []list.Item{
		modelbase.Item("create-vpc"),
		modelbase.Item("create-ec2"),
	}
	return &EC2Model{
		modelbase.BaseListModel{
			List:        list.New(items, modelbase.ItemDelegate{}, modelbase.DefaultWidth, modelbase.ListHeight),
			ParentModel: m,
		},
	}
}
