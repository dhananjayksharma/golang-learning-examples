package urls_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Urls", func() {
	Describe("GetShortURL", func() {
		Context("when result are success", func() {
			It("get short url - success", func() {
				inputData := gin.H{
					"url": "https://news.google.com/topstories?tab=mn&hl=en-IN&gl=IN&ceid=IN:en",
				}
				var c *gin.Context
				_ = inputData
				c.Bod
				out, _ := GetShortURL(c)

				Expect(out.data.ShortURL).To(Equal("https://bit.ly/3Bg26nL"))
			})
		})
	})
})

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
