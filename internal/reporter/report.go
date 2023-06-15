package reporter

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
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
	// ErrEmptyReport is returned when report is empty
	ErrEmptyReport = errors.New("nothing to print due to empty report")
)

// Report is a slice of report entities
type Report []ReportEntity

// ReportEntity is a report entity of each executed task
type ReportEntity struct {
	ReqTarget     string
	ReqAuthMethod string
	Result        string
	RespBody      string
	RespStatus    string
}

// Print prints out each reportEntity in Report
func (report Report) Print() error {
	if len(report) == 0 {
		return ErrEmptyReport
	}

	for _, v := range report {
		v.printTarget()
		v.printResult()
		v.printResponse()
	}

	return nil
}

// printTarget prints reqTarget in reportEntity
func (entity ReportEntity) printTarget() {
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, entity.ReqTarget)
}

// printTarget prints result in reportEntity
func (entity ReportEntity) printResult() {
	var c *color.Color

	if entity.Result == ResultRequestSent {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	c.Printf("%s Result: %s\n", smallArrow, entity.Result)
}

// printResponse prints respBody and respStatus in reportEntity
// when result is not RequestNotSent
func (entity ReportEntity) printResponse() {
	if entity.Result == ResultRequestSent {
		c := color.New(color.FgBlue)
		c.Printf("%s Response state: ", smallArrow)
		fmt.Println(entity.RespStatus)
		c.Printf("%s Response body: ", smallArrow)
		fmt.Println(entity.RespBody)
	}
}
