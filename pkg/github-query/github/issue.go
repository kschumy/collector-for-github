package github

import (
	"time"
)

const (
	IssuesAsString = "activeLockReason,author{login},authorAssociation,closed,closedAt,comments{totalCount}," +
		"createdAt,id,labels(first:100){nodes{name,updatedAt}},lastEditedAt,locked,number,participants{totalCount}," +
		"publishedAt,repository{id,name},state,title,updatedAt,url"

	GitHubIssueAsStringWithLeadIn = "{issueCount,nodes{...on Issue{" + IssuesAsString + "}}}"
)

//
type Issue struct {
	ActiveLockReason string `json:"activeLockReason"`
	Author           struct {
		Login string `json:"login"`
	} `json:"author"`
	AuthorAssociation string    `json:"authorAssociation"`
	Closed            bool      `json:"closed"`
	ClosedAt          time.Time `json:"closedAt"`
	Comments          struct {
		TotalCount int `json:"closedAt"`
	} `json:"comments"`
	CreatedAt    time.Time          `json:"createdAt"`
	IssueId      string             `json:"id"`
	LabelsList   LabelsQueryResults `json:"labels"`
	LastEditedAt time.Time          `json:"lastEditedAt"`
	Locked       bool               `json:"locked"`
	Number       int                `json:"number"` // Identifies the issue number. Different than IssueId.
	Participants struct {
		TotalCount int `json:"totalCount"`
	} `json:"participants"`
	PublishedAt time.Time `json:"publishedAt"`
	Repository  struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"repository"`
	State     string    `json:"state"` // open or closed
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url       string    `json:"url"`
}

// GetLabelNames returns a list of label names for provided gitHubIssue.
func (gitHubIssue *Issue) GetLabelNames() []string {
	return gitHubIssue.LabelsList.GetNames()
}

func (gitHubIssue *Issue) GetTitle() string {
	return gitHubIssue.Title
}

func (gitHubIssue *Issue) GetAuthorLogin() string {
	return gitHubIssue.Author.Login
}

func (gitHubIssue *Issue) GetURL() string {
	return gitHubIssue.Url
}

func (gitHubIssue *Issue) GetDateCreated() time.Time {
	return gitHubIssue.CreatedAt
}

func (gitHubIssue *Issue) GetRepoName() string {
	return gitHubIssue.Repository.Name
}

func (gitHubIssue *Issue) GetNumber() int {
	return gitHubIssue.Number
}
