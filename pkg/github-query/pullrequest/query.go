package issue

import (
	"github.com/collector-for-GitHub/pkg/github-query/internal/query"
	"github.com/collector-for-GitHub/pkg/github-query/types"
)

// Return struct with most recent time
func (prsRequest *PRsRequest) GetPullRequests() (*types.GitHubList, error) {
	gitHubRequest, err := prsRequest.GetGitHubRequest()
	if err != nil {
		return nil, err
	}

	return query.CollectPullRequests(gitHubRequest)
}
