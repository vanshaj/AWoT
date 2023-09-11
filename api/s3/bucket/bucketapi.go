package bucket

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/vanshaj/awot/internal"
)

func CreateBucket(bucketName string) error {
	cfg := internal.Config.Config.(aws.Config)
	svc := s3.NewFromConfig(cfg)
	internal.Logger.Debugf("Creating bucket %s\n", bucketName)
	_, err := svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteBucket(bucketName string) error {
	cfg := internal.Config.Config.(aws.Config)
	svc := s3.NewFromConfig(cfg)
	internal.Logger.Debugf("Delete bucket %s\n", bucketName)
	_, err := svc.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

func ListBuckets() (*s3.ListBucketsOutput, error) {
	cfg := internal.Config.Config.(aws.Config)
	svc := s3.NewFromConfig(cfg)
	internal.Logger.Debugf("Listing buckets")
	res, err := svc.ListBuckets(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
