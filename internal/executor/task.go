package executor

import (
	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string                `json:"targetAPI"`
	AuthMethod string                `json:"authMethod"`
	Data       myhttp.CustomizedData `json:"data"`
}

// Perform a task
func (t Task) Perform(store *myhttp.Endpoints, methods *myhttp.AuthMethods) *reporter.Report {
	endpoint := store.Endpoint(t.TargetAPI)
	auth := methods.AuthInfo(t.AuthMethod)
	data := t.Data
	req, err := myhttp.NewRequest(endpoint, auth, data).HTTPRequest()
	if err != nil {
		return t.generateReport(myhttp.DefaultResponse(), reporter.ResultRequestNotSent)
	}

	client := myhttp.NewHTTPClient()
	resp, err := myhttp.SendHTTPRequest(req, client)
	if err != nil {
		return t.generateReport(myhttp.DefaultResponse(), reporter.ResultRequestError)
	}

	return t.generateReport(resp, reporter.ResultRequestSent)
}

func (t Task) generateReport(resp *myhttp.Response, result string) *reporter.Report {
	return reporter.NewReport(
		result,
		t.TargetAPI,
		t.AuthMethod,
		resp.Body,
		resp.Status,
	)
}
