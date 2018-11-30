package awsresource

import (
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"gopkg.in/yaml.v2"
)

// TODO: Refactor same logic out of GetYaml and SaveYaml
func GetYaml() (*Config, error) {
	envVars, err := getAWSVars()
	if err != nil {
		return nil, err
	}
	svc, err := createS3Client(envVars)
	if err != nil {
		return nil, fmt.Errorf("error getting S3 client: %#v", err)
	}

	objectFound, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(*envVars.getBucketName()),
		Key:    aws.String(*envVars.getConfigFileName()),
	})
	if err != nil {
		return nil, err
	}
	defer objectFound.Body.Close()

	var result Config
	if err := yaml.NewDecoder(objectFound.Body).Decode(&result); err != nil {
		fmt.Printf("Error decoding: %#v\n", result)
		return nil, fmt.Errorf("error wit decoding: %#v", err)
	}
	return &result, nil
}

//
func SaveYaml(config *Config) error {
	envVars, err := getAWSVars()
	if err != nil {
		return err
	}
	svc, err := createS3Client(envVars)
	if err != nil {
		return err
	}
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	bodyText := io.ReadSeeker(strings.NewReader(string(yamlData)))

	if _, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(*envVars.getBucketName()),
		Key:    aws.String(*envVars.getConfigFileName()),
		Body:   bodyText,
	}); err != nil {
		return err
	}

	logger.Success("Successfully uploaded %s with time date  %s\n", *envVars.getConfigFileName(), config.GetUpdatedTime())
	return nil

}
