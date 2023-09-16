package modelbase

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	s3local "github.com/vanshaj/awot/api/s3"
	"github.com/vanshaj/awot/internal"
)

type BaseSpinnerModel struct {
	Action      string
	Spinner     spinner.Model
	ParentModel tea.Model
	Data        []string
}

type SpinnerOption func(*BaseSpinnerModel)

func WithActionSpinner(action string) SpinnerOption {
	return func(m *BaseSpinnerModel) {
		m.Action = action
	}
}

func WithCustomSpinner() SpinnerOption {
	return func(m *BaseSpinnerModel) {
		spinnerModel := spinner.New()
		spinnerModel.Spinner = spinner.Dot
		spinnerModel.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
		m.Spinner = spinnerModel
	}
}

func WithParentModelSpinner(parentModel tea.Model) SpinnerOption {
	return func(m *BaseSpinnerModel) {
		m.ParentModel = parentModel
	}
}

func WithDataSpinner(data ...string) SpinnerOption {
	return func(m *BaseSpinnerModel) {
		m.Data = data
	}
}

func NewBaseSpinnerModel(options ...SpinnerOption) *BaseSpinnerModel {
	model := &BaseSpinnerModel{}
	for _, o := range options {
		o(model)
	}
	return model
}

func (m BaseSpinnerModel) Init() tea.Cmd {
	return nil
}

func (m BaseSpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	done := make(chan tea.Model)
	quit := make(chan struct{})
	go func() {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				quit <- struct{}{}
				return
			}
		}
	}()
	go func() {
		s3client := s3local.NewS3Client()
		switch m.Action {
		case "list-buckets":
			s3Client := s3local.NewS3Client()
			res, err := s3Client.ListBucketsViaClient()
			if err != nil {
				internal.Logger.Debugf("Error during list buckets %s\n", err.Error())
			}
			items := make([]string, len(res.Buckets))
			for index, value := range res.Buckets {
				items[index] = *value.Name
			}
			listBucketModel := NewBaseListModel(
				WithList(items...),
				WithParentModelList(m.ParentModel))
			done <- listBucketModel
		case "create-bucket":
			bucketName := m.Data[0]
			policyPath := m.Data[1]
			err := s3client.CreateBucketViaClient(bucketName, policyPath)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- struct{}{}
			}
			done <- m.ParentModel
		case "delete-bucket":
			bucketName := m.Data[0]
			err := s3client.DeleteBucketViaClient(bucketName)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- struct{}{}
			}
			done <- m.ParentModel
		case "put-object":
			bucketName := m.Data[0]
			keyName := m.Data[1]
			filePath := m.Data[2]
			err := s3client.CreateObjectViaClient(bucketName, keyName, filePath)
			if err != nil {
				internal.Logger.Debug(err.Error())
				quit <- struct{}{}
			}
			done <- m.ParentModel
		}
		done <- m
		return
	}()
	select {
	case rmodel := <-done:
		return rmodel, nil
	case <-quit:
		return m.ParentModel, tea.Quit
	}
}

func (m BaseSpinnerModel) View() string {
	str := fmt.Sprintf("\n\n   %s Running your task......\n\n", m.Spinner.View())
	return str
}
