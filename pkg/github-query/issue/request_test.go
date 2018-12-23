package issue

import (
	. "github.com/collector-for-github/pkg/github-query/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query", func() {
	var (
		testQueryAWS = IssuesRequest{
			Terms:      []string{"aws", "eks"},
			Labels:     []string{"sig/aws", "area/platform/aws", "area/platform/eks"},
			SearchIn:   Title,
			State:      Open,
			OwnerLogin: "kubernetes",
			//RepoName: "test-infra",
		}
	)

	// TODO: this test is terrible for obvious reasons and one with fake issues from an achieved org/repo.
	Context("When initialized without any values", func() {
		It("should default with GitHubObjectType of AnyEvent", func() {
			results, err := testQueryAWS.GetIssues()
			Expect(len(results)).To(Equal(161))
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
