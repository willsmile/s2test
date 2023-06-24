package executor

import (
	"errors"
	"testing"

	myhttp "github.com/willsmile/s2test/internal/http"
)

func TestPlanExecute_WithTasks(t *testing.T) {
	store := myhttp.Endpoints{
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

	plan := Plan{
		Goal:              "For test",
		EndpointsFilepath: "testdata/api.json",
		AuthDataset:       myhttp.AuthDataset{},
		Tasks: []Task{
			{
				TargetAPI: "GET a sample post",
				Auth:      "",
			},
			{
				TargetAPI: "Undefined target",
				Auth:      "",
			},
		},
	}

	appInfo := "test"
	_, err := plan.Execute(&store, appInfo)
	if err != nil {
		t.Fatalf("plan.Execute(&store, appInfo), expected none error, got %s", err)
	}
}

func TestPlanExecute_WithoutTasks(t *testing.T) {
	store := myhttp.Endpoints{
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

	plan := Plan{
		Goal:              "For test",
		EndpointsFilepath: "testdata/api.json",
		AuthDataset:       myhttp.AuthDataset{},
		Tasks:             []Task{},
	}

	appInfo := "test"
	_, err := plan.Execute(&store, appInfo)
	if !errors.Is(err, ErrNoTasksToExecute) {
		t.Fatalf("plan.Execute(&store, appInfo), expected %s, got %s", ErrNoTasksToExecute, err)
	}
}
