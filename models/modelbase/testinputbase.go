package modelbase

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type BaseTextInputModel struct {
	TextInputs  []textinput.Model
	ParentModel tea.Model
	Action      string
}

type TextOption func(*BaseTextInputModel)

func WithTextInputs(inputs ...string) TextOption {
	return func(m *BaseTextInputModel) {
		textInputs := make([]textinput.Model, len(inputs))
		for index, input := range inputs {
			model := textinput.New()
			model.Placeholder = input
			model.PlaceholderStyle = SelectedItemStyle
			if index == 0 {
				model.Focus()
			}
			textInputs[index] = model
		}
		m.TextInputs = textInputs
	}
}

func WithActionText(action string) TextOption {
	return func(m *BaseTextInputModel) {
		m.Action = action
	}
}

func WithParentModelText(parentModel tea.Model) TextOption {
	return func(m *BaseTextInputModel) {
		m.ParentModel = parentModel
	}
}

func NewBaseTextInputModel(options ...TextOption) *BaseTextInputModel {
	model := &BaseTextInputModel{}
	for _, o := range options {
		o(model)
	}
	return model
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
			switch m.Action {
			case "create-bucket", "delete-bucket", "put-object":
				details := make([]string, len(m.TextInputs))
				for index, textInput := range m.TextInputs {
					details[index] = textInput.Value()
				}
				return NewBaseSpinnerModel(
					WithParentModelSpinner(m.ParentModel),
					WithDataSpinner(details...),
					WithCustomSpinner(),
					WithActionSpinner(m.Action)), nil
			}
		case "tab":
			for index, textInputs := range m.TextInputs {
				if textInputs.Focused() {
					if index < len(m.TextInputs) {
						m.TextInputs[index].Blur()
						m.TextInputs[index+1].Focus()
						break
					}
				}
			}

		case "esc":
			return m.ParentModel, nil
		}
	}
	for index, textInputs := range m.TextInputs {
		if textInputs.Focused() {
			m.TextInputs[index], cmd = m.TextInputs[index].Update(msg)
			break
		}
	}
	return m, cmd
}

func (m BaseTextInputModel) View() string {
	view := ""
	for _, textInput := range m.TextInputs {
		view = fmt.Sprintf("%s\n%s", view, textInput.View())
	}
	return fmt.Sprintf(
		"\n  %s\n\n%s",
		view,
		"(esc to go back)",
	) + "\n"
}
