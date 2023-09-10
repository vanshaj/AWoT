package bucket

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/modelbase"
)

type S3BucketModel struct {
	modelbase.BaseTextInputModel
	Action string
}

func NewS3BucketModel(m tea.Model, action string) *S3BucketModel {
	model := textinput.New()
	model.PlaceholderStyle = modelbase.SelectedItemStyle
	model.Focus()
	model.Placeholder = "bucket name"
	return &S3BucketModel{
		modelbase.BaseTextInputModel{
			TextInput:   model,
			ParentModel: m,
		},
		action,
	}
}

func (m S3BucketModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			bucketName := m.TextInput.Value()
			m.TextInput.SetValue("")
			return NewS3BucketActionModel(m.ParentModel, m.Action, bucketName), nil
		case "esc":
			return m.ParentModel, nil
		}
	}
	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}
