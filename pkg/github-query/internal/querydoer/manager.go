package querydoer

import (
	"time"

	"github.com/collector-for-GitHub/pkg/github-query/github"
	"github.com/collector-for-GitHub/pkg/github-query/types"
)

type doQueryManager struct {
	currRelativeTime types.RelativeTime
	getString        func(t time.Time) (string, error)
	objectType       types.ObjectType
	resultsCount     int

	issueResults []github.Issue
	prResults    []github.PR
}

func (manager *doQueryManager) GetResultsForIssues() []github.Issue {
	return manager.issueResults
}

func (manager *doQueryManager) GetResultsForPRs() []github.PR {
	return manager.prResults
}

func (manager *doQueryManager) GetLastCreatedIssueTime() *time.Time {
	if len(manager.issueResults) == 0 {
		return nil
	}
	return &(manager.issueResults[len(manager.issueResults)-1]).CreatedAt
}

func (manager *doQueryManager) GetLastCreatedPRTime() *time.Time {
	if len(manager.prResults) == 0 {
		return nil
	}
	return &(manager.prResults[len(manager.prResults)-1]).CreatedAt
}

func (manager *doQueryManager) GetPRs() []github.PR {
	return manager.prResults
}
