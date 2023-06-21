package executor

import (
	"testing"

	myhttp "github.com/willsmile/s2test/internal/http"
	"github.com/willsmile/s2test/internal/reporter"
)

func TestTaskPerform_existingTarget(t *testing.T) {
	endpoints := createEndpoints()
	methods := createAuthMethods()
	task := Task{
		TargetAPI:  "GET a sample post",
		AuthMethod: "",
		Variables:  myhttp.Variables{},
	}

	report := task.Perform(endpoints, methods)
	result := report.GetResult()

	if result != reporter.ResultRequestSent {
		t.Fatalf("task.Perform(&endpoints, &methods), expected ResultRequestSent as result, got %s", result)
	}
}

func TestTaskPerform_notExistingTarget(t *testing.T) {
	endpoints := createEndpoints()
	methods := createAuthMethods()
	task := Task{
		TargetAPI:  "Not Existing Target",
		AuthMethod: "",
		Variables:  myhttp.Variables{},
	}

	report := task.Perform(endpoints, methods)
	result := report.GetResult()

	if result != reporter.ResultRequestNotSent {
		t.Fatalf("task.Perform(&endpoints, &methods), expected ResultRequestNotSent as result, got %s", result)
	}
}

func createEndpoints() *myhttp.Endpoints {
	return &myhttp.Endpoints{
		"GET a sample post": myhttp.Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": myhttp.Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/todos/1/",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
	}
}

func createAuthMethods() *myhttp.AuthMethods {
	return &myhttp.AuthMethods{
		"cookieA": {
			"type":  "Cookie",
			"name":  "cookieName",
			"value": "cookieValue",
		},
		"tokenA": {
			"type":   "OAuth 2.0",
			"prefix": "Bearer",
			"value":  "tokenValue",
		},
	}
}
