package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/issue"
	"github.com/collector-for-github/pkg/github-query/types"
)

func main() {
	currTime := time.Now()
	request, err := getRequest(currTime)
	if err != nil {
		fmt.Printf("error while getting request: %v", err)
	}

	results, err := request.GetIssues()
	if err = writeResults(request, results); err != nil {
		fmt.Printf("error while getting request: %v", err)
	}
}

func writeResults(request *issue.IssuesRequest, results []github.Issue) error {
	requestTime := request.GetRelativeTime().GetTime().Local()
	file, err := os.Create(fmt.Sprintf("results/%d-%02d-%02d-%02d-%02d.txt", requestTime.Year(), requestTime.Month(), requestTime.Day(), requestTime.Minute(), requestTime.Second()))
	time.Sleep(10 * time.Second)
	fmt.Printf("%#v", file)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, fmt.Sprintf("For after %s,\n found %d results for:\n%+v\n\n", requestTime.String(), len(results), request))
	for i, result := range results {
		fmt.Fprintf(file, fmt.Sprintf(
			"%d. Title: %s\n\t- Repo: %s\n\t- User: %s\n\t- Created: %s\n\t- Labels: %s\n\t- URL: %s\n\n",
			i+1,
			result.GetTitle(),
			result.GetRepoName(),
			result.GetAuthorLogin(),
			result.GetDateCreated(),
			strings.Join(result.GetLabelNames(), ", "),
			result.GetURL()),
		)
	}
	return nil
}

func getRequest(currTime time.Time) (*issue.IssuesRequest, error) {
	relativeTime, err := types.NewRelativeTime(types.AfterDateTime, currTime.UTC().AddDate(-2, 0, 0))
	if err != nil {
		return nil, err
	}

	return &issue.IssuesRequest{
		Terms:         []string{"aws"},
		Labels:        []string{"sig/aws", "area/platform/aws"},
		SearchIn:      types.Title,
		State:         types.Open,
		OwnerLogin:    "kubernetes",
		QueryDateTime: *relativeTime,
	}, nil
}
