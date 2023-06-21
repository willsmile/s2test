package executor

import (
	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI string           `json:"targetAPI"`
	Auth      string           `json:"auth"`
	Variables myhttp.Variables `json:"variables"`
}

// Perform a task
func (t Task) Perform(endpoints *myhttp.Endpoints, dataset *myhttp.AuthDataset) *reporter.Report {
	request := myhttp.NewRequest(
		endpoints.GetEndpoint(t.TargetAPI),
		dataset.Select(t.Auth).NewAuthInfo(),
		t.Variables,
	)

	req, err := request.HTTPRequest()
	if err != nil {
		return t.reportNotSent(request)
	}

	client := myhttp.NewClient()
	resp, err := myhttp.Send(req, client)
	if err != nil {
		return t.reportError(request)
	}

	return t.reportSent(request, resp)
}

func (t Task) reportSent(req *myhttp.Request, resp *myhttp.Response) *reporter.Report {
	return reporter.NewReport(
		reporter.ResultRequestSent,
		t.TargetAPI,
		t.Auth,
		req,
		resp,
	)
}

func (t Task) reportNotSent(req *myhttp.Request) *reporter.Report {
	return reporter.NewReport(
		reporter.ResultRequestNotSent,
		t.TargetAPI,
		t.Auth,
		req,
		myhttp.DefaultResponse(),
	)
}

func (t Task) reportError(req *myhttp.Request) *reporter.Report {
	return reporter.NewReport(
		reporter.ResultRequestError,
		t.TargetAPI,
		t.Auth,
		req,
		myhttp.DefaultResponse(),
	)
}
