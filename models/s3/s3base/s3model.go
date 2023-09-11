package s3base

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	bucketapi "github.com/vanshaj/awot/api/s3/bucket"
	"github.com/vanshaj/awot/internal"
	"github.com/vanshaj/awot/models/modelbase"
	"github.com/vanshaj/awot/models/s3/bucket"
)

type S3Model struct {
	modelbase.BaseListModel
}

func NewS3Model(m tea.Model) *S3Model {
	items := []list.Item{
		modelbase.Item("create-bucket"),
		modelbase.Item("delete-bucket"),
		modelbase.Item("list-buckets"),
	}
	return &S3Model{
		modelbase.BaseListModel{
			List:        list.New(items, modelbase.ItemDelegate{}, modelbase.DefaultWidth, modelbase.ListHeight),
			ParentModel: m,
		},
	}
}

func (m S3Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			switch m.Choice {
			case "create-bucket":
				return bucket.NewS3BucketModel(m, "create-bucket"), nil
			case "delete-bucket":
				return bucket.NewS3BucketModel(m, "delete-bucket"), nil
			case "list-buckets":
				res, err := bucketapi.ListBuckets()
				if err != nil {
					internal.Logger.Debugf("Error during list buckets %s\n", err.Error())
				}
				return bucket.NewS3BucketListModel(m, res), nil
			}
		case "esc":
			return m.ParentModel, nil
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
