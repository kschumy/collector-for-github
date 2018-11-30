package issue

import (
	. "github.com/collector-for-GitHub/pkg/github-query/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TESTS: Query
var _ = Describe("Query", func() {
	var (
		testQueryAWS = IssuesRequest{
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
			_, err := testQueryAWS.GetIssues()
			//Expect(len(results)).To(Equal(163))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
