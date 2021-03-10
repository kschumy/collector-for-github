package query

import (
	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/issue"
	"github.com/collector-for-github/pkg/github-query/types"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"time"
)

func QueryForIssues(startTime time.Time, terms []string) ([]github.Issue, error) {
	// TODO: add more validation checks either here or where terms is handled
	if terms == nil || len(terms) == 0 {
		return nil, fmt.Errorf("query terms cannot be nil or empty")
	}
	
	relativeTime, err := types.NewRelativeTime(types.AfterDateTime, startTime)
	if err != nil {
		return nil, err
	}
	logger.Info("Starting query at: %s: ", startTime.String())
	
	issueRequest := issue.IssuesRequest{
		Terms:         terms,
		Labels:        []string{"sig/aws", "area/platform/aws", "area/platform/eks"},
		SearchIn:      types.Title,
		State:         types.Open,
		OwnerLogin:    "kubernetes",
		QueryDateTime: *relativeTime,
	}

	return issueRequest.GetIssues()
}
