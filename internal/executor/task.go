package executor

import (
	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
	"github.com/willsmile/s2test/internal/storage"
)

// Task is a task definition for test
type Task struct {
	TargetAPI string           `json:"targetAPI"`
	Auth      []string         `json:"auth"`
	Variables myhttp.Variables `json:"variables"`
}

// Perform a task
func (t Task) Perform(endpoints *storage.Endpoints, dataset *myhttp.AuthDataset, ua string) *reporter.Report {
	request := myhttp.NewRequest(
		endpoints.GetEndpoint(t.TargetAPI),
		dataset.Select(t.Auth).NewAuthInfo(),
		t.Variables,
		ua,
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
		reporter.RequestSent,
		t.TargetAPI,
		t.Auth,
		req,
		resp,
	)
}

func (t Task) reportNotSent(req *myhttp.Request) *reporter.Report {
	return reporter.NewReport(
		reporter.RequestNotSent,
		t.TargetAPI,
		t.Auth,
		req,
		myhttp.DefaultResponse(),
	)
}

func (t Task) reportError(req *myhttp.Request) *reporter.Report {
	return reporter.NewReport(
		reporter.RequestError,
		t.TargetAPI,
		t.Auth,
		req,
		myhttp.DefaultResponse(),
	)
}
