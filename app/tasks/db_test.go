package tasks_test

import (
	"github.com/leedapps/leedprojects/app/tasks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DB related tasks", func() {
	Describe("DBCommands", func() {
		It("should have init command", func() {
			dbCommands := tasks.DBCommands(lpapp)
			taskNames := getTaskNames(dbCommands)
			Expect(taskNames[0]).To(Equal("init"))
		})
	})

})
