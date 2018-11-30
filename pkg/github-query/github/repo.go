package github

const RepoAsString = "repository{id,name,owner{id,login,__typename}}"

type repoResults struct {
	Data struct {
		Organization struct {
			Repositories struct {
				TotalCount int     `json:"totalCount"`
				Repos      []*Repo `json:"nodes"`
			} `json:"repositories"`
		} `json:"organization"`
	} `json:"data"`
}

type Repo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		ID       string `json:"id"`
		Login    string `json:"login"`
		TypeName string `json:"__typename"`
	} `json:"owner"`
}

func (repos *repoResults) getRepos() []*Repo {
	return repos.Data.Organization.Repositories.Repos
}
