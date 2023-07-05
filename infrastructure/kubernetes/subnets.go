package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithSharedCredentialsFiles([]string{"default"}),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the VPC client
	svc := ec2.NewFromConfig(cfg)

	SubnetInputs := []*ec2.CreateSubnetInput{}

	SubnetInputs = append(SubnetInputs, &ec2.CreateSubnetInput{
		VpcId:            aws.String("vpc-07491f7eb0091a4c6"),
		AvailabilityZone: aws.String("us-east-1a"),
		CidrBlock:        aws.String("10.20.0.0/24"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("kubernetes-devopsdriver-control-plane-1a"),
					},
				},
			},
		},
	})

	SubnetInputs = append(SubnetInputs, &ec2.CreateSubnetInput{
		VpcId:            aws.String("vpc-07491f7eb0091a4c6"),
		AvailabilityZone: aws.String("us-east-1b"),
		CidrBlock:        aws.String("10.20.1.0/24"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("kubernetes-devopsdriver-node-1b"),
					},
				},
			},
		},
	})

	SubnetInputs = append(SubnetInputs, &ec2.CreateSubnetInput{
		VpcId:            aws.String("vpc-07491f7eb0091a4c6"),
		AvailabilityZone: aws.String("us-east-1c"),
		CidrBlock:        aws.String("10.20.2.0/24"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("kubernetes-devopsdriver-node-1c"),
					},
				},
			},
		},
	})

	SubnetInputs = append(SubnetInputs, &ec2.CreateSubnetInput{
		VpcId:            aws.String("vpc-07491f7eb0091a4c6"),
		AvailabilityZone: aws.String("us-east-1d"),
		CidrBlock:        aws.String("10.20.3.0/24"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSubnet,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("kubernetes-devopsdriver-node-1d"),
					},
				},
			},
		},
	})

	for _, SubnetInput := range SubnetInputs {
		resp, err := svc.CreateSubnet(context.TODO(), SubnetInput)

		if err != nil {
			log.Fatalf("failed to create vpc, %v", err)
		}

		fmt.Println("Subnet ID: ", *resp.Subnet.SubnetId)
	}

}
