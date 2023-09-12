package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsConfig struct {
	config.Config
}

var Config *AwsConfig

func NewAwsConfig() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		return err
	}
	Config = &AwsConfig{cfg}
	return nil
}
