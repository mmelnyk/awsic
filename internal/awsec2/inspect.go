package awsec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func Inspect(id string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Configuration error, " + err.Error())
		return err
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []string{id},
	}

	result, err := GetInstances(context.TODO(), client, input)
	if err != nil {
		return err
	}
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			fmt.Println("Instance-ID : ", *i.InstanceId)
			fmt.Println("Image ID    : ", *i.ImageId)
			fmt.Println("State       : ", i.State.Name)
			if i.PrivateIpAddress != nil {
				fmt.Println("Private IP Address: ", *i.PrivateIpAddress)
			}
			if i.PublicIpAddress != nil {
				fmt.Println("Public IP Address: ", *i.PublicIpAddress)
			}
			if i.PublicDnsName != nil {
				fmt.Println("Public DNS Name: ", *i.PublicDnsName)
			}
			fmt.Println("Tags:")
			for _, v := range i.Tags {
				fmt.Println("\t", *v.Key, " : ", *v.Value)
			}
		}
	}

	return nil
}
