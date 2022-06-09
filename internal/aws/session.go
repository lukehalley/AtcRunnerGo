package aws

import (
	"atc-runner/internal/env"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func NewAWSSession() *session.Session {

	AWSProfile := env.LoadEnv("AWS_PROFILE")
	AWSRegion := env.LoadEnv("AWS_REGION")

	creds := credentials.NewEnvCredentials()

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
