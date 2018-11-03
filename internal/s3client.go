package internal

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Returns an S3 service client
func createS3Client(vars *awsVars) (*s3.S3, error) {
	newSession, err := session.NewSession(&aws.Config{
		Region: aws.String(*vars.getRegion())},
	)
	if err != nil {
		return nil, fmt.Errorf("error while creating session: %v", err)
	}
	return s3.New(newSession), nil
}
