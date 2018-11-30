package github

import (
	"time"
)

type GitHubList interface {
	Length() int
	CreatedLast() time.Time
	//Append(GitHubList) (GitHubList, error)
	//GetLabelNames() []string
	//GetTitle() string
	//GetAuthorLogin() string
	//GetURL() string
	//GetDateCreated() time.Time
	//GetRepoName() string
	//GetNumber() int
}

type GitHubObject interface {
	GetLabelNames() []string
	GetTitle() string
	GetAuthorLogin() string
	GetURL() string
	GetDateCreated() time.Time
	GetRepoName() string
	GetNumber() int
}

//func AppendLists(listOne, listTwo GitHubList) (GitHubList, error) {
//	//	if !isPR(listOne) {
//	//		return listOne, fmt.Errorf("not a prlist")
//	//	}
//	//	return append(listTwo, listTwo...), nil
//	//}
//	//return append(listOne, listTwo...)
//	//return listOne.Append(listTwo)
//	//return newList
//}

func isIssue(t interface{}) bool {
	switch t.(type) {
	case IssueList:
		return true
	default:
		return false
	}
}