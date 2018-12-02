package query

import (
	"github.com/collector-for-GitHub/pkg/github-query/github"
	"time"
)

type resultsInterface interface {
	GetIssues() []github.Issue
	GetPRs() []github.PR

	GetLastCreatedIssueTime() *time.Time
	GetLastCreatedPRTime() *time.Time
}

type results struct {
	issues []github.Issue
	prs    []github.PR

	lastCreatedIssue *time.Time
	lastCreatedPR    *time.Time
}

func (results *results) GetIssues() []github.Issue {
	return results.issues
}

func (results *results) GetPRs() []github.PR {
	return results.prs
}

func (currResults *results) append(newResults *resultsInterface) {
	if newResults == nil {
		return
	}
	if len((*newResults).GetIssues()) > 0 {
		currLastCreatedIssueTime := currResults.GetLastCreatedIssueTime()
		newResultsLastCreatedIssueTime := (*newResults).GetLastCreatedIssueTime()
		if currLastCreatedIssueTime == nil || newResultsLastCreatedIssueTime.After(*currLastCreatedIssueTime) {
			currResults.lastCreatedIssue = newResultsLastCreatedIssueTime
		}
		currResults.issues = append(currResults.issues, (*newResults).GetIssues()...)
	}
	if len((*newResults).GetPRs()) > 0 {
		currLastCreatedPRTime := currResults.GetLastCreatedPRTime()
		newResultsLastCreatedPRTime := (*newResults).GetLastCreatedPRTime()
		if currLastCreatedPRTime == nil || newResultsLastCreatedPRTime.After(*currLastCreatedPRTime) {
			currResults.lastCreatedPR = newResultsLastCreatedPRTime
		}
		currResults.prs = append(currResults.prs, (*newResults).GetPRs()...)
	}
}

func (results *results) GetLastCreatedIssueTime() *time.Time {
	// QUESTION: do you need currResults.lastCreatedPR.IsZero()?
	if results.lastCreatedIssue == nil || results.lastCreatedIssue.IsZero() {
		if len(results.issues) == 0 {
			return nil
		}
		results.lastCreatedIssue = &(results.issues[len(results.issues)-1]).CreatedAt
	}
	return results.lastCreatedIssue
}

func (results *results) GetLastCreatedPRTime() *time.Time {
	// QUESTION: do you need currResults.lastCreatedPR.IsZero()?
	if results.lastCreatedPR == nil || results.lastCreatedPR.IsZero() {
		if len(results.prs) == 0 {
			return nil
		}
		results.lastCreatedPR = &(results.prs[len(results.prs)-1]).CreatedAt
	}
	return results.lastCreatedPR
}
