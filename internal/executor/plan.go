package executor

import (
	"errors"

	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
	"github.com/willsmile/s2test/internal/storage"
)

var (
	// ErrNoTasksToExecute is returned when there are no tasks to execute
	ErrNoTasksToExecute = errors.New("there are no tasks to execute")
)

// Plan is a plan that contains test information
type Plan struct {
	Goal              string             `json:"goal"`
	EndpointsFilepath string             `json:"endpoints"`
	UserAgent         string             `json:"ua"`
	AuthDataset       myhttp.AuthDataset `json:"auths"`
	Tasks             []Task             `json:"tasks"`
}

// Execute a plan
func (p Plan) Execute(store *storage.Endpoints, info string) (reporter.Reports, error) {
	var reports reporter.Reports

	if len(p.Tasks) == 0 {
		return reports, ErrNoTasksToExecute
	}

	ua := p.GetUserAgent(info)
	for _, task := range p.Tasks {
		entity := task.Perform(store, &p.AuthDataset, ua)
		reports = append(reports, *entity)
	}

	return reports, nil
}

func (p Plan) GetEndpointsPath(s string) string {
	if s == "" {
		return p.EndpointsFilepath
	} else {
		return s
	}
}

func (p Plan) GetUserAgent(appInfo string) string {
	if p.UserAgent != "" {
		return p.UserAgent
	}

	return appInfo
}
