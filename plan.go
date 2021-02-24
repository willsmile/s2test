package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Plan is a plan that contains test information
type Plan struct {
	Goal            string                       `json:"goal"`
	TargetPath      string                       `json:"targetPath"`
	PreparedCookies map[string]map[string]string `json:"preparedcookies"`
	Tasks           []task                       `json:"tasks"`
}

// task is a task definition for test
type task struct {
	TargetAPI   string `json:"targetAPI"`
	UsedCookies string `json:"usedCookies"`
}

// NewPlan constructs an empty plan.
func NewPlan() *Plan {
	return &Plan{}
}

// LoadPlan loads a plan from a JSON file
func LoadPlan(path string) *Plan {
	p := NewPlan()

	if path == "" {
		log.Fatal("[Invaild Input Error] empty argument of path")
		os.Exit(1)
	}

	raw, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("[File Loading Error] ", error)
		os.Exit(1)
	}

	json.Unmarshal(raw, &p)

	return p
}

// Execute excutes a plan
func (p Plan) Execute(store *Store) Report {
	var report Report

	for _, task := range p.Tasks {
		target := (*store)[task.TargetAPI]
		cookies := p.PreparedCookies[task.UsedCookies]
		body, status := HTTPRequest(target.Method, target.URL, target.Headers, cookies)

		entity := reportEntity{
			reqTarget:  task.TargetAPI,
			reqCookies: task.UsedCookies,
			respBody:   body,
			respStatus: status,
		}

		report = append(report, entity)
	}

	return report
}
