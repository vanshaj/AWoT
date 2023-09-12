package bucket

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	s3local "github.com/vanshaj/awot/api/s3"
	"github.com/vanshaj/awot/internal"
	"github.com/vanshaj/awot/models/modelbase"
)

type S3BucketActionModel struct {
	modelbase.BaseSpinnerModel
	Action     string
	BucketName string
	//RegionName string
}

type statusMsg int

type errMsg struct{ err error }

func NewS3BucketActionModel(m tea.Model, action string, name string) *S3BucketActionModel {
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
		s3client := s3local.NewS3Client()
		switch m.Action {
		case "create-bucket":
			err := s3client.CreateBucketViaClient(m.BucketName)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- err
				return
			}
		case "delete-bucket":
			err := s3client.DeleteBucketViaClient(m.BucketName)
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
