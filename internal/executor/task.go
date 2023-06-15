package executor

import (
	"github.com/willsmile/s2test/internal/reporter"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string         `json:"targetAPI"`
	AuthMethod string         `json:"authMethod"`
	Data       CustomizedData `json:"data"`
}

// Perform a task
func (t Task) Perform(store *Endpoints, methods *authMethods) *reporter.ReportEntity {
	endpoint := t.endpoint(store)
	auth := t.authInfo(methods)
	data := t.Data
	req, err := NewRequest(endpoint, auth, data).HTTPRequest()
	if err != nil {
		return t.reportEntity(DefaultResponse(), reporter.ResultRequestNotSent)
	}

	client := NewHTTPClient()
	resp, err := SendHTTPRequest(req, client)
	if err != nil {
		return t.reportEntity(DefaultResponse(), reporter.ResultRequestError)
	}

	return t.reportEntity(resp, reporter.ResultRequestSent)
}

func (t Task) endpoint(store *Endpoints) Endpoint {
	return (*store)[t.TargetAPI]
}

func (t Task) authInfo(methods *authMethods) AuthInfo {
	info := (*methods)[t.AuthMethod]
	return NewAuthInfo(info)
}

func (t Task) reportEntity(resp *Response, result string) *reporter.ReportEntity {
	return &reporter.ReportEntity{
		ReqTarget:     t.TargetAPI,
		ReqAuthMethod: t.AuthMethod,
		Result:        result,
		RespBody:      resp.Body,
		RespStatus:    resp.Status,
	}
}
