package reporter

import (
	"errors"

	myhttp "github.com/willsmile/s2test/internal/http"
)

const (
	arrow      = "==>"
	smallArrow = " ->"
)

var (
	// ErrEmptyReport is returned when reports is empty
	ErrEmptyReport = errors.New("nothing to print due to empty reports")
)

// Reports is a slice of Report
type Reports []Report

// Report records the results on the execution of each task
type Report struct {
	result   result
	target   string
	auth     string
	request  *myhttp.Request
	response *myhttp.Response
}

func NewReport(result result, target string, auth string, req *myhttp.Request, resp *myhttp.Response) *Report {
	return &Report{
		result:   result,
		target:   target,
		auth:     auth,
		request:  req,
		response: resp,
	}
}
