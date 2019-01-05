package github

import "time"

// BUG: see below
// NOTE: should be included if this program implements graphql calls that are have heavily-nested queries in the future,
// or if there is a need to get over about 200k results/hour (according to current hits/query for issues and PRs).
// At this time, GitHub is extremely forgiving in how graphql hits are counted against the rate limit. For more info
// about how this rate is calculated, see: https://developer.github.com/v4/guides/resource-limitations/

const RateLimitAsString = "rateLimit{cost,limit,nodeCount,remaining,resetAt}"

type rateLimit struct {
	Cost      int       `json:"cost"`      // how many points the query cost
	Limit     int       `json:"limit"`     // how many points can be used in rate limit window (60 minutes)
	NodeCount int       `json:"nodeCount"` // number of nodes results were in this query
	Remaining int       `json:"remaining"` // how many point are remaining in the rate limit
	ResetAt   time.Time `json:"resetAt"`   // when the rate limit resets in UTC seconds
}
