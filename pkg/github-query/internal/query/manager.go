package query

import (
	"fmt"
	"github.com/collector-for-GitHub/pkg/github-query/internal/querydoer"
	. "github.com/collector-for-GitHub/pkg/github-query/internal/querystrings"
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

//func getIssues(request *RequestProvider) ([]github.Issue, error) {
//	results, err := getResults(request)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get results: %s", err)
//	}
//	return results.GetIssues(), nil
//}

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

	///    / idk what's up with this
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

//
func (manager *Manager) doQueryForIssues() (resultsInterface, error) {
	queryStringFunc, err := manager.queryStrings.BuildQueryStringFactory(manager.request)
	if err != nil {
		return nil, err
	}
	newDoQueryInput := querydoer.DoQueryInput{
		QueryFactory:      queryStringFunc,
		RelativeTimeStart: (*manager.request).GetRelativeTime(),
		ObjectType: (*manager.request).GetObjectType(),
	}
	return querydoer.DoQuery(newDoQueryInput)
}
