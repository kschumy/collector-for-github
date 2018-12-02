package issue

import (
	"github.com/collector-for-GitHub/pkg/github-query/github"
	"github.com/collector-for-GitHub/pkg/github-query/internal/query"
)

// Return struct with most recent time
func (issuesRequest *IssuesRequest) GetIssues() ([]github.Issue, error) {
	return query.GetIssues(*issuesRequest)
	//gitHubRequest, err := issuesRequest.GetGitHubRequest()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return query.CollectIssues(gitHubRequest)
}
