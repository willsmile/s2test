package executor

import (
	"errors"

	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
)

var (
	// ErrNoTasksToExecute is returned when there are no tasks to execute
	ErrNoTasksToExecute = errors.New("there are no tasks to execute")
)

// Plan is a plan that contains test information
type Plan struct {
	Goal        string             `json:"goal"`
	TargetPath  string             `json:"endpoints"`
	AuthDataset myhttp.AuthDataset `json:"auths"`
	Tasks       []Task             `json:"tasks"`
}

// Execute a plan
func (p Plan) Execute(store *myhttp.Endpoints) (reporter.Reports, error) {
	var reports reporter.Reports

	if len(p.Tasks) == 0 {
		return reports, ErrNoTasksToExecute
	}

	for _, task := range p.Tasks {
		entity := task.Perform(store, &p.AuthDataset)
		reports = append(reports, *entity)
	}

	return reports, nil
}

func (p Plan) APIPath(s string) string {
	if s == "" {
		return p.TargetPath
	} else {
		return s
	}
}
