package query

import (
	"fmt"

	"github.com/collector-for-GitHub/pkg/github-query/github"
	. "github.com/collector-for-GitHub/pkg/github-query/internal/request"
)

func GetIssues(iqr IssueQueryRequest) ([]github.Issue, error) {
	request, err := GetRequestForIssues(iqr)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}

	results, err := getResults(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get results: %s", err)
	}
	return results.GetResultsForIssues(), nil
}
