package Util

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func CreateInstance(svc ec2iface.EC2API, name, value *string) (*ec2.Reservation, error) {

	res, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId: 	aws.String("ami-0dbc3d7bc646e8516"),
        InstanceType: aws.String("t2.micro"),
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
	})
	
	if err != nil {
		log.Fatal(err)
	}

	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{res.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key: name,
				Value: value,
			},
		},
	})

	if err != nil {
        return nil, err
    }

	return res, err
}

func GetInstanceDetails(svc ec2iface.EC2API, sess *session.Session) (*ec2.DescribeInstancesOutput, error) {
	res, err := svc.DescribeInstances(nil)
	if err != nil {
        return nil, err
    }

    return res, nil
}

func TerminateInstance(svc ec2iface.EC2API, instanceId string) error {
	_, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceId)},
	})

	if err != nil {
        return err
    }

    return nil
}