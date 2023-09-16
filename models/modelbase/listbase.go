package modelbase

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vanshaj/awot/internal"
)

const ListHeight = 14
const DefaultWidth = 20

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string { return string(i) }

type ItemDelegate struct{}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type BaseListModel struct {
	List        list.Model
	Choice      string
	Quitting    bool
	ParentModel tea.Model
}

type Option func(*BaseListModel)

func WithList(items ...string) Option {
	return func(m *BaseListModel) {
		internal.Logger.Debugf("len of list is %d", len(items))
		listItems := make([]list.Item, len(items))
		for index, value := range items {
			listItem := Item(value)
			listItems[index] = listItem
		}
		m.List = list.New(listItems, ItemDelegate{}, DefaultWidth, ListHeight)
	}
}

func WithParentModelList(parentModel tea.Model) Option {
	return func(m *BaseListModel) {
		m.ParentModel = parentModel
	}
}

func NewBaseListModel(option ...Option) *BaseListModel {
	model := &BaseListModel{}
	for _, o := range option {
		o(model)
	}
	return model
}

func (m BaseListModel) Init() tea.Cmd {
	internal.Logger.Debug("from base model")
	return nil
}

func (m BaseListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			switch m.Choice {
			case "ec2":
				list := []string{"create-ec2", "create-vpc"}
				ec2Model := NewBaseListModel(
					WithList(list...),
					WithParentModelList(m))
				return ec2Model, nil
			case "s3":
				list := []string{"create-bucket", "delete-bucket", "list-buckets", "put-object"}
				s3Model := NewBaseListModel(
					WithList(list...),
					WithParentModelList(m))
				return s3Model, nil
			case "create-bucket":
				items := []string{"bucket-name", "policy-path", "region-name"}
				createBucketModel := NewBaseTextInputModel(
					WithTextInputs(items...),
					WithParentModelText(m),
					WithActionText("create-bucket"))
				return createBucketModel, nil
			case "delete-bucket":
				items := []string{"bucket-name", "region-name"}
				createBucketModel := NewBaseTextInputModel(
					WithTextInputs(items...),
					WithParentModelText(m),
					WithActionText("delete-bucket"))
				return createBucketModel, nil
			case "list-buckets":
				listBucketModel := NewBaseSpinnerModel(
					WithCustomSpinner(),
					WithActionSpinner("list-buckets"),
					WithParentModelSpinner(m))
				return listBucketModel.Update(msg)
			case "put-object":
				items := []string{"bucket-name", "object-name", "file-path"}
				createBucketModel := NewBaseTextInputModel(
					WithTextInputs(items...),
					WithParentModelText(m),
					WithActionText("put-object"))
				return createBucketModel, nil
			}
		case "esc":
			if m.ParentModel != nil {
				return m.ParentModel, nil
			} else {
				return m, tea.Quit
			}
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m BaseListModel) View() string {
	m.List.SetShowTitle(false)
	m.List.SetShowStatusBar(false)
	return m.List.View()
}
