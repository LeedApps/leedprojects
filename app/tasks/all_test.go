package tasks_test

import (
	"github.com/leedapps/leedprojects/app"
	"github.com/leedapps/leedprojects/app/tasks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/cli/v2"
)

func getTaskNames(commands []*cli.Command) []string {
	names := make([]string, len(commands))
	for index, command := range commands {
		names[index] = command.Name
	}
	return names
}

var lpapp app.LPApp = app.InitializeApp(false)

var _ = Describe("all tasks", func() {
	Describe("AllCommands", func() {
		It("Consolidates all commands", func() {
			allCommands := tasks.AllTasks(lpapp)
			commandNames := getTaskNames(allCommands)
			Expect(commandNames).To(ConsistOf("db"))
			Expect(len(allCommands)).To(Equal(1))
		})
	})
})
