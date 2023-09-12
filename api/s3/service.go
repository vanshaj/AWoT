package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/vanshaj/awot/internal"
)

type S3Client struct {
	*s3.Client
}

var s3client *S3Client

func NewS3Client() *S3Client {
	if s3client == nil {
		s3client = &S3Client{s3.NewFromConfig(internal.Config.Config.(aws.Config))}
	}
	return s3client
}
