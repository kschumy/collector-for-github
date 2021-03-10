// package github_issue_generator generates random issues on GitHub
package github_issue_generator

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

const (
	MaxAllowed    = 150 // Must be >0 and should be >= 150. Anything more than this risk getting rate-restricted by GitHub.
	SecondsToWait = 4   // Number of seconds between calls. Due to GitHub rate limits, cannot be less than 2.
)

type Seed struct {
	UserRepo   string // where issues will be added. Format must be: <org/owner's username>/<UserRepo name>"
	TermOne    string // term to be included randomly in some issue title and/or body
	TermTwo    string // another term to be included randomly in some issue title and/or body
	ForceTerms bool   // if 'true', TermOne and TermTwo will be included in all titles and bodies
	// TODO: implement EnsureUnique
	EnsureUnique bool // if 'true', will only seed if TermOne and TermTwo do not

	LabelOne string // a label to be randomly added to some issues if IncludeLabels
	LabelTwo string // another label to be randomly added to some issues if IncludeLabels

	Quantity int // Ending value for issues. Should be greater than 1 and MUST be greater than StartNum.

	StartNum   int  // Starting value for issues. Should be greater than 0 and MUST be less than EndNum.
	IncludeNum bool // if 'true', includes number (starting at StartNum and incrementing) in each issue title
}

type seedValues struct {
	termOne    string // term to be included randomly in some issue title and/or body
	termTwo    string // another term to be included randomly in some issue title and/or body
	forceTerms bool   // if 'true', TermOne and TermTwo will be included in all titles and bodies

	labelOne string // a label to be randomly added to some issues if IncludeLabels
	labelTwo string // another label to be randomly added to some issues if IncludeLabels

	termOneTexts   stringList
	termTwoTexts   stringList
	bothTermsTexts stringList
}

var gitHubToken = os.Getenv("GITHUB_GRAPHQL_TOKEN") // must be a valid token with write access to UserRepo

// Creates new issue(s) for UserRepo on GitHub. Number of issues created should be equal to EndNum - StartNum (see note
// below) and will include TermOne and/or TermTwo in some randomly-selected issue titles and bodies, as well as including
// LabelOne and/or LabelTwo on randomly-selected issue if IncludeLabels is set to true.
// NOTE: Once the program makes somewhere between 200-300 calls, GitHub will likely flag the token holder's account for
// abuse of service and prohibit any additional API calls for a period of datetime. It is highly recommended that any
// high-volume use of this program are broken up into 200 call batches.
func (seed *Seed) SeedIssues() error {
	seedInfo, err := seed.createSeedIssues()
	if err != nil {
		return err
	}
	
	sleepDuration := SecondsToWait * time.Second
	
	for i := seed.StartNum; i < seed.StartNum+seed.Quantity; i++ {
		reqString := seedInfo.getReqString(i, seed.IncludeNum)
		logger.Debug("%v: %#v", i, reqString)
		err := generateIssue(reqString, seed.UserRepo)
		if err != nil {
			return err
		}
		time.Sleep(sleepDuration) // GitHub rate limits prohibit more than 30 calls/minute
	}
	logger.Info("Seeded %v issues to %s", seed.Quantity, seed.UserRepo)
	return nil
}

// Creates and returns a string needed for an API call body to created an issue with a title, body, and possibly labels,
// which is determined by if IncludeLabels is true and if the program has randomly determined to create a label(s).
func (sv *seedValues) getReqString(i int, includeNum bool) string {
	reqString := strings.Builder{}
	reqString.WriteString("{\n  \"title\":\"")
	if includeNum {
		reqString.WriteString(fmt.Sprintf("#%v: ", i))
	}
	reqString.WriteString(fmt.Sprintf("%s\",\n \"body\":\"%s\"", sv.getRandomTitle(), sv.getRandomBody()))
	label := sv.getRandomLabel()
	if label != "" {
		reqString.WriteString(fmt.Sprintf(",\n \"labels\": [\n    \"%s\"\n  ]", label))
	}
	reqString.WriteString("\n}")

	return reqString.String()
}

// Creates a new issue on UserRepo on GitHub.
func generateIssue(reqString, userRepo string) error {
	reqBody := strings.NewReader(reqString)
	url := "https://api.github.com/repos/" + userRepo + "/issues"

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return fmt.Errorf("ERROR with Request: %v", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Authorization", "bearer "+gitHubToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ERROR with call: %v", err)
	}
	defer res.Body.Close()
	if res.Status != "201 Created" {
		return fmt.Errorf("ERROR with status: %v", res.Status)
	}

	fmt.Println(res)
	return nil
}

// Returns a string intended to be used an issue body.
// The returned string will contain both contain both TermOne and TermTwo if ForceTerms is true. Otherwise, the text may
// include one term, both terms, or neither term.
func (sv *seedValues) getRandomBody() string {
	if sv.forceTerms {
		return sv.bothTermsTexts.getRandomString()
	}

	randNum := getRandomNumBasedOnNumOfTerms(sv)

	switch randNum {
	case 0:
		return noTermsTexts.getRandomString()
	case 1:
		return sv.termOneTexts.getRandomString()
	case 2:
		return sv.termTwoTexts.getRandomString()
	default:
		return sv.termOneTexts.getRandomString()
	}
}

// Returns a string intended to be used an issue Title.
// The returned string will contain both contain both TermOne and TermTwo if ForceTerms is true. Otherwise, the text may
// include one term, both terms, or neither term.
func (sv *seedValues) getRandomTitle() string {
	titleWordOne := phrases[getRandomNum(len(phrases))]
	titleWordTwo := phrases[getRandomNum(len(phrases))]

	var insertString string
	if sv.forceTerms {
		if sv.onlyOneTerm() {
			insertString = sv.termOne
		} else {
			insertString = sv.termOne + " " + sv.termOne
		}
	} else {
		
		randNum := getRandomNumBasedOnNumOfTerms(sv)
		
		switch randNum {
		case 0:
			return titleWordOne + " " + titleWordTwo
		case 1:
			insertString = sv.termOne
		case 2:
			insertString = sv.termOne + " " + sv.termTwo
		default:
			insertString = sv.termTwo
		}
	}

	switch getRandomNum(3) {
	case 0:
		return insertString + " " + titleWordOne + " " + titleWordTwo
	case 1:
		return titleWordOne + " " + insertString + " " + titleWordTwo
	default:
		return titleWordOne + " " + titleWordTwo + " " + insertString
	}
}

// Returns either an empty string or a string comprised of LabelOne and/ or LabelTwo
func (sv *seedValues) getRandomLabel() string {
	if !sv.hasAtLeastOneLabel() {
		return ""
	}

	if sv.labelOne == "" {
		return sv.labelTwo
	} else if sv.labelTwo == "" {
		return sv.labelOne
	}

	switch getRandomNum(4) {
	case 0:
		return ""
	case 1:
		return sv.labelOne + "\", \"" + sv.labelTwo
	case 2:
		return sv.labelOne
	default:
		return sv.labelTwo
	}
}

// Returns a random int between 0 (inclusive) and n (exclusive).
// n must be greater than 0
func getRandomNum(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func getRandomNumBasedOnNumOfTerms(sv *seedValues) int {
	if sv.onlyOneTerm() {
		return getRandomNum(2)
	}
	return getRandomNum(4)
}

func (seed *Seed) createSeedIssues() (*seedValues, error) {
	seedInfo := &seedValues{}

	errors := []string{}
	if seed.Quantity < 1 {
		errors = append(errors, fmt.Sprintf("quatity %v cannot be less than 1", seed.Quantity))
	}
	// Seeding more than MaxAllowed at a time runs the risk of breaking GitHub's rate limit
	if seed.Quantity > MaxAllowed {
		errors = append(errors, fmt.Sprintf("quatity %v cannot be greater than %v", seed.Quantity, MaxAllowed))
	}
	termOne := strings.TrimSpace(seed.TermOne)
	if len(termOne) == 0 {
		errors = append(errors, fmt.Sprintf("term one %v must not be empty or only have white spaces", termOne))
	}
	// TODO: additional validation check for UserRepo
	if seed.UserRepo > "" {
		errors = append(errors, fmt.Sprintf("repo %s cannot be empty", seed.UserRepo))
	}
	if len(errors) > 0 {
		return seedInfo, fmt.Errorf("errors initializing seed: %s", strings.Join(errors, ", "))
	}

	seedInfo.termOne = termOne
	seedInfo.termTwo = strings.TrimSpace(seed.TermTwo)
	seedInfo.labelOne = strings.TrimSpace(seed.LabelOne)
	seedInfo.labelTwo = strings.TrimSpace(seed.LabelTwo)
	seedInfo.forceTerms = seed.ForceTerms
	seedInfo.termOneTexts = *getTermOneTexts(seedInfo.termOne)
	seedInfo.termTwoTexts = *getTermTwoTexts(seedInfo.termTwo)
	seedInfo.bothTermsTexts = *getBothTermsTexts(seedInfo.termOne, seedInfo.termTwo)

	return seedInfo, nil
}

func (sv *seedValues) onlyOneTerm() bool {
	return sv.termTwo == ""
}

func (sv *seedValues) hasAtLeastOneLabel() bool {
	return sv.labelOne != "" || sv.labelTwo != ""
}
