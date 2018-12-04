package query

import (
	"github.com/collector-for-GitHub/pkg/github-query/github"
	"github.com/collector-for-GitHub/pkg/github-query/issue"
	"github.com/collector-for-GitHub/pkg/github-query/types"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"time"
)

func QueryForIssues(startTime time.Time) ([]github.Issue, error) {
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
