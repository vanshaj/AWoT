package bucket

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/modelbase"
)

type S3BucketModel struct {
	BucketNameInput modelbase.BaseTextInputModel
	//BucketRegionInput modelbase.BaseTextInputModel
	Action string
}

func NewS3BucketModel(m tea.Model, action string) *S3BucketModel {
	modelName := textinput.New()
	modelName.PlaceholderStyle = modelbase.SelectedItemStyle
	modelName.Focus()
	modelName.Placeholder = "bucket name"
	//modelRegion := textinput.New()
	//modelRegion.PlaceholderStyle = modelbase.SelectedItemStyle
	//modelRegion.Placeholder = "region name"
	return &S3BucketModel{
		modelbase.BaseTextInputModel{
			TextInput:   modelName,
			ParentModel: m,
		},
		//modelbase.BaseTextInputModel{
		//TextInput:   modelRegion,
		//ParentModel: m,
		//},
		action,
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
		//case "tab":
		//m.BucketNameInput.TextInput.Blur()
		//m.BucketRegionInput.TextInput.Focus()
		case "enter":
			bucketName := m.BucketNameInput.TextInput.Value()
			//bucketRegion := m.BucketRegionInput.TextInput.Value()
			m.BucketNameInput.TextInput.SetValue("")
			//m.BucketRegionInput.TextInput.SetValue("")
			return NewS3BucketActionModel(m.BucketNameInput.ParentModel, m.Action, bucketName), nil
		case "esc":
			return m.BucketNameInput.ParentModel, nil
		}
	}
	m.BucketNameInput.TextInput, cmd = m.BucketNameInput.TextInput.Update(msg)
	//if m.BucketNameInput.TextInput.Focused() {
	//m.BucketNameInput.TextInput, cmd = m.BucketNameInput.TextInput.Update(msg)
	//} else if m.BucketRegionInput.TextInput.Focused() {
	//m.BucketRegionInput.TextInput, cmd = m.BucketRegionInput.TextInput.Update(msg)
	//}
	return m, cmd
}

func (m S3BucketModel) View() string {
	//return fmt.Sprintf("%s\n%s", m.BucketNameInput.TextInput.View(), m.BucketRegionInput.TextInput.View())
	return fmt.Sprintf("%s", m.BucketNameInput.TextInput.View())
}
