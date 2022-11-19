package aws

import (
	"atc-runner/internal/env"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession() *session.Session {

	AWSProfile := env.LoadEnv("AWS_PROFILE")
	AWSRegion := env.LoadEnv("AWS_REGION")

	AWSSession, err := session.NewSessionWithOptions(session.Options{
		Profile: AWSProfile,
		Config: aws.Config{
			Region: aws.String(AWSRegion),
		},
	})

	if err != nil {
		fmt.Printf("Failed To Tnitialize New AWS Session: %v", err)
		return nil
	}

	return AWSSession
}
