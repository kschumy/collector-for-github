package request

import (
	"fmt"

	"github.com/collector-for-github/pkg/github-query/internal/fmtstrings"
	"github.com/collector-for-github/pkg/github-query/types"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

type GenericRequest interface {
	GetTerms() []string
	GetLabels() []string
	GetSearchIn() types.SearchIn
	GetOwnerLogin() string
	GetRepoName() string
	GetAccessible() types.Accessible
	GetRelativeTime() types.RelativeTime
}

type IssueQueryRequest interface {
	GenericRequest
	GetState() types.IssueState
}

type PRQueryRequest interface {
	GenericRequest
	GetState() types.PRState
	IsMerged() bool
}

type gitHubRequest struct {
	// Words and/or phrases to gh for. Optional.
	terms []string
	// Valid labels for OwnerLogin or RepoName. Optional.
	// query returns results that match any of the labels and are not required to match all of the labels.
	labels []string
	// Types of GitHUb objects the gh will apply to. Optional.
	// Options: Results, PRs, or AllTypes. Default is AllTypes.
	objectType types.ObjectType
	// Where gh will search for Terms. Optional.
	// Options: Body, Comments, Title, or AnyLocation. Default is AnyLocation.
	searchIn types.SearchIn
	// query for issues based on open/closed state. Optional.
	// Options: Open, Closed, or AnyIssueState. Default is AnyIssueState.
	state types.State
	// Login for repo's owner or org. ❗ REQUIRED ❗
	ownerLogin string
	// Repo belonging to OwnerLogin and that will be queried. Optional.
	// If RepoName is not provided, gh will look at all of OwnerLogin's repos, excluding those restricted by the
	// Accessible value or private repos that the client does not have access to with their token.
	repoName string
	// query includes the public-accessibility status of the repos. Optional.
	// Options: Public, Private, or PublicOrPrivate. Default is PublicOrPrivate.
	accessible types.Accessible
	// Relation to created/updated times for queried items. Default is any datetime.
	relativeDateTime types.RelativeTime
	// merged...
	merged bool
}

func GetRequestForIssues(iqr IssueQueryRequest) (*gitHubRequest, error) {
	// Assigns value to defaults
	newGitHubRequest, err := getDefaultGitHubRequest(iqr)
	if err != nil {
		return nil, err
	}

	// Assigns value to state
	if !iqr.GetState().IsValid() {
		return nil, fmt.Errorf("invalid IssueState for gitHubRequest: %v", iqr.GetState())
	}
	state, err := types.ConvertIssueStateToState(iqr.GetState())
	if err != nil {
		return nil, err
	}
	newGitHubRequest.state = state

	// Assigns value to objectType
	newGitHubRequest.objectType = types.Issues

	return newGitHubRequest, nil
}

func GetRequestForPullRequest(prqr PRQueryRequest) (*gitHubRequest, error) {
	newGitHubRequest, err := getDefaultGitHubRequest(prqr)
	if err != nil {
		return nil, err
	}

	// Assigns value to state
	if !prqr.GetState().IsValid() {
		return nil, fmt.Errorf("invalid IssueState for gitHubRequest: %v", prqr.GetState())
	}
	state, err := types.ConvertPRStateToState(prqr.GetState())
	if err != nil {
		return nil, err
	}
	newGitHubRequest.state = state

	// Assigns value to objectType
	newGitHubRequest.objectType = types.PRs
	// Assigns value to merged
	newGitHubRequest.merged = prqr.IsMerged()

	return newGitHubRequest, nil
}

// getDefault returns a gitHubRequest with the following fields set if provided a value by qr:
//		- accessible		(returns error if invalid)
//		- labels
// 		- ownerLogin		(returns error if empty or only contains whitespaces)
// 		- repoName			(returns error if contains whitespaces between characters)
//		- relativeDateTime	(returns error if invalid)
//		- searchIn			(returns error if invalid)
//		- terms
func getDefaultGitHubRequest(qr GenericRequest) (*gitHubRequest, error) {
	ownerLogin, err := formatOwnerIfNoError(qr.GetOwnerLogin())
	if err != nil {
		return nil, err
	}
	newRelativeTime, err := types.GetCopyOrDefault(qr.GetRelativeTime())
	if err != nil {
		return nil, err
	}

	// Assigns value to ownerLogin and relativeDateTime
	newGitHubRequest := gitHubRequest{
		ownerLogin:       ownerLogin,
		relativeDateTime: *newRelativeTime,
	}

	// Assigns value to repoName
	repoName, err := formatRepoIfNoError(qr.GetRepoName())
	if err != nil {
		return nil, err
	}
	newGitHubRequest.repoName = repoName

	// Assigns value to searchIn
	if !qr.GetSearchIn().IsValid() {
		return nil, fmt.Errorf("invalid SearchIn for gitHubRequest: %v", qr.GetSearchIn())
	}
	newGitHubRequest.searchIn = qr.GetSearchIn()

	// Assigns value to accessible
	if !qr.GetAccessible().IsValid() {
		return nil, fmt.Errorf("invalid Accessible for gitHubRequest: %v", qr.GetAccessible())
	}
	newGitHubRequest.accessible = qr.GetAccessible()

	// Assigns value to terms
	newGitHubRequest.terms = fmtstrings.ToLowercaseUniqueTrimmedList(qr.GetTerms())
	// Assigns value to labels
	newGitHubRequest.labels = fmtstrings.ToLowercaseUniqueTrimmedList(qr.GetLabels())

	logger.Debug("GitHub Request: %+v", newGitHubRequest)
	return &newGitHubRequest, nil
}

func (r gitHubRequest) GetTerms() []string {
	return r.terms
}

func (r gitHubRequest) GetLabels() []string {
	return r.labels
}

func (r gitHubRequest) GetObjectType() types.ObjectType {
	return r.objectType
}

func (r gitHubRequest) GetSearchIn() types.SearchIn {
	return r.searchIn
}

func (r gitHubRequest) GetState() types.State {
	return r.state
}

func (r gitHubRequest) GetOwnerLogin() string {
	return r.ownerLogin
}

func (r gitHubRequest) GetRepoName() string {
	return r.repoName
}

func (r gitHubRequest) GetAccessible() types.Accessible {
	return r.accessible
}

func (r gitHubRequest) GetRelativeTime() types.RelativeTime {
	return r.relativeDateTime
}

func (r gitHubRequest) GetLabelAtIndex(index int) (string, error) {
	if index >= len(r.labels) || index < 0 {
		return "", fmt.Errorf("invalid index to set labels")
	}
	return r.labels[index], nil
}

func formatOwnerIfNoError(ownerLogin string) (string, error) {
	login, err := fmtstrings.GetTrimmedOrErrorIfRemainingWhiteSpaces(ownerLogin)
	if err != nil {
		return "", fmt.Errorf("invalid owner login: %s", ownerLogin)
	}
	if len(login) == 0 {
		return login, fmt.Errorf("invalid owner login: cannot be empty or only whitespace string")
	}
	return login, nil
}

func formatRepoIfNoError(repoName string) (string, error) {
	repo, err := fmtstrings.GetTrimmedOrErrorIfRemainingWhiteSpaces(repoName)
	if err != nil {
		return "", fmt.Errorf("invalid repo name: %s", repoName)
	}
	return repo, nil
}
