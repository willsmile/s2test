package reporter

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

// Print prints out each Report in Reports
func (reports Reports) Print(m printMode) error {
	if len(reports) == 0 {
		return ErrEmptyReport
	}

	for _, v := range reports {
		v.printTarget()
		v.printResult()
		switch m {
		case FullMode:
			v.printResponseState()
			v.printResponseBody()
			v.printRequest()
		case NormalMode:
			v.printResponseState()
			v.printResponseBody()
		case ShortMode:
			v.printResponseState()
		}
	}

	return nil
}

// printTarget prints target of Report
func (report Report) printTarget() {
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, report.target)
}

// printTarget prints result of Report
func (report Report) printResult() {
	var c *color.Color

	if report.result == RequestSent {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	c.Printf("%s Result: %s\n", smallArrow, report.result)
}

// printResponseState prints response state of Report
// when result is not RequestNotSent
func (report Report) printResponseState() {
	if report.result == RequestSent {
		c := color.New(color.FgBlue)
		c.Printf("%s Response state: ", smallArrow)
		fmt.Println(report.response.Status)
	}
}

// printResponseBody prints response body of Report
// when result is not RequestNotSent
func (report Report) printResponseBody() {
	if report.result == RequestSent {
		c := color.New(color.FgBlue)
		c.Printf("%s Response body: ", smallArrow)
		fmt.Println(report.response.Body)
	}
}

// printRequest prints request of Report
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
