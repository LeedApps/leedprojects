package app_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/leedapps/leedprojects/app"
)

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

var _ = Describe("app", func() {
	Describe("InitializeApp", func() {
		It("Includes the router", func() {
			lpapp := app.InitializeApp(true)
			Expect(fmt.Sprintf("%T", lpapp.Router)).To(Equal("*gin.Engine"))
		})
	})
})
