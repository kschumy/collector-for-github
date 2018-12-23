package querystrings

import (
	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/types"
)

// TODO: change where the strings are held
func querySchemaToString(objectType types.ObjectType) string {
	switch objectType {
	case types.Issues:
		return github.GitHubIssueAsStringWithLeadIn
	case types.PRs:
		return github.PRAsStringWithLeadIn
	default: // AllTypes
		return "" //TODO: error!
	}
}
