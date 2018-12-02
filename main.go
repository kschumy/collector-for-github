package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/collector-for-GitHub/internal/awsresource"
	"github.com/collector-for-GitHub/pkg/github-query/github"
	"github.com/collector-for-GitHub/pkg/github-query/issue"
	"github.com/collector-for-GitHub/pkg/github-query/types"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

const (
	secondsToSleep = 4
)

var url = "https://api.github.com/teams/" + os.Getenv("BOARD_ID") + "/discussions"

type discussionPost struct {
	Author      string    `json:"author"`
	DateCreated time.Time `json:"dateCreated"`
	Labels      string    `json:"labels"`
	Number      int       `json:"number"`
	RepoName    string    `json:"repoName"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
}

//
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	//manageProgram()
	lambda.Start(manageProgram)
}

//
func manageProgram() (Response, error) {
	logger.Level = 3
	configFile, err := awsresource.GetYaml()
	if err != nil {
		return getStandardErrorResponse("config", err)
	}

	issues, err := queryForIssues(configFile.GetUpdatedTime())
	if err != nil {
		return getStandardErrorResponse("query", err)
	}

	err = postAllIssues(issues, configFile)
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

//
func queryForIssues(startTime time.Time) ([]*github.Issue, error) {
	logger.Info("Starting query at: %s: ", startTime.String())
	relativeTime, err := types.NewRelativeTime(types.AfterDateTime, startTime)
	if err != nil {
		return nil, err
	}
	issueRequest := issue.IssuesRequest{
		Terms:         []string{"aws", "eks"},
		Labels:        []string{"sig/aws", "area/platform/aws", "area/platform/eks"},
		SearchIn:      types.Title,
		State:         types.Open,
		OwnerLogin:    "kubernetes",
		QueryDateTime: *relativeTime,
	}

	return issueRequest.GetIssues()
}

//
func postAllIssues(issues []*github.Issue, configFile *awsresource.Config) error {
	for i, issue := range issues {
		if i == 20 {
			return fmt.Errorf("timedout before finishing. Created %v out of %v. Lambda will re-run if it can", i, len(issues))
		}

		issueInfo := getObjectInfo(*issue)
		if err := createPost(issueInfo); err != nil {
			return fmt.Errorf("error creating post (repo %q, #\"%v): %v\"", issue.GetRepoName(), issue.GetNumber(), err)
		}

		logger.Info("%v Created post (repo %q, #\"%v) from %s", i, issue.GetRepoName(), issue.GetNumber(), issue.GetDateCreated().String())
		configFile.SetUpdatedTime(issue.GetDateCreated())
		time.Sleep(secondsToSleep * time.Second) // delay is due to GitHub's rate limits
	}
	return nil
}

//
func createPost(info discussionPost) error {
	payload := strings.NewReader(info.getHTTPPRequestBody())

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return fmt.Errorf("error with creating post: %#v", err)
	}
	req.Header.Add("Accept", "application/vnd.github.echo-preview+json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Authorization", "bearer "+os.Getenv("GITHUB_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error with creating post: %#v", err)
	}
	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		errorMsg := fmt.Sprintf("when creating a post, got status code %v instead of 201", resp.StatusCode)
		if err != nil {
			return fmt.Errorf(errorMsg)
		}
		return fmt.Errorf("%v and message %v", errorMsg, string(body))
	}
	defer resp.Body.Close()
	return nil
}

// FIXME: find a way to encode the title
// TODO: test for other input that will break this program.
func getObjectInfo(issue github.Issue) discussionPost {
	// Unescaped quotes in titles cause a JSON parsing error when creating a new discussion board post.
	return discussionPost{
		Author:      issue.GetAuthorLogin(),
		DateCreated: issue.GetDateCreated(),
		Labels:      strings.Join(issue.GetLabelNames(), ", "),
		Number:      issue.GetNumber(),
		RepoName:    issue.GetRepoName(),
		Title:       strings.Replace(issue.GetTitle(), `"`, "\\\"", -1),
		Url:         issue.GetURL(),
	}
}

//
func (post discussionPost) getHTTPPRequestBody() string {
	return fmt.Sprintf("{\"title\": \"%s, #%v\",\"body\": \"%v\",\"assignees\": [],\"labels\": []}",
		post.RepoName,
		post.Number,
		fmt.Sprintf(
			"- **Title**: %s\\n- **Repo**: %s\\n- **User**: %s\\n- **Created**: %s\\n- **Labels**: %s\\n- **URL**: %s",
			post.Title,
			post.RepoName,
			post.Author,
			post.DateCreated,
			post.Labels,
			post.Url),
	)
}

//
func getStandardErrorResponse(errorWith string, err error) (Response, error) {
	logger.Critical("Error with %s: %v", errorWith, err)
	return Response{
		Message: fmt.Sprintf("Error with %s from: %#v", errorWith, time.Now().String()),
		Ok:      false,
	}, err
}
