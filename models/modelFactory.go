package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vanshaj/awot/models/ec2/ec2base"
	"github.com/vanshaj/awot/models/s3/s3base"
)

func GetModels(m tea.Model, serviceType string) tea.Model {
	switch serviceType {
	case "ec2":
		return ec2base.NewEC2Model(m)
	case "s3":
		return s3base.NewS3Model(m)
	}
	return nil
}
