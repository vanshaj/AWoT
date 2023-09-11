package bucket

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/modelbase"
)

type S3BucketListModel struct {
	modelbase.BaseListModel
}

func NewS3BucketListModel(m tea.Model, res *s3.ListBucketsOutput) *S3BucketListModel {
	items := []list.Item{}
	for _, value := range res.Buckets {
		items = append(items, modelbase.Item(*value.Name))
	}
	return &S3BucketListModel{
		modelbase.BaseListModel{
			List:        list.New(items, modelbase.ItemDelegate{}, modelbase.DefaultWidth, modelbase.ListHeight),
			ParentModel: m,
		},
	}
}

//func (m S3BucketListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//go func() {
//res, err := bucketapi.ListBuckets()
//if err != nil {
//internal.Logger.Debugf("Error during list buckets %s\n", err.Error())
//quit <- struct{}{}
//} else {
//done <- res
//}

//}()
//go func() {
//switch msg := msg.(type) {
//case tea.KeyMsg:
//switch keypress := msg.String(); keypress {
//case "ctrl+c", "q":
//return m, tea.Quit
//case "esc":
//return m.ParentModel, nil
//}
//}
//}()
//count := 0
//for {
//select {
//case res := <-done:
//if res == nil {
//return m.ParentModel, nil
//}
//for _, value := range res.Buckets {
//internal.Logger.Debugf("bucket name %s\n", *value.Name)
//m.List.InsertItem(count, modelbase.Item(*value.Name))
//count++
//}
//return m, nil
//case <-quit:
//return m, tea.Quit
//}
//}
//var cmd tea.Cmd
//m.List, cmd = m.List.Update(msg)
////return m, cmd
//}

//func (m S3BucketListModel) View() string {
//m.List.SetShowTitle(false)
//m.List.SetShowStatusBar(false)
//res, err := bucketapi.ListBuckets()
//if err != nil {
//internal.Logger.Debugf("Error during list buckets %s\n", err.Error())
//}
//count := 0
//for _, value := range res.Buckets {
//internal.Logger.Debugf("bucket name %s\n", *value.Name)
//m.List.InsertItem(count, modelbase.Item(*value.Name))
//count++
//}
//return m.List.View()
//}
