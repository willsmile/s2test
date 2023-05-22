package main

import (
	"encoding/json"
	"os"
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

// LoadPlan loads a plan from a JSON file
func LoadPlan(path string) (*Plan, error) {
	plan := NewPlan()

	if path == "" {
		return plan, ErrEmptyPath
	}

	src, error := os.ReadFile(path)
	if error != nil {
		return plan, ErrReadFile
	}

	json.Unmarshal(src, &plan)

	return plan, nil
}

// Execute excutes a plan
func (p Plan) Execute(store *Store) Report {
	var report Report

	for _, task := range p.Tasks {
		target := (*store)[task.TargetAPI]
		cookies := p.PreparedCookies[task.UsedCookies]
		resp, _ := HTTPRequest(target.Method, target.URL, target.Headers, cookies)

		entity := reportEntity{
			reqTarget:  task.TargetAPI,
			reqCookies: task.UsedCookies,
			respBody:   resp.Body,
			respStatus: resp.Status,
		}

		report = append(report, entity)
	}

	return report
}
