package awsec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"go.melnyk.org/awsic/internal/table"
)

// EC2DescribeInstancesAPI defines the interface for the DescribeInstances function.
// We use this interface to test the function using a mocked service.
type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// GetInstances retrieves information about your Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DescribeInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeInstances.
func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func List() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Configuration error, " + err.Error())
		return err
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}

	result, err := GetInstances(context.TODO(), client, input)
	if err != nil {
		return err
	}

	column := []string{
		"Instance ID",
		"Reservation",
		"Status",
		"Name",
	}

	data := make([][]string, 0)

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			row := make([]string, 0, 8)
			row = append(row, *i.InstanceId, *r.ReservationId, string(i.State.Name), "")
			for _, v := range i.Tags {
				if *v.Key == "Name" {
					row[3] = *v.Value
					break
				}
			}
			data = append(data, row)
		}
	}

	table.Print(column, data)

	return nil
}
