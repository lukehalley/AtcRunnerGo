package aws

import (
	"atc-runner/internal/env"
	"github.com/seqsense/s3sync"
	"os"
	"path/filepath"
)

func SyncABIsFromS3() {

	// Create ABI Path If It Doesn't Exist
	ABIPath := filepath.Join("static", "abi")
	_ = os.Mkdir(ABIPath, os.ModePerm)

	AWSSession := NewAWSSession()

	S3SyncManager := s3sync.New(AWSSession)

	S3BucketName := env.LoadEnv("AWS_S3_BUCKET_NAME")
	S3PrefixName := env.LoadEnv("AWS_S3_ABI_PREFIX")

	S3Path := "s3://" + S3BucketName + "/" + S3PrefixName

	// Sync ABIs From S3 To Local
	S3SyncError := S3SyncManager.Sync(S3Path, ABIPath)

	if S3SyncError != nil {
		return
	}
}
