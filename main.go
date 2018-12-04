package main

import (
	"fmt"
	"time"

	"github.com/collector-for-GitHub/internal/post"
	"github.com/collector-for-GitHub/query"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	resp, err := manageProgram()

	fmt.Printf("Response: %#v", resp)
	fmt.Printf("Error: %#v", err)
	//lambda.Start(manageProgram)
}

func manageProgram() (Response, error) {
	logger.Level = 3
	//configFile, err := awsresource.GetYaml()
	//if err != nil {
	//	return getStandardErrorResponse("config", err)
	//}

	issues, err := query.QueryForIssues( time.Now().UTC().AddDate(0, -1, 0))

	//issues, err := query.QueryForIssues(configFile.GetUpdatedTime())
	if err != nil {
		return getStandardErrorResponse("query", err)
	}

	err = post.PostAllIssues(issues)//, configFile)
	//awsresource.SaveYaml(configFile)
	if err != nil {
		return getStandardErrorResponse("posting", err)
	}

	return Response{
		Message: fmt.Sprintf(
			"Finished request from %#v, with last recorded time %#v",
			time.Now().String(),
			"hi",
			///configFile.GetUpdatedTime().String(),
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
