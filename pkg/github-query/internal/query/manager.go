package query

import (
	"fmt"
	"github.com/collector-for-github/pkg/github-query/internal/querydoer"
	. "github.com/collector-for-github/pkg/github-query/internal/querystrings"
)

type Manager struct {
	queryStrings *QueryStrings
	request      *RequestProvider
}

//
func initializeManager(request *RequestProvider) (*Manager, error) {
	queryStrings, err := InitializeToDefault(request)
	if err != nil {
		return nil, err
	}

	return &Manager{
		request:      request,
		queryStrings: queryStrings,
	}, nil
}

// BUG: getting over about 200k results/hour will results in the user exceeding the rate limit.
// TODO: implement a way to return results if this error is encountered.
func getResults(request RequestProvider) (*results, error) {
	manager, err := initializeManager(&request)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize manager: %s", err)
	}

	results := results{}

	newResults, err := manager.doQueryForIssues()
	if err != nil {
		return nil, fmt.Errorf("error while querying for issues: %s", err)
	}
	results.append(&newResults)

	// TODO: what's up with this? It was in an earlier version but is it still needed?
	//	// If there are no labels or the query was for everything, there are no more results to be collected.
	//	if len(r.QueryLabels) == 0 || len(r.QueryTerms) == 0 {
	//		return queryResults, nil
	//	}

	if len((*manager.request).GetLabels()) == 0 {
		return &results, nil
	}

	// Do not want duplicate results, so add "NOT" in front of each term
	manager.queryStrings.SetTermsWithNot(*manager.request)

	for i := 0; i < len((*manager.request).GetLabels()); i++ {
		manager.queryStrings.AddLabel(*manager.request, i)
		newResults, err := manager.doQueryForIssues()
		if err != nil {
			return nil, fmt.Errorf("error while querying for issues: %s", err)
		}
		results.append(&newResults)
	}

	return &results, nil
}

func (manager *Manager) doQueryForIssues() (resultsInterface, error) {
	queryStringFunc, err := manager.queryStrings.BuildQueryStringFactory(manager.request)
	if err != nil {
		return nil, err
	}
	newDoQueryInput := querydoer.DoQueryInput{
		QueryFactory:      queryStringFunc,
		RelativeTimeStart: (*manager.request).GetRelativeTime(),
		ObjectType:        (*manager.request).GetObjectType(),
	}
	return querydoer.DoQuery(newDoQueryInput)
}
