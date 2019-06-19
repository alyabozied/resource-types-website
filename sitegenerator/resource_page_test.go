package sitegenerator_test

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	. "github.com/concourse/dutyfree/matchers"
	"github.com/concourse/dutyfree/sitegenerator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourcePage", func() {
	It("renders the template", func() {
		resourceModel := sitegenerator.ResourceModel{
			Resource: sitegenerator.Resource{
				Name:       "git resource",
				Repository: "https://github.com/concourse/git-resource",
			},
			Identifier:        "concourse-git-resource",
			AuthorHandle:      "concourse",
			AuthorProfileLink: "https://github.com/concourse",
			Readme:            "<div>foobar readme</div>",
		}

		b := bytes.Buffer{}

		ip := sitegenerator.NewResourcePage("", resourceModel)
		err := ip.Generate(&b)

		Expect(err).ToNot(HaveOccurred())

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b.Bytes()))

		Expect(err).ToNot(HaveOccurred())

		Expect(doc).To(
			SatisfyAll(
				ContainSelectorWithText("h1", Equal("git resource")),
				ContainSelectorWithText("p", Equal("https://github.com/concourse/git-resource")),
				ContainSelectorWithText("#github-readme > div", Equal("foobar readme"))),
		)
	})
})
