package querystrings

import (
	"strings"
	"time"

	"github.com/collector-for-GitHub/pkg/github-query/internal/fmtstrings"
	"github.com/collector-for-GitHub/pkg/github-query/types"
)

// TODO: fix logic
func relativeTimeDateToString(r types.RelativeTime) string {
	dateString := &strings.Builder{}
	dateString.WriteString("created:")
	dateString.WriteString(fmtstrings.QuoteMark)

	if r.GetRelative() == types.AnyDateTime || r.GetRelative() == types.AfterDateTime {
		dateString.WriteString(">")
	} else {
		dateString.WriteString("<")
	}
	if r.GetRelative() == types.AnyDateTime {
		// GitHub was founded on February 5, 2008, so there should not be any issues or PRs before this date.
		gitHubStartTime := time.Date(2008, 2, 5, 9, 0, 0, 0, time.UTC)
		dateString.WriteString(formatTime(&gitHubStartTime))
	} else {
		dateTime := r.GetTime()
		dateString.WriteString(formatTime(&dateTime))
	}
	dateString.WriteString(fmtstrings.QuoteMark)

	return dateString.String()
}

//
func formatTime(t *time.Time) string {
	return types.ConvertToUTC(t).Format(time.RFC3339)
}
