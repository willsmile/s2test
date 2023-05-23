package main

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
func (p Plan) Execute(store *Store) Report {
	var report Report

	for _, task := range p.Tasks {
		target := (*store)[task.TargetAPI]
		authMethod := p.AuthMethods[task.AuthMethod]
		authInfo := NewAuthInfo(authMethod)
		resp, _ := HTTPRequest(target.Method, target.URL, target.Headers, authInfo)

		entity := reportEntity{
			reqTarget:     task.TargetAPI,
			reqAuthMethod: task.AuthMethod,
			respBody:      resp.Body,
			respStatus:    resp.Status,
		}

		report = append(report, entity)
	}

	return report
}
