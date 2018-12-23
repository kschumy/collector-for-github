package query

import (
	"fmt"

	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/internal/request"
)

func GetIssues(iqr request.IssueQueryRequest) ([]github.Issue, error) {
	gitHubRequest, err := request.GetRequestForIssues(iqr)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	results, err := getResults(gitHubRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to get results: %s", err)
	}
	return results.GetResultsForIssues(), nil
}
