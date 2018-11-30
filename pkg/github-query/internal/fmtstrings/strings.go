package fmtstrings

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const QuoteMark = "\\\\\\\""

func GetInQuotes(s string) string {
	return GetInQuotesWithPrefix("", s)
}

func GetInQuotesWithPrefix(prefix, s string) string {
	return strings.Join([]string{prefix, QuoteMark, s, QuoteMark}, "")
}

// FIXME: this seems very inefficient. Find a better way to do this.
func ToLowercaseUniqueTrimmedList(s []string) []string {
	lowercaseUniqueTrimmedStrings := []string{}
	if len(s) == 0 {
		return lowercaseUniqueTrimmedStrings
	}

	formattedStrSet := map[string]bool{}
	for _, str := range s {
		// FIXME: There must be a better way.
		formattedString, _ := GetTrimmedOrErrorIfRemainingWhiteSpaces(strings.ToLower(str))
		if formattedString == "" || formattedStrSet[formattedString] || formattedString == "or" || formattedString == "not" {
			continue
		}
		formattedStrSet[formattedString] = true
	}

	for key, _ := range formattedStrSet {
		lowercaseUniqueTrimmedStrings = append(lowercaseUniqueTrimmedStrings, key)
	}
	return lowercaseUniqueTrimmedStrings
}

func GetTrimmedOrErrorIfRemainingWhiteSpaces(str string) (string, error) {
	trimmed := strings.TrimSpace(str)
	matched, err := regexp.MatchString("\\s", trimmed)
	if err != nil {
		return "", fmt.Errorf("error intializing regex: %s", err)
	}
	if matched {
		return "", fmt.Errorf("contains whitespace")
	}
	return trimmed, nil
}

func PrintJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
