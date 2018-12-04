package github

import "time"

const RateLimitAsString = "rateLimit{cost,limit,nodeCount,remaining,resetAt}"

type rateLimit struct {
	Cost      int       `json:"cost"`      // how many points the query cost
	Limit     int       `json:"limit"`     // how many points can be used in rate limit window (60 minutes)
	NodeCount int       `json:"nodeCount"` // number of nodes results were in this gh
	Remaining int       `json:"remaining"` // how many point are remaining in the rate limit
	ResetAt   time.Time `json:"resetAt"`   // when the rate limit resets in UTC epoch seconds
}
