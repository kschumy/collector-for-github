package querydoer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/collector-for-github/pkg/github-query/github"
	"github.com/collector-for-github/pkg/github-query/types"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

const url = "https://api.github.com/graphql"

var gitHubToken = os.Getenv("GITHUB_TOKEN")

// doQuery executes and updates the query as needed to get all the requested results.
func (manager *doQueryManager) doQuery() error {
	logger.Level = 4
	queryResults, expectedTotalResults, err := manager.getResultsFromGitHub()
	if err != nil {
		return err
	}
	numOfNewResults, err := manager.addToResults(&queryResults)
	if err != nil {
		return err
	}
	manager.resultsCount = numOfNewResults
	for expectedTotalResults > manager.resultsCount {
		logger.Info("Found %v/%v results. Continuing to query...\n", manager.resultsCount, expectedTotalResults)

		// Using Github's cursor-based pagination implementation for search using graphql (and its traditional REST API)
		// does not allow for pagination past 1,000 results. Changing the date for the query to get results circumvents
		// this problem. However, if GitHub gets rid of this limit, switching to traditional pagination should be
		// strongly considered.
		if err = manager.paginateDateTime(); err != nil {
			return err
		}

		queryResults, remainingTotalResults, err := manager.getResultsFromGitHub()
		if err != nil {
			return err
		}
		if queryResults == nil || len(queryResults) == 0 {
			return fmt.Errorf("error: got nil results back")
		}
		if expectedTotalResults != remainingTotalResults+manager.resultsCount {
			fmt.Printf(
				"Number of expected results has changed from %v to %v\n",
				expectedTotalResults,
				remainingTotalResults+manager.resultsCount,
			)
			expectedTotalResults = remainingTotalResults + manager.resultsCount
		}

		numOfNewResults, err := manager.addToResults(&queryResults)
		if err != nil {
			return err
		}
		manager.resultsCount = numOfNewResults
	}
	logger.Info("Found %v results\n", expectedTotalResults)
	return nil
}

// getResultsFromGitHub makes a call to GitHUb for results and returns a list of results in json.RawMessage, the number
// of possible results in the query, and an error if there is one.
// The number of possible results in the query may be more than the results produced from the results in json.RawMessage,
// as the GitHub does not return more than 100 results at a time.
func (manager *doQueryManager) getResultsFromGitHub() (json.RawMessage, int, error) {
	bodyString, err := manager.getString(manager.currRelativeTime.GetTime())
	if err != nil {
		return nil, 0, fmt.Errorf("error getting string for query: %#v", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(bodyString))
	if err != nil {
		return nil, 0, fmt.Errorf("error intializing new http issue: %#v", err)
	}

	req.Header.Set("Authorization", "bearer "+gitHubToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.starfire-preview+json")
	req.Header.Set("Accept", "application/vnd.github.ocelot-preview+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error with HTTP response: %#v", err)
	}
	if resp.StatusCode != 200 {
		return nil, 0, fmt.Errorf("error with HTTP status code: %#v", resp)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading response body: %#v", err)
	}

	var githubResults struct {
		Data struct {
			RateLimit struct{} `json:"rateLimit"`
			Search    struct {
				ResultsCount int             `json:"issueCount"` // total number of nodes available, which could be more than returned
				Results      json.RawMessage `json:"nodes"`
			} `json:"search"`
		} `json:"data"`
	}

	err = json.Unmarshal(respBody, &githubResults)
	if err != nil {
		return nil, 0, fmt.Errorf("error unmarshalling response body: %#v", err)
	}

	return githubResults.Data.Search.Results, githubResults.Data.Search.ResultsCount, nil
}

// TODO: DRY up?
// addToResults adds the results from message to their respective list depending on their type and returns a int representing
// the number of results in message and an error if there is one.
func (manager *doQueryManager) addToResults(message *json.RawMessage) (int, error) {
	numOfNewResults := 0

	switch manager.objectType {
	case types.Issues:
		var newList []github.Issue
		err := json.Unmarshal(*message, &newList)
		if err != nil {
			return 0, fmt.Errorf("error while unmarshalling Issues: %s", err)
		}
		numOfNewResults = len(newList)
		manager.issueResults = append(manager.issueResults, newList...)

	case types.PRs:
		var newList []github.PR
		err := json.Unmarshal(*message, &newList)
		if err != nil {
			return 0, fmt.Errorf("error while unmarshalling PRs: %s", err)
		}
		numOfNewResults = len(newList)
		manager.prResults = append(manager.prResults, newList...)

	default:
		return 0, fmt.Errorf("unknown type to unmarshall: %T", manager.objectType)
	}

	return numOfNewResults, nil
}

// NOTE: This assumes the order of results will always be in the direction that is desired for paginating by date. This
// may not be the case if future versions of this program allows for the order to change or for results to be queried
// in a different order.
func (manager *doQueryManager) paginateDateTime() error {
	var newDateTime time.Time
	switch manager.objectType {
	case types.Issues:
		newDateTime = *manager.GetLastCreatedIssueTime()
	case types.PRs:
		newDateTime = *manager.GetLastCreatedPRTime()
	default:
		return fmt.Errorf("unknown type to unmarshall: %T", manager.objectType)
	}
	fmt.Printf("Setting new time to: %#v\n", newDateTime)
	return manager.currRelativeTime.SetTime(newDateTime)
}
