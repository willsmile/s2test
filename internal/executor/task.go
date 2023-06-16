package executor

import (
	"github.com/willsmile/s2test/internal/connector"
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string                   `json:"targetAPI"`
	AuthMethod string                   `json:"authMethod"`
	Data       connector.CustomizedData `json:"data"`
}

// Perform a task
func (t Task) Perform(store *connector.Endpoints, methods *authMethods) *reporter.Report {
	endpoint := t.endpoint(store)
	auth := t.authInfo(methods)
	data := t.Data
	req, err := connector.NewRequest(endpoint, auth, data).HTTPRequest()
	if err != nil {
		return t.createReport(connector.DefaultResponse(), reporter.ResultRequestNotSent)
	}

	client := connector.NewHTTPClient()
	resp, err := connector.SendHTTPRequest(req, client)
	if err != nil {
		return t.createReport(connector.DefaultResponse(), reporter.ResultRequestError)
	}

	return t.createReport(resp, reporter.ResultRequestSent)
}

func (t Task) endpoint(store *connector.Endpoints) connector.Endpoint {
	return (*store)[t.TargetAPI]
}

func (t Task) authInfo(methods *authMethods) connector.AuthInfo {
	info := (*methods)[t.AuthMethod]
	return connector.NewAuthInfo(info)
}

func (t Task) createReport(resp *connector.Response, result string) *reporter.Report {
	return &reporter.Report{
		ReqTarget:     t.TargetAPI,
		ReqAuthMethod: t.AuthMethod,
		Result:        result,
		RespBody:      resp.Body,
		RespStatus:    resp.Status,
	}
}