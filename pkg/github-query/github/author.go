package github

const AuthorAsString = "author{login}"

type Author struct {
	Login string `json:"login"`
}
