package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/vanshaj/awot/internal"
)

func (svc S3Client) CreateBucketViaClient(bucketName string) error {
	internal.Logger.Debugf("Creating bucket %s\n", bucketName)
	//fn := func(opt *s3.Options) { opt.Region = regionName }
	//fns := []func(opt *s3.Options){fn}
	_, err := svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc S3Client) DeleteBucketViaClient(bucketName string) error {
	internal.Logger.Debugf("Delete bucket %s\n", bucketName)
	_, err := svc.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc S3Client) ListBucketsViaClient() (*s3.ListBucketsOutput, error) {
	internal.Logger.Debugf("Listing buckets")
	res, err := svc.ListBuckets(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
