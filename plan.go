package main

const (
	RequestSent    = "SENT"
	RequestNotSent = "NOT SENT"
)

// Plan is a plan that contains test information
type Plan struct {
	Goal        string                       `json:"goal"`
	TargetPath  string                       `json:"targetPath"`
	AuthMethods map[string]map[string]string `json:"authMethods"`
	Tasks       []task                       `json:"tasks"`
}

// task is a task definition for test
type task struct {
	TargetAPI  string `json:"targetAPI"`
	AuthMethod string `json:"authMethod"`
}

// Execute excutes a plan
func (p Plan) Execute(store *Store) (Report, error) {
	var report Report

	if len(p.Tasks) == 0 {
		return report, ErrNoTasksToExecute
	}

	for _, task := range p.Tasks {
		resp := DefaultResponse()
		result := RequestNotSent

		target, err := (*store).Search(task.TargetAPI)
		authMethod := p.AuthMethods[task.AuthMethod]
		authInfo := NewAuthInfo(authMethod)
		if err == nil {
			resp, _ = HTTPRequest(target.Method, target.URL, target.Headers, authInfo)
			result = RequestSent
		}

		entity := reportEntity{
			reqTarget:     task.TargetAPI,
			reqAuthMethod: task.AuthMethod,
			result:        result,
			respBody:      resp.Body,
			respStatus:    resp.Status,
		}

		report = append(report, entity)
	}

	return report, nil
}
