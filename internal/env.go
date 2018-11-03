package internal

import (
	"fmt"
	"os"
)

type awsVars struct {
	awsRegion  string
	bucketName string
	configFile string
}

func getAWSVars() (*awsVars, error) {
	vars := &awsVars{}
	var hasEnv bool

	if vars.awsRegion, hasEnv = os.LookupEnv("AWS_REGION"); !hasEnv {
		return nil, fmt.Errorf("unable to get %q from env", "AWS_REGION")
	}

	if vars.bucketName, hasEnv = os.LookupEnv("BUCKET_NAME"); !hasEnv {
		return nil, fmt.Errorf("unable to get %q from env", "BUCKET_NAME")
	}

	if vars.configFile, hasEnv = os.LookupEnv("CONFIG_FILE_NAME"); !hasEnv {
		return nil, fmt.Errorf("unable to get %q from env", "CONFIG_FILE_NAME")
	}

	fmt.Printf("%#v\n", vars)

	return vars, nil
}

//
func (vars *awsVars) getRegion() *string {
	return &vars.awsRegion
}

//
func (vars *awsVars) getBucketName() *string {
	return &vars.bucketName
}

//
func (vars *awsVars) getConfigFileName() *string {
	return &vars.configFile
}
