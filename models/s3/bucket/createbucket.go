package bucket

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/internal"
)

type S3BucketModel struct {
	TextInput   textinput.Model
	ParentModel tea.Model
	Action      string
}

func NewS3BucketModel(m tea.Model, action string) *S3BucketModel {
	model := textinput.New()
	model.Placeholder = "bucket name"
	model.Focus()
	return &S3BucketModel{
		model, m, action,
	}
}

func (m S3BucketModel) Init() tea.Cmd {
	return nil
}

func (m S3BucketModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			bucketname := m.TextInput.Value()
			internal.Logger.Debug(bucketname)
			return m, tea.Quit
		case "esc":
			return m.ParentModel, nil
		}
	}
	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m S3BucketModel) View() string {
	return m.TextInput.View()
}
