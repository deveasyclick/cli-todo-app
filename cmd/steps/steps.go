// Package steps provides utility for creating
// each step of the CLI
package steps

// A StepSchema contains the data that is used
// for an individual step of the CLI
type StepSchema struct {
	StepName string   // The name of a given step
	Options  []string // The slice of each option for a given step
	Headers  string   // The title displayed at the top of a given step
	Field    string
}

// Steps contains a slice of steps
type Steps struct {
	Steps map[string]StepSchema
}

// An Item contains the data for each option
// in a StepSchema.Options
type Item struct {
	Flag, Title, Desc string
}

// InitSteps initializes and returns the *Steps to be used in the CLI program
func InitSteps() *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"actions": {
				StepName: "TODO Actions",
				Options: []string{
					"Add a new todo",
					"View all your todos",
				},
				Headers: "Select an action to continue or q to quit.",
			},
		},
	}

	return steps
}
