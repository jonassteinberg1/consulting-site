package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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

	// Define default VPC
	input := ec2.CreateVpcInput{
          CidrBlock: aws.String("10.20.0.0/16"),
	}

	// Build vpc
	resp, err := svc.CreateVpc(context.TODO(), &input)

	if err != nil {
	    log.Fatalf("failed to create vpc, %v", err)
	}

	fmt.Println(resp)
}