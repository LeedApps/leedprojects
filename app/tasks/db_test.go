package tasks_test

import (
	"github.com/leedapps/leedprojects/app/tasks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DB related tasks", func() {
	Describe("DBCommands", func() {
		It("should be empty for now", func() {
			Expect(len(tasks.DBCommands())).To(Equal(0))
		})
	})

})
