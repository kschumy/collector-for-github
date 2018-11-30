package querydoer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/collector-for-GitHub/pkg/github-query/github"
	"github.com/collector-for-GitHub/pkg/github-query/types"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

const url = "https://api.github.com/graphql"

var gitHubToken = os.Getenv("GITHUB_TOKEN")

type doQueryManager struct {
	currRelativeTime types.RelativeTime
	getString        func(t time.Time) (string, error)
	objectType       types.ObjectType
	resultsCount     int

	issueResults []github.Issue
	prResults    []github.PR
}

func (manager *doQueryManager) GetIssues() []github.Issue {
	return manager.issueResults
}

func (manager *doQueryManager) GetLastCreatedIssueTime() *time.Time {
	if len(manager.issueResults) == 0 {
		return nil
	}
	return &(manager.issueResults[len(manager.issueResults)-1]).CreatedAt
}

func (manager *doQueryManager) GetLastCreatedPRTime() *time.Time {
	if len(manager.prResults) == 0 {
		return nil
	}
	return &(manager.prResults[len(manager.prResults)-1]).CreatedAt
}

func (manager *doQueryManager) GetPRs() []github.PR {
	return manager.prResults
}


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
		logger.Info("Found %v/%v issues. Continuing to queryStrings...\n", manager.resultsCount, expectedTotalResults)
		err = manager.paginateDateTime()
		if err != nil {
			return err
		}

		queryResults, remainingTotalResults, err := manager.getResultsFromGitHub()
		if err != nil {
			return err
		}
		if queryResults == nil {
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

//
func (manager *doQueryManager) getResultsFromGitHub() (json.RawMessage, int, error) {
	bodyString, err := manager.getString(manager.currRelativeTime.GetTime())
	if err != nil {
		return nil, 0, fmt.Errorf("error getting string for query: %#v", err)
	}
	body := strings.NewReader(bodyString)

	req, err := http.NewRequest("POST", url, body)
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
		return nil, 0, fmt.Errorf("error unmarchalling response body: %#v", err)
	}

	return githubResults.Data.Search.Results, githubResults.Data.Search.ResultsCount, nil
}


//
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

//
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
