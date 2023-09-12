package bucket

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	bucketapi "github.com/vanshaj/awot/api/s3/bucket"
	"github.com/vanshaj/awot/internal"
	"github.com/vanshaj/awot/models/modelbase"
)

type S3BucketActionModel struct {
	modelbase.BaseSpinnerModel
	Action     string
	BucketName string
	RegionName string
}

type statusMsg int

type errMsg struct{ err error }

func NewS3BucketActionModel(m tea.Model, action string, name string, region string) *S3BucketActionModel {
	model := spinner.New()
	model.Spinner = spinner.Dot
	model.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
	return &S3BucketActionModel{
		modelbase.BaseSpinnerModel{
			Spinner:     model,
			ParentModel: m,
		},
		action,
		name,
		region,
	}
}

func (m S3BucketActionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	done := make(chan struct{})
	quit := make(chan error)
	go func() {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				done <- struct{}{}
				return
			}
		}
	}()
	go func() {
		switch m.Action {
		case "create-bucket":
			err := bucketapi.CreateBucket(m.BucketName)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- err
				return
			}
		case "delete-bucket":
			err := bucketapi.DeleteBucket(m.BucketName)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- err
				return
			}
		}
		done <- struct{}{}
		return
	}()
	select {
	case <-done:
		return m.ParentModel, nil
	case <-quit:
		return m.ParentModel, tea.Quit
	}
}
