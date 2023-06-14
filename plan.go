package main

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
func (p Plan) Execute(store *Endpoints) (Report, error) {
	var report Report

	if len(p.Tasks) == 0 {
		return report, ErrNoTasksToExecute
	}

	for _, task := range p.Tasks {
		entity := task.Perform(store, &p.AuthMethods)
		report = append(report, *entity)
	}

	return report, nil
}
