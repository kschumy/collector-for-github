package github

import (
	"time"
)

const (
	PRAsString = "additions," + AuthorAsString + ",authorAssociation," + "," + //+ AssigneesAsString + "," +
		"baseRef{id,name," + RepoAsString + "},changedFiles,closed,closedAt,comments(first:100){totalCount}," +
		"createdAt,deletions,id," + LabelsAsString + ",locked,mergeable,merged,mergedAt," +
		"mergedBy{login},number,permalink," + RepoAsString + "reviews(first:100){totalCount," +
		"nodes{" + AuthorAsString + ",createdAt,id,state,url}},state,title,updatedAt,url"

	PRAsStringWithLeadIn = "{issueCount,nodes{...on PullRequest{" + PRAsString + "}}}"
)

//
type PRQueryResults struct {
	Data struct {
		RateLimit struct{} `json:"rateLimit"`
		Search    struct {
			ResultsCount int   `json:"issueCount"` // total number of nodes available, which could be more than returned
			Results      []*PR `json:"nodes"`
		} `json:"search"`
	} `json:"data"`
}

type PR struct {
	Additions         int    `json:"additions"`
	Author            Author `json:"author"`
	AuthorAssociation string `json:"authorAssociation"`
	//Assignees         Assignees `json:"assignees"`
	BaseReference struct {
		IdFoo      string `json:"id"`
		Name       string `json:"name"`
		Repository Repo   `json:"repository"`
	} `json:"baseRef"`
	ChangedFiles int       `json:"changedFiles"`
	Closed       bool      `json:"closed"`
	ClosedAt     time.Time `json:"closedAt"`
	Comments     struct {
		TotalCount int `json:"totalCount"`
	} `json:"comments"`
	CreatedAt  time.Time          `json:"createdAt"`
	Deletions  int                `json:"deletions"`
	Id         string             `json:"id"`
	LabelsList LabelsQueryResults `json:"labels"`
	Locked     bool               `json:"locked"`
	Mergeable  string             `json:"mergeable"`
	Merged     bool               `json:"merged"`
	MergedAt   time.Time          `json:"mergedAt"`
	MergedBy   struct {
		Login string `json:"login"`
	} `json:"mergedBy"`
	Number     int    `json:"number"`
	Permalink  string `json:"permalink"`
	Repository Repo   `json:"repository"`
	Reviews    struct {
		TotalCount int `json:"totalCount"`
		Review     struct {
			Author    Author    `json:"author"`
			CreatedAt time.Time `json:"createdAt"`
			Id        string    `json:"id"`
			State     string    `json:"state"`
			Url       string    `json:"url"`
		} `json:"nodes"`
	} `json:"reviews"`
	State     string    `json:"state"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url       string    `json:"url"`
}

type PullRequestList []PR

func (prlist PullRequestList) Length() int {
	return len(prlist)
}

func (prlist PullRequestList) CreatedLast() time.Time {
	return prlist[len(prlist)-1].CreatedAt
}

func (prlist PullRequestList) Append(newList PullRequestList) PullRequestList {
	return append(prlist, newList...)
}

//func (prlist PullRequestList) Append(newList GitHubList) (PullRequestList, error) {
//	if !isPR(newList) {
//		return prlist, fmt.Errorf("not a prlist")
//	}
//	return append(prlist, newList...), nil
//}

func isPR(t interface{}) bool {
	switch t.(type) {
	case PullRequestList:
		return true
	default:
		return false
	}
}

//func (prlist PullRequestList) Append(newList PullRequestList) PullRequestList {
//	return append(prlist, newList...)
//}

// GetLabelNames returns a list of label names for provided pr.
func (pr *PR) GetLabelNames() []string {
	return pr.LabelsList.GetNames()
}

//
func (pr *PR) GetTitle() string {
	return pr.Title
}

//
func (pr *PR) GetAuthorLogin() string {
	return pr.Author.Login
}

//
func (pr *PR) GetURL() string {
	return pr.Url
}

//
func (pr *PR) GetDateCreated() time.Time {
	return pr.CreatedAt
}

//
func (pr *PR) GetRepoName() string {
	return pr.Repository.Name
}

//
func (pr *PR) GetNumber() int {
	return pr.Number
}
