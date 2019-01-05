package github

import "time"

const ReviewsAsString = "reviews(first:100){totalCount," + "nodes{" + AuthorAsString + ",createdAt,id,state,url}}"

type ReviewsQueryResults struct {
	TotalCount int       `json:"totalCount"`
	Labels     []*Review `json:"nodes"`
}

type Review struct {
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	State     string    `json:"state"`
	Url       string    `json:"url"`
}
