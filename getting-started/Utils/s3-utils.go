package Utils

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BucketUtils struct {
	S3Client *s3.Client
}

func (util BucketUtils) ListBucket() ([]types.Bucket, error) {
	result, err := util.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var bucket []types.Bucket

	if err != nil {
		log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	} else {
		bucket = result.Buckets
	}

	return  bucket, err
} 

func (util BucketUtils) CreateBucket(name string, region string) error {
	_, err := util.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})

	if err != nil {
		log.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
			name, region, err)
	} else {
		log.Println("Creation of bucket successfulll")
	}
	return err

}

func (util BucketUtils) DeleteBucket(bucketName string) error {
	_, err := util.S3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName)})
	if err != nil {
		log.Printf("Couldn't delete bucket %v. Here's why: %v\n", bucketName, err)
	}
	log.Printf("Successfully deleted bucket with name: %v",bucketName)

	return err
}