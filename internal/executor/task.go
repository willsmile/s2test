package executor

import (
	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string           `json:"targetAPI"`
	AuthMethod string           `json:"authMethod"`
	Variables  myhttp.Variables `json:"variables"`
}

// Perform a task
func (t Task) Perform(endpoints *myhttp.Endpoints, methods *myhttp.AuthMethods) *reporter.Report {
	request := myhttp.NewRequest(
		endpoints.GetEndpoint(t.TargetAPI),
		methods.GetAuthInfo(t.AuthMethod),
		t.Variables,
	)

	req, err := request.HTTPRequest()
	if err != nil {
		return t.generateReport(
			reporter.ResultRequestNotSent,
			request,
			myhttp.DefaultResponse(),
		)
	}

	client := myhttp.NewClient()
	resp, err := myhttp.Send(req, client)
	if err != nil {
		return t.generateReport(
			reporter.ResultRequestError,
			request,
			myhttp.DefaultResponse(),
		)
	}

	return t.generateReport(reporter.ResultRequestSent, request, resp)
}

func (t Task) generateReport(result string, req *myhttp.Request, resp *myhttp.Response) *reporter.Report {
	return reporter.NewReport(
		result,
		t.TargetAPI,
		t.AuthMethod,
		req,
		resp,
	)
}
