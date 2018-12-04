package querystrings

import (
	"fmt"
	"strings"

	"github.com/collector-for-GitHub/pkg/github-query/internal/fmtstrings"
)

func (qs *QueryStrings) AddLabel(input RequestProvider, index int) error {
	if !qs.isInitialized() {
		return fmt.Errorf("cannot set terms with not on uninitilized QueryStrings")
	}
	label, err := input.GetLabelAtIndex(index)
	if err != nil {
		return err
	}

	qs.labels = addLabelToLabelString(label, input.GetLabels()[:index])
	return nil
}

func addLabelToLabelString(labelToInclude string, existingLabels []string) string {
	newLabel := fmtstrings.GetInQuotesWithPrefix("label:", labelToInclude)
	if len(existingLabels) == 0 {
		return newLabel
	}
	between := fmtstrings.GetInQuotes(" -label:")
	return strings.Join([]string{newLabel, fmtstrings.GetInQuotesWithPrefix(" -label:", strings.Join(existingLabels, between))}, "")
}
