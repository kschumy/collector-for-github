package issue

import (
	"github.com/collector-for-GitHub/pkg/github-query/types"
)

type PRsRequest struct {
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
	State types.PRState
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
	// Merged represents whether or not the PR has been merged.
	Merged bool
}

//func (prsRequest *PRsRequest) GetGitHubRequest() (*request.GitHubRequest, error) {
//	return request.GetRequestForIssues(prsRequest)
//}

func (prsRequest *PRsRequest) GetTerms() []string {
	return prsRequest.Terms
}

func (prsRequest *PRsRequest) GetLabels() []string {
	return prsRequest.Labels
}

func (prsRequest *PRsRequest) GetRepoName() string {
	return prsRequest.OwnerLogin
}

func (prsRequest *PRsRequest) GetAccessible() types.Accessible {
	return prsRequest.Accessible
}

func (prsRequest *PRsRequest) GetRelativeTime() types.RelativeTime {
	return prsRequest.QueryDateTime
}

func (prsRequest *PRsRequest) GetOwnerLogin() string {
	return prsRequest.OwnerLogin
}

func (prsRequest *PRsRequest) GetSearchIn() types.SearchIn {
	return prsRequest.SearchIn
}

func (prsRequest *PRsRequest) GetState() types.PRState {
	return prsRequest.State
}

func (prsRequest *PRsRequest) IsMerged() bool {
	return prsRequest.Merged
}
