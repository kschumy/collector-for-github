package querystrings

import (
	"fmt"
	"strings"
)

func (qs *QueryStrings) SetTermsWithOr(input RequestProvider) error {
	if !qs.isInitialized() {
		return fmt.Errorf("cannot set terms with or on uninitilized QueryStrings")
	}
	qs.terms = strings.Join(input.GetTerms(), " OR ")
	return nil
}

func (qs *QueryStrings) SetTermsWithNot(input RequestProvider) error {
	if !qs.isInitialized() {
		return fmt.Errorf("cannot set terms with not on uninitilized QueryStrings")
	}
	qs.terms = termsToStringLeadAndSeparatedByNot(input.GetTerms())
	return nil
}

// BUG: what if they include "NOT" or "OR"?
func termsToStringLeadAndSeparatedByNot(terms []string) string {
	if len(terms) == 0 {
		return ""
	}
	return strings.Join([]string{"NOT ", strings.Join(terms, " NOT ")}, "")
}
