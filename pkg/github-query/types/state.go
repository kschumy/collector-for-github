package types

import "fmt"

type (
	IssueState int
	PRState    int
	State int
)

const (
	AnyIssueState IssueState = iota
	Closed
	Open
)

const (
	AnyPRState PRState = iota
	Approved
	ChangesRequested
	Commented
	Dismissed
	Pending
)

const (
	AnyStateForIssue State = iota
	ClosedIssues
	OpenIssue
	AnyStateForPR
	ApprovedPR
	ChangesRequestedPR
	CommentedPR
	DismissedPR
	PendingPR
)

func (issueState IssueState) IsValid() bool {
	return issueState == AnyIssueState || issueState == Closed || issueState == Open
}

func (prState PRState) IsValid() bool {
	return prState == AnyPRState || prState == Approved || prState == ChangesRequested || prState == Commented || prState == Dismissed || prState == Pending
}

func (state State) IsValid() bool {
	return state == AnyStateForIssue || state == ClosedIssues || state == OpenIssue ||
		state == AnyStateForPR || state == ApprovedPR || state == ChangesRequestedPR || state == CommentedPR || state == DismissedPR || state == PendingPR
}

func ConvertPRStateToState(prState PRState) (State, error) {
	switch prState {
	case AnyPRState:
		return AnyStateForPR, nil
	case Approved:
		return ApprovedPR, nil
	case ChangesRequested:
		return ChangesRequestedPR, nil
	case Commented:
		return CommentedPR, nil
	case Dismissed:
		return DismissedPR, nil
	case Pending:
		return PendingPR, nil
	default:
		return -1, fmt.Errorf("cannot convert unkown PRState %v to state", prState)
	}
}

func ConvertIssueStateToState(issueState IssueState)  (State, error) {
	switch issueState {
	case AnyIssueState:
		return AnyStateForIssue, nil
	case Open:
		return OpenIssue, nil
	case Closed:
		return ClosedIssues, nil
	default:
		return -1, fmt.Errorf("cannot convert unkown IssueState %v to state", issueState)
	}
}
