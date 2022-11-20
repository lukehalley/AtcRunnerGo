package aws

import (
	"atc-runner/internal/env"
	"atc-runner/internal/file"
	"github.com/seqsense/s3sync"
	"log"
	"os"
	"path/filepath"
)

func SyncABIsFromS3() int {

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
