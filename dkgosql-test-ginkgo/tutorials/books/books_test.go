package books_test

import (
	"log"

	"dkgosql.com/tutorials-abc/books"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book.IsReligious()", func() {
	BeforeEach(func() {
		log.Println("Its religious table test")
	})

	Context("when the book is a religious book", func() {
		It("returns true", func() {
			book := books.Book{Name: "Ramayan", Religious: true}
			response := book.IsReligious()
			Expect(response).To(BeTrue())
		})
	})

	Context("when the book is a not religious book", func() {

		It("returns false", func() {
			book := books.Book{Name: "Champak", Religious: false}
			response := book.IsReligious()
			Expect(response).To(BeFalse())
		})
	})

	DescribeTable("is religious table test",
		func(name string, r bool, response bool) {
			b := books.Book{Name: name, Religious: r}
			Expect(b.IsReligious()).To(Equal(response))
		},
		Entry("when book is religious", "Mahabharat", true, true),
		Entry("when book is not religious", "Chhotta Bheem", false, false),
	)
})
