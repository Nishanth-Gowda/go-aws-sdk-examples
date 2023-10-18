package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nishanth-gowda/go-aws-sdk-examples/getting-started/Utils"
)

// main uses the AWS SDK for Go V2 to create an Amazon Simple Storage Service
// (Amazon S3) client and list up to 10 buckets in your account.
// This example uses the default settings specified in your shared credentials
// and config files.
func main() {

	region := "ap-south-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	client := s3.NewFromConfig(config)

	fmt.Println("Please enter the desired S3 Operation ")
	fmt.Printf("\n 1. List Buckets \n 2. Create Bucket \n 3. Delete Bucket\n")
	

	utils := Utils.BucketUtils{
		S3Client: client,
	}

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1 : {

			fmt.Printf("Let's list buckets for your account.\n")
			buckets, err := utils.ListBucket()

			if err != nil {
				log.Printf("Error listing buckets: %v\n", err)
				return
			}

			// Process the list of buckets
			for _, bucket := range buckets {
				fmt.Printf("Bucket Name: %s\n", *bucket.Name)
			}
	}

	case 2 : {
			fmt.Println("Please enter the name to create Bucket in S3")
			var bucketName string
			fmt.Scanln(&bucketName)
			fmt.Println("Please enter the region to create Bucket in S3")
			var region string
			fmt.Scanln(&region)
		
		
			err = utils.CreateBucket(bucketName, region)
			log.Fatal(err)
		}
	
	case 3 : {
			fmt.Println("Please enter the name to delete Bucket in S3")
			var bucketName string
			fmt.Scanln(&bucketName)
			err = utils.DeleteBucket(bucketName)
			if err != nil {
				log.Fatal(err)
			}
		}
	
	default: fmt.Println("Please Enter Valid Input")

	}

}