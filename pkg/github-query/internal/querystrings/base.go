package querystrings

import (
	"fmt"
	"strings"

	"github.com/collector-for-github/pkg/github-query/types"
)

//
func getBase(request *RequestProvider) (string, error) {
	sourceToQuery, err := requestSourceToString((*request).GetOwnerLogin(), (*request).GetRepoName())
	if err != nil {
		return "", err
	}
	queryString := &strings.Builder{}
	queryString.WriteString(sourceToQuery)

	writeToBuilder := func(s string) {
		if s == "" {
			return
		}
		queryString.WriteString(" ")
		queryString.WriteString(s)
	}

	writeToBuilder(objectTypeToString((*request).GetObjectType()))
	writeToBuilder(stateToString((*request).GetState()))
	writeToBuilder(searchInToString((*request).GetSearchIn()))
	writeToBuilder(sortOrderToString())

	return strings.Join([]string{
		sourceToQuery,
		objectTypeToString((*request).GetObjectType()),
		stateToString((*request).GetState()),
		searchInToString((*request).GetSearchIn()),
		sortOrderToString()}, " "), nil
}

func requestSourceToString(ownerLogin, repoName string) (string, error) {
	if ownerLogin == "" {
		return "", fmt.Errorf("must include an OwnerLogin (the username for a user or org to query)")
	}
	// TODO: need quotes?
	if repoName == "" {
		return fmt.Sprintf("org:%v", ownerLogin), nil
	}
	// TODO: change to quotemark
	return fmt.Sprintf("repo:\\\\\\\"%v/%v\\\\\\\"", ownerLogin, repoName), nil
}

// TODO: return error if not a known type
func objectTypeToString(objectType types.ObjectType) string {
	switch objectType {
	case types.Issues:
		return "is:issue"
	case types.PRs:
		return "is:pr"
	default: // AllTypes
		return ""
	}
}

// TODO: These are likely not correct
// TODO: return error if not a known type
func stateToString(state types.State) string {
	switch state {
	case types.ClosedIssues:
		return "is:closed"
	case types.OpenIssue:
		return "is:open"
	case types.ApprovedPR:
		return "is:approved"
	case types.ChangesRequestedPR:
		return "is:changes-requested"
	case types.CommentedPR:
		return "is:commented"
	case types.DismissedPR:
		return "is:dismissed"
	case types.PendingPR:
		return "is:pending"
	default: // AnyStateForIssue and AnyStateForPR
		return ""
	}
}

// TODO: return error if not a known type
func searchInToString(searchIn types.SearchIn) string {
	switch searchIn {
	case types.Body:
		return "in:body"
	case types.Comments:
		return "in:comments"
	case types.Title:
		return "in:title"
	default: // AnyLocation
		return ""
	}
}

// NOTE: will change when future versions allow for user to specify this
func sortOrderToString() string {
	return "sort:created-asc"
}
