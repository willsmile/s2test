package executor

import (
	"errors"

	"github.com/willsmile/s2test/internal/connector"
	"github.com/willsmile/s2test/internal/reporter"
)

var (
	// ErrNoTasksToExecute is returned when there are no tasks to execute
	ErrNoTasksToExecute = errors.New("there are no tasks to execute")
)

// Plan is a plan that contains test information
type Plan struct {
	Goal        string      `json:"goal"`
	TargetPath  string      `json:"targetPath"`
	AuthMethods authMethods `json:"authMethods"`
	Tasks       []Task      `json:"tasks"`
}

// AuthMethods is a store of prepared information of methods for authentication
type authMethods map[string]map[string]string

// Execute a plan
func (p Plan) Execute(store *connector.Endpoints) (reporter.Report, error) {
	var report reporter.Report

	if len(p.Tasks) == 0 {
		return report, ErrNoTasksToExecute
	}

	for _, task := range p.Tasks {
		entity := task.Perform(store, &p.AuthMethods)
		report = append(report, *entity)
	}

	return report, nil
}

func (p Plan) APIPath(s string) string {
	if s == "" {
		return p.TargetPath
	} else {
		return s
	}
}
