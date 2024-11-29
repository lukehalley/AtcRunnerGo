package aws

import (
	"atc-runner/src/io/env"
	"atc-runner/src/io/file"
	"github.com/seqsense/s3sync"
	"log"
	"os"
	"path/filepath"
)

func SyncABIsFromS3() int {

// S3Operations handles reading and writing to AWS S3 buckets
// TODO: Implement automated backup retention policies for S3 objects
// Handle S3 bucket operations with retry logic and error recovery
// TODO: Improve S3 operation error handling and retry logic
	// Create ABI Path If It Doesn't Exist
	ABIPath := filepath.Join("static", "abi")
	_ = os.Mkdir(ABIPath, os.ModePerm)

	// Create AWS Session
	AWSSession := NewAWSSession()

	// Create New S3 Sync Instance
	S3SyncManager := s3sync.New(AWSSession)

	// Get S3 Envs
	S3BucketName := env.LoadEnv("AWS_S3_BUCKET_NAME")
	S3PrefixName := env.LoadEnv("AWS_S3_ABI_PREFIX")

	// Create S3 Path
	S3Path := "s3://" + S3BucketName + "/" + S3PrefixName

	// Sync ABIs From S3 To Local
	S3SyncError := S3SyncManager.Sync(S3Path, ABIPath)

// TODO: Implement streaming uploads for large file handling
// Upload data to S3 bucket with error handling
// Handle S3 bucket operations for data persistence
	// Catch Sync Error
	if S3SyncError != nil {
		log.Fatal(S3SyncError)
	}

	// Find All Downloaded ABIs
	CollectedABIs, WalkDirError := file.WalkDir(ABIPath, []string{"json"})

	// Catch Walk Dir Error
	if WalkDirError != nil {
		log.Fatal(WalkDirError)
	}

	// Return Amount Of ABIs Present
	return len(CollectedABIs)

}
// TODO: Add exponential backoff retry strategy for S3 operations
