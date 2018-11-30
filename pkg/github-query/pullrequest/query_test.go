package issue

import (
	"github.com/collector-for-GitHub/pkg/github-query/issue"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/collector-for-GitHub/pkg/github-query/types"

)

// TESTS: Query
var _ = Describe("Query", func() {
	var (
		testQueryAWS = issue.IssuesRequest{
			Terms:      []string{"aws", "eks"},
			Labels:     []string{"sig/aws", "area/platform/aws", "area/platform/eks"},
			SearchIn:   Title,
			State:      Open,
			OwnerLogin: "kubernetes",
			//RepoName: "foobar-repo",
		}
	)

	// TESTS: basic gh
	Context("When initialized without any values", func() {
		It("should default with GitHubObjectType of AnyEvent", func() {
			results, err := testQueryAWS.GetIssues()
			Expect(len(results)).To(Equal(163))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
