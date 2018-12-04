package pullrequest

import (
	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/internal/query"
	"github.com/collector-for-github/pkg/github-query/types"
)

type PRsRequest struct {
	// Words to query for. Optional.
	// Each term must be only one word and contain no extra whitespaces.
	Terms []string
	// Valid labels for org or repo being queries. Optional.
	// Results will match any of these labels, not all of them.
	Labels []string
	// Where query will search for Terms. Optional.
	// Options: Body, Comments, Title, or AnyLocation. Default is AnyLocation.
	SearchIn types.SearchIn
	// query for pull requests based on open/closed state. Optional.
	// TODO: options
	State types.PRState
	// Login for repo's owner or org. ❗ REQUIRED ❗
	OwnerLogin string
	// Repo belonging to OwnerLogin and that will be queried. Optional.
	// If RepoName is not provided, query will look at all of OwnerLogin's repos, excluding those restricted by the
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

// BUG: getting over about 200k results/hour will results in the user exceeding the rate limit.
// TODO: implement a way to return results if this error is encountered.
func (prRequest *PRsRequest) GetPullRequests() ([]github.PR, error) {
	return query.GetPullRequests(prRequest)
}

func (prsRequest *PRsRequest) GetTerms() []string {
	return prsRequest.Terms
}

func (prsRequest *PRsRequest) GetLabels() []string {
	return prsRequest.Labels
}

func (prsRequest *PRsRequest) GetRepoName() string {
	return prsRequest.RepoName
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

func (prsRequest PRsRequest) GetState() types.PRState {
	return prsRequest.State
}

func (prsRequest *PRsRequest) IsMerged() bool {
	return prsRequest.Merged
}
