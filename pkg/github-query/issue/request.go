package issue

import (
	"github.com/collector-for-GitHub/pkg/github-query/types"
)

type IssuesRequest struct {
	// Words and/or phrases to gh for. Optional.
	Terms []string
	// Valid labels for OwnerLogin or RepoName. Optional.
	// query returns results that match any of the labels and are not required to match all of the labels.
	Labels []string
	// Where gh will search for Terms. Optional.
	// Options: Body, Comments, Title, or AnyLocation. Default is AnyLocation.
	SearchIn types.SearchIn
	// query for issues based on open/closed state. Optional.
	// Options: Open, Closed, or AnyIssueState. Default is AnyIssueState.
	State types.IssueState
	// Login for repo's owner or org. ❗ REQUIRED ❗
	OwnerLogin string
	// Repo belonging to OwnerLogin and that will be queried. Optional.
	// If RepoName is not provided, gh will look at all of OwnerLogin's repos, excluding those restricted by the
	// Accessible value or private repos that the client does not have access to with their token.
	RepoName string
	// query includes the public-accessibility status of the repos. Optional.
	// Options: Public, Private, or PublicOrPrivate. Default is PublicOrPrivate.
	Accessible types.Accessible
	// Relation to created/updated times for queried items. Default is any datetime.
	QueryDateTime types.RelativeTime
}

//func (issuesRequest *IssuesRequest) GetGitHubRequest() (*request.GitHubRequest, error) {
//	return request.GetRequestForIssues(issuesRequest)
//}

func (issuesRequest IssuesRequest) GetTerms() []string {
	return issuesRequest.Terms
}

func (issuesRequest IssuesRequest) GetLabels() []string {
	return issuesRequest.Labels
}

func (issuesRequest IssuesRequest) GetRepoName() string {
	return issuesRequest.RepoName
}

func (issuesRequest IssuesRequest) GetAccessible() types.Accessible {
	return issuesRequest.Accessible
}

func (issuesRequest IssuesRequest) GetRelativeTime() types.RelativeTime {
	return issuesRequest.QueryDateTime
}

func (issuesRequest IssuesRequest) GetOwnerLogin() string {
	return issuesRequest.OwnerLogin
}

func (issuesRequest IssuesRequest) GetSearchIn() types.SearchIn {
	return issuesRequest.SearchIn
}

func (issuesRequest IssuesRequest) GetState() types.IssueState {
	return issuesRequest.State
}
