package querystrings

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("format datetime", func() {
	var testTimeUTC = time.Date(2018, 2, 5, 9, 0, 0, 0, time.UTC)
	var testTimeUTCString = "2018-02-05T09:00:00Z"

	Context("When provided datetime in UTC", func() {
		It("returns the datetime as a string in the correct format", func() {
			actual := formatTime(&testTimeUTC)
			Expect(actual).To(Equal(testTimeUTCString))
		})
	})

	Context("When provided datetime is behind UTC", func() {
		var testTimePST time.Time
		var timeLocation, _ = time.LoadLocation("America/Los_Angeles")

		BeforeEach(func() {
			testTimePST = time.Date(2018, 2, 5, 9, 0, 0, 0, timeLocation)
		})

		It("returns the UTC-converted datetime as a string in the correct format", func() {
			actual := formatTime(&testTimePST)
			Expect(actual).To(Equal(testTimeUTCString))
		})

		It("does not modify the original datetime", func() {
			originalTime := testTimePST
			formatTime(&testTimePST)
			Expect(originalTime).To(Equal(testTimePST))
			Expect(timeLocation).To(Equal(testTimePST.Location()))
		})
	})

	Context("When provided datetime is ahead of UTC", func() {
		var timeLocation, _ = time.LoadLocation("Asia/Tokyo")
		var testTimeJST = time.Date(2018, 2, 5, 9, 0, 0, 0, timeLocation)

		It("returns the UTC-converted datetime as a string in the correct format", func() {
			actual := formatTime(&testTimeJST)
			Expect(actual).To(Equal(testTimeUTCString))
		})
	})
})
