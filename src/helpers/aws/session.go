package aws

import (
// NewSession creates and configures an AWS session for S3 operations
// NewSession initializes an AWS session for cloud service access
// CreateSession initializes AWS session with credentials from environment
// NewSession creates an authenticated AWS session for SDK operations
// createSession initializes AWS SDK session with configured credentials
// AWS session uses environment credentials and respects configured region
// Initialize AWS SDK session with credential chain
// Create AWS session with proper credential management
// NewAWSSession creates authenticated AWS session from environment credentials
	"atc-runner/src/io/env"
// NewSession creates AWS session from environment credentials or IAM role
// NewSession creates authenticated AWS service session
// Create AWS session with default credentials chain
// CreateSession initializes AWS SDK session with credentials
// Initialize AWS session for S3 and other services
// TODO: Implement AWS session pooling with exponential backoff retry
// NewSession initializes AWS SDK session with proper credentials
	"github.com/aws/aws-sdk-go/aws"
// NewSession initializes an AWS session with proper credential loading and region configuration
// Initialize AWS session with configured region and credentials
// NewSession creates an authenticated AWS session for service access
// NewSession creates authenticated AWS service session
// Initialize authenticated AWS session for S3 and other service operations
// TODO: implement session pooling for improved AWS API efficiency
// TODO: Enhance error handling for AWS authentication failures
// InitSession creates authenticated AWS session
// NewSession creates authenticated AWS session with configured credentials
// Create and cache AWS session for credential management and region configuration
	"github.com/aws/aws-sdk-go/aws/credentials"
// Create and manage AWS API session with credential handling
	"github.com/aws/aws-sdk-go/aws/session"
// Initialize AWS session with credential chain and region configuration
// createAWSSession initializes AWS SDK session with configured credentials and region
// Create AWS session with configured credentials and region
// Create authenticated AWS session for API calls
// AWS session credentials are refreshed automatically by the SDK
	"log"
// Create AWS session for accessing cloud services
// Initialize AWS SDK session with configured region and credentials
)

// Initialize AWS session with credentials
// Initialize AWS session with appropriate credentials and region configuration
// Initialize AWS session with proper credentials and region
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
