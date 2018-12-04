package fmtstrings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	whiteSpaces = []string{"   ", " ", "\t", "\n", "\r"}

	oneWord               = []string{"foo"}
	whiteSpacesAndOneWord = append(whiteSpaces, oneWord...)
)

// TESTS: toLowercaseUniqueTrimmedList()
var _ = Describe("get lowercase, trimmed, unique list with toLowercaseUniqueTrimmedList()", func() {

	// Empty list
	Context("With empty list", func() {
		It("should return empty list", func() {
			Expect(ToLowercaseUniqueTrimmedList([]string{})).To(Equal([]string{}))
		})
	})

	// Empty strings or only whitespaces
	Context("With empty strings or only whitespaces in string", func() {
		Context("With only whitespaces", func() {
			It("should return empty list", func() {
				Expect(ToLowercaseUniqueTrimmedList(whiteSpaces)).To(Equal([]string{}))
			})
		})

		Context("With only empty strings", func() {
			It("should return nil", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"", ""})).To(Equal([]string{}))
			})
		})
	})

	Context("With one word", func() {
		Context("Without whitespaces", func() {
			It("should return the list", func() {
				Expect(ToLowercaseUniqueTrimmedList(oneWord)).To(Equal(oneWord))
			})
		})

		Context("With whitespaces", func() {
			It("should trim space and return the list", func() {
				Expect(ToLowercaseUniqueTrimmedList(whiteSpacesAndOneWord)).To(Equal(oneWord))
			})
		})

		Context("With a capitol letter", func() {
			It("should make the word lowercase and return the list", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"HellO"})).To(Equal([]string{"hello"}))
			})
		})
	})

	//
	Context("With two words", func() {
		Context("Without whitespaces", func() {
			It("return the list", func() {
				actual := ToLowercaseUniqueTrimmedList([]string{"hello", "world"})
				Expect(actual).To(Equal([]string{"hello", "world"}))
			})
		})

		Context("With one word containing only whitespaces", func() {
			It("return a list with only the non-whitespace word", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"hello", "       "})).To(Equal([]string{"hello"}))
			})
		})

		Context("With duplicate words", func() {
			It("return a list with only one of the word", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"world", "world"})).To(Equal([]string{"world"}))
			})
		})

		Context("With duplicate words of different cases", func() {
			It("return a list with only one of the word as lowercase", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"worLD", "wORld"})).To(Equal([]string{"world"}))
			})
		})

		Context("With duplicate words with whitespaces", func() {
			It("return a list with only one of the word and without whitespaces", func() {
				Expect(ToLowercaseUniqueTrimmedList([]string{"world   ", " world"})).To(Equal([]string{"world"}))
			})
		})
	})
})
