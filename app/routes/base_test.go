package routes_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/leedapps/leedprojects/app/routes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Base", func() {
	Describe("LPRouter", func() {
		It("has 'ping' route", func() {
			router := routes.LPRouter()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ping", nil)
			router.ServeHTTP(w, req)
			fmt.Println(w.Code)
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("pong"))
		})
	})
})
