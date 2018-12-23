package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/collector-for-github/internal/awsresource"
	"github.com/collector-for-github/internal/post"
	"github.com/collector-for-github/query"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	lambda.Start(queryAndPost)
}

func queryAndPost() (Response, error) {
	logger.Level = 3
	configFile, err := awsresource.GetYaml()
	if err != nil {
		return getStandardErrorResponse("config", err)
	}

	issues, err := query.QueryForIssues(configFile.GetUpdatedTime())
	if err != nil {
		return getStandardErrorResponse("query", err)
	}

	err = post.PostAllIssues(issues, configFile)
	awsresource.SaveYaml(configFile)
	if err != nil {
		return getStandardErrorResponse("posting", err)
	}

	return Response{
		Message: fmt.Sprintf(
			"Finished request from %#v, with last recorded time %#v",
			time.Now().String(),
			configFile.GetUpdatedTime().String(),
		),
		Ok: true,
	}, nil
}

func getStandardErrorResponse(errorWith string, err error) (Response, error) {
	logger.Critical("Error with %s: %v", errorWith, err)
	return Response{
		Message: fmt.Sprintf("Error with %s from: %#v", errorWith, time.Now().String()),
		Ok:      false,
	}, err
}
