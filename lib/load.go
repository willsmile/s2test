package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Plan uses for providing a struct for test plan information
type Plan struct {
	Goal            string                       `json:"goal"`
	TargetPath      string                       `json:"targetPath"`
	PreparedCookies map[string]map[string]string `json:"preparedcookies"`
	Tasks           []Task                       `json:"tasks"`
}

// Task uses for providing a struct for task definition
type Task struct {
	TargetAPI   string `json:"targetAPI"`
	UsedCookies string `json:"usedCookies"`
}

// APIStore uses for providing a struct for storing API information
type APIStore map[string]API

// API uses for providing a struct for information of a single API
type API struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// LoadPlan uses for loading a test plan from a JSON file
func LoadPlan(path string) Plan {
	var plan Plan

	if path == "" {
		log.Fatal("[Invaild Input Error] empty argument of path")
		os.Exit(1)
	}

	raw, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("[File Loading Error] ", error)
		os.Exit(1)
	}

	json.Unmarshal(raw, &plan)

	return plan
}

// LoadAPIStore uses for loading APIStore from a JSON file
func LoadAPIStore(path string) APIStore {
	var store APIStore

	if path == "" {
		log.Fatal("[Invaild Input Error] empty argument of path")
		os.Exit(1)
	}

	raw, error := ioutil.ReadFile(path)
	if error != nil {
		log.Fatal("[File Loading Error] ", error)
		os.Exit(1)
	}

	json.Unmarshal(raw, &store)

	return store
}
