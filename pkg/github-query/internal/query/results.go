package query

import (
	"time"

	"github.com/collector-for-github/pkg/github-query/github"
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

	lastCreatedIssueTime *time.Time
	lastCreatedPRTime    *time.Time
}

func (results *results) GetResultsForIssues() []github.Issue {
	return results.issues
}

func (results *results) GetResultsForPRs() []github.PR {
	return results.prs
}

// append add newResults to currResults and updates the appropriate last created time value if needed.
func (currResults *results) append(newResults *resultsInterface) {
	if newResults == nil {
		return
	}
	if len((*newResults).GetResultsForIssues()) > 0 {
		currLastCreatedIssueTime := currResults.GetLastCreatedIssueTime()
		newResultsLastCreatedIssueTime := (*newResults).GetLastCreatedIssueTime()
		if currLastCreatedIssueTime == nil || newResultsLastCreatedIssueTime.After(*currLastCreatedIssueTime) {
			currResults.lastCreatedIssueTime = newResultsLastCreatedIssueTime
		}
		currResults.issues = append(currResults.issues, (*newResults).GetResultsForIssues()...)
	}
	if len((*newResults).GetResultsForPRs()) > 0 {
		currLastCreatedPRTime := currResults.GetLastCreatedPRTime()
		newResultsLastCreatedPRTime := (*newResults).GetLastCreatedPRTime()
		if currLastCreatedPRTime == nil || newResultsLastCreatedPRTime.After(*currLastCreatedPRTime) {
			currResults.lastCreatedPRTime = newResultsLastCreatedPRTime
		}
		currResults.prs = append(currResults.prs, (*newResults).GetResultsForPRs()...)
	}
}

func (results *results) GetLastCreatedIssueTime() *time.Time {
	// QUESTION: do you need currResults.lastCreatedPRTime.IsZero()?
	if results.lastCreatedIssueTime == nil || results.lastCreatedIssueTime.IsZero() {
		if len(results.issues) == 0 {
			return nil
		}
		results.lastCreatedIssueTime = &(results.issues[len(results.issues)-1]).CreatedAt
	}
	return results.lastCreatedIssueTime
}

func (results *results) GetLastCreatedPRTime() *time.Time {
	// QUESTION: do you need currResults.lastCreatedPRTime.IsZero()?
	if results.lastCreatedPRTime == nil || results.lastCreatedPRTime.IsZero() {
		if len(results.prs) == 0 {
			return nil
		}
		results.lastCreatedPRTime = &(results.prs[len(results.prs)-1]).CreatedAt
	}
	return results.lastCreatedPRTime
}
