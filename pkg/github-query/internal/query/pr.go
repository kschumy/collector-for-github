package query

import (
	"fmt"

	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/internal/request"
)

func GetPullRequests(iqr request.PRQueryRequest) ([]github.PR, error) {
	gitHubRequest, err := request.GetRequestForPullRequest(iqr)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	results, err := getResults(gitHubRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to get results: %s", err)
	}
	return results.GetResultsForPRs(), nil
}
