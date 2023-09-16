package s3

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/vanshaj/awot/internal"
)

func (svc S3Client) CreateBucketViaClient(bucketName string, policyPath string) error {
	internal.Logger.Debugf("Creating bucket %s\n", bucketName)
	//fn := func(opt *s3.Options) { opt.Region = regionName }
	//fns := []func(opt *s3.Options){fn}
	_, err := svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}
	_, err = svc.PutPublicAccessBlock(context.TODO(), &s3.PutPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       false,
			BlockPublicPolicy:     false,
			IgnorePublicAcls:      true,
			RestrictPublicBuckets: false,
		},
	})
	if err != nil {
		return err
	}
	f, err := os.Open(policyPath)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(f)
	internal.Logger.Debugf("Bucket policy is %s\n", string(data))
	if err != nil {
		return err
	}
	_, err = svc.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(string(data)),
	})

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

func (svc S3Client) CreateObjectViaClient(bucketName string, keyName string, filePath string) error {
	internal.Logger.Debugf("Creating object")
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketName),
		Body:   f,
	})
	if err != nil {
		return err
	}
	return nil
}
