package reporter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	myhttp "github.com/willsmile/s2test/internal/http"
)

const (
	ResultRequestSent    = "SENT"
	ResultRequestNotSent = "NOT SENT"
	ResultRequestError   = "ERROR"
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
	result   string
	target   string
	auth     string
	request  *myhttp.Request
	response *myhttp.Response
}

func NewReport(result string, target string, auth string, req *myhttp.Request, resp *myhttp.Response) *Report {
	return &Report{
		result:   result,
		target:   target,
		auth:     auth,
		request:  req,
		response: resp,
	}
}

// Print prints out each Report in Reports
func (reports Reports) Print() error {
	if len(reports) == 0 {
		return ErrEmptyReport
	}

	for _, v := range reports {
		v.printTarget()
		v.printResult()
		v.printResponse()
		v.printRequest()
	}

	return nil
}

func (report Report) GetResult() string {
	return report.result
}

// printTarget prints ReqTarget of Report
func (report Report) printTarget() {
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, report.target)
}

// printTarget prints Result of Report
func (report Report) printResult() {
	var c *color.Color

	if report.result == ResultRequestSent {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	c.Printf("%s Result: %s\n", smallArrow, report.result)
}

// printResponse prints Response field of Report
// when result is not RequestNotSent
func (report Report) printResponse() {
	if report.result == ResultRequestSent {
		c := color.New(color.FgBlue)
		c.Printf("%s Response state: ", smallArrow)
		fmt.Println(report.response.Status)
		c.Printf("%s Response body: ", smallArrow)
		fmt.Println(report.response.Body)
	}
}

// printRequest prints Request field of Report
func (report Report) printRequest() {
	c := color.New(color.FgCyan)
	c.Printf("%s Request URL: ", smallArrow)
	fmt.Println(report.request.URL)

	c.Printf("%s Request Method: ", smallArrow)
	fmt.Println(report.request.Method)

	c.Printf("%s Request Headers: ", smallArrow)
	printHeaders(report.request.Headers)

	c.Printf("%s Request Cookies: ", smallArrow)
	printCookies(report.request.Cookies)

	c.Printf("%s Request Body: ", smallArrow)
	printBody(report.request.Body)
}

func printHeaders(headers http.Header) {
	if len(headers) != 0 {
		fmt.Printf("\n")
		for k, v := range headers {
			fmt.Println("   ", k, ":", v)
		}
	} else {
		fmt.Println("Empty")
	}
}

func printCookies(cookies []*http.Cookie) {
	if len(cookies) != 0 {
		fmt.Printf("\n")
		for _, c := range cookies {
			fmt.Println("   ", c)
		}
	} else {
		fmt.Println("Empty")
	}
}

func printBody(body string) {
	if body == "" {
		fmt.Println("Empty")
	} else {
		fmt.Println(body)
	}
}
