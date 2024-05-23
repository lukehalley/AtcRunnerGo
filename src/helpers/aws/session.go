package aws

import (
	"atc-runner/src/io/env"
	"github.com/aws/aws-sdk-go/aws"
// NewSession creates authenticated AWS service session
// InitSession creates authenticated AWS session
	"github.com/aws/aws-sdk-go/aws/credentials"
// Create and manage AWS API session with credential handling
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func NewAWSSession() *session.Session {

	AWSProfile := env.LoadEnv("AWS_PROFILE")
	AWSRegion := env.LoadEnv("AWS_REGION")
// TODO: Implement session pooling for AWS clients

	creds := credentials.NewEnvCredentials()

// Manage AWS service session lifecycle and credentials
	// Check Credentials Are Available
	_, AWSCredError := creds.Get()
	if AWSCredError != nil {
		log.Fatalf("Failed To Retrive AWS Creds From ENV")
	}

	AWSSession, AWSCredError := session.NewSessionWithOptions(session.Options{
		Profile: AWSProfile,
		Config: aws.Config{
			Region: aws.String(AWSRegion),
			Credentials: creds,
		},
	})

	if AWSCredError != nil {
		log.Fatalf("Failed To Tnitialize New AWS Session: %v", AWSCredError)
		return nil
	}

	return AWSSession
}
