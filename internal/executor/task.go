package executor

import (
	"github.com/willsmile/s2test/internal/connector"
	"github.com/willsmile/s2test/internal/depository"
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string                    `json:"targetAPI"`
	AuthMethod string                    `json:"authMethod"`
	Data       depository.CustomizedData `json:"data"`
}

// Perform a task
func (t Task) Perform(store *depository.Endpoints, methods *depository.AuthMethods) *reporter.Report {
	endpoint := t.endpoint(store)
	auth := t.authInfo(methods)
	data := t.Data
	req, err := connector.NewRequest(endpoint, auth, data).HTTPRequest()
	if err != nil {
		return t.generateReport(connector.DefaultResponse(), reporter.ResultRequestNotSent)
	}

	client := connector.NewHTTPClient()
	resp, err := connector.SendHTTPRequest(req, client)
	if err != nil {
		return t.generateReport(connector.DefaultResponse(), reporter.ResultRequestError)
	}

	return t.generateReport(resp, reporter.ResultRequestSent)
}

func (t Task) endpoint(store *depository.Endpoints) depository.Endpoint {
	return (*store)[t.TargetAPI]
}

func (t Task) authInfo(methods *depository.AuthMethods) depository.AuthInfo {
	info := (*methods)[t.AuthMethod]
	return depository.NewAuthInfo(info)
}

func (t Task) generateReport(resp *connector.Response, result string) *reporter.Report {
	return &reporter.Report{
		ReqTarget:     t.TargetAPI,
		ReqAuthMethod: t.AuthMethod,
		Result:        result,
		RespBody:      resp.Body,
		RespStatus:    resp.Status,
	}
}
