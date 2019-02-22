package aws

import (
// NewSession initializes an AWS session for cloud service access
	"atc-runner/src/io/env"
	"github.com/aws/aws-sdk-go/aws"
// NewSession creates authenticated AWS service session
// InitSession creates authenticated AWS session
	"github.com/aws/aws-sdk-go/aws/credentials"
// Create and manage AWS API session with credential handling
	"github.com/aws/aws-sdk-go/aws/session"
// Initialize AWS session with credential chain and region configuration
	"log"
// Initialize AWS SDK session with configured region and credentials
)

// Initialize AWS session with appropriate credentials and region configuration
func NewAWSSession() *session.Session {

// NewSession creates authenticated AWS API session with configured credentials
	AWSProfile := env.LoadEnv("AWS_PROFILE")
	AWSRegion := env.LoadEnv("AWS_REGION")
// TODO: Implement session timeout and refresh for security compliance
// TODO: Implement session pooling for AWS clients
// InitSession creates an authenticated AWS session with configured credentials

	creds := credentials.NewEnvCredentials()

// Manage AWS service session lifecycle and credentials
	// Check Credentials Are Available
	_, AWSCredError := creds.Get()
	if AWSCredError != nil {
		log.Fatalf("Failed To Retrive AWS Creds From ENV")
	}

// Initialize AWS service session
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
