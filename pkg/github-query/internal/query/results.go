package query

import (
	"time"

	"github.com/collector-for-GitHub/pkg/github-query/github"
)

//TODO: test this!
type resultsInterface interface {
	GetResultsForIssues() []github.Issue
	GetResultsForPRs() []github.PR

	GetLastCreatedIssueTime() *time.Time
	GetLastCreatedPRTime() *time.Time
}

type results struct {
	issues []github.Issue
	prs    []github.PR

	lastCreatedIssue *time.Time
	lastCreatedPR    *time.Time
}

func (results *results) GetResultsForIssues() []github.Issue {
	return results.issues
}

func (results *results) GetResultsForPRs() []github.PR {
	return results.prs
}

func (currResults *results) append(newResults *resultsInterface) {
	if newResults == nil {
		return
	}
	if len((*newResults).GetResultsForIssues()) > 0 {
		currLastCreatedIssueTime := currResults.GetLastCreatedIssueTime()
		newResultsLastCreatedIssueTime := (*newResults).GetLastCreatedIssueTime()
		if currLastCreatedIssueTime == nil || newResultsLastCreatedIssueTime.After(*currLastCreatedIssueTime) {
			currResults.lastCreatedIssue = newResultsLastCreatedIssueTime
		}
		currResults.issues = append(currResults.issues, (*newResults).GetResultsForIssues()...)
	}
	if len((*newResults).GetResultsForPRs()) > 0 {
		currLastCreatedPRTime := currResults.GetLastCreatedPRTime()
		newResultsLastCreatedPRTime := (*newResults).GetLastCreatedPRTime()
		if currLastCreatedPRTime == nil || newResultsLastCreatedPRTime.After(*currLastCreatedPRTime) {
			currResults.lastCreatedPR = newResultsLastCreatedPRTime
		}
		currResults.prs = append(currResults.prs, (*newResults).GetResultsForPRs()...)
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
