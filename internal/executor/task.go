package executor

const (
	ResultRequestSent    = "SENT"
	ResultRequestNotSent = "NOT SENT"
	ResultRequestError   = "ERROR"
)

// Task is a task definition for test
type Task struct {
	TargetAPI  string         `json:"targetAPI"`
	AuthMethod string         `json:"authMethod"`
	Data       CustomizedData `json:"data"`
}

// Perform a task
func (t Task) Perform(store *Endpoints, methods *authMethods) *reportEntity {
	endpoint := t.endpoint(store)
	auth := t.authInfo(methods)
	data := t.Data
	req, err := NewRequest(endpoint, auth, data).HTTPRequest()
	if err != nil {
		return NewReportEntity(&t, DefaultResponse(), ResultRequestNotSent)
	}

	client := NewHTTPClient()
	resp, err := SendHTTPRequest(req, client)
	if err != nil {
		return NewReportEntity(&t, DefaultResponse(), ResultRequestError)
	}

	return NewReportEntity(&t, resp, ResultRequestSent)
}

func (t Task) endpoint(store *Endpoints) Endpoint {
	return (*store)[t.TargetAPI]
}

func (t Task) authInfo(methods *authMethods) AuthInfo {
	info := (*methods)[t.AuthMethod]
	return NewAuthInfo(info)
}
