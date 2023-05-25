package main

import (
	"errors"
	"testing"
)

func TestPlanExecute_WithTasks(t *testing.T) {
	store := Store{
		"GET a sample post": spec{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": spec{
			URL:    "https://jsonplaceholder.typicode.com/todos/1/",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
	}

	plan := Plan{
		Goal:        "For test",
		TargetPath:  "testdata/api.json",
		AuthMethods: map[string]map[string]string{},
		Tasks: []task{
			{
				TargetAPI:  "GET a sample post",
				AuthMethod: "",
			},
			{
				TargetAPI:  "Undefined target",
				AuthMethod: "",
			},
		},
	}

	_, err := plan.Execute(&store)
	if err != nil {
		t.Fatalf("plan.Execute(&store), expected none error, got %s", err)
	}
}

func TestPlanExecute_WithoutTasks(t *testing.T) {
	store := Store{
		"GET a sample post": spec{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": spec{
			URL:    "https://jsonplaceholder.typicode.com/todos/1/",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
	}

	plan := Plan{
		Goal:        "For test",
		TargetPath:  "testdata/api.json",
		AuthMethods: map[string]map[string]string{},
		Tasks:       []task{},
	}

	_, err := plan.Execute(&store)
	if !errors.Is(err, ErrNoTasksToExecute) {
		t.Fatalf("plan.Execute(&store), expected %s, got %s", ErrNoTasksToExecute, err)
	}
}
