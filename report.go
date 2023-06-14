package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	arrow      = "==>"
	smallArrow = " ->"
)

// Report is a slice of report entities
type Report []reportEntity

// reportEntity is a report entity of each executed task
type reportEntity struct {
	reqTarget     string
	reqAuthMethod string
	result        string
	respBody      string
	respStatus    string
}

func NewReportEntity(t *Task, resp *Response, result string) *reportEntity {
	return &reportEntity{
		reqTarget:     t.TargetAPI,
		reqAuthMethod: t.AuthMethod,
		result:        result,
		respBody:      resp.Body,
		respStatus:    resp.Status,
	}
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
func (entity reportEntity) printTarget() {
	c := color.New(color.FgYellow, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, entity.reqTarget)
}

// printTarget prints result in reportEntity
func (entity reportEntity) printResult() {
	var c *color.Color

	if entity.result == ResultRequestSent {
		c = color.New(color.FgGreen)
	} else {
		c = color.New(color.FgRed)
	}

	c.Printf("%s Result: %s\n", smallArrow, entity.result)
}

// printResponse prints respBody and respStatus in reportEntity
// when result is not RequestNotSent
func (entity reportEntity) printResponse() {
	if entity.result == ResultRequestSent {
		c := color.New(color.FgBlue)
		c.Printf("%s Response state: ", smallArrow)
		fmt.Println(entity.respStatus)
		c.Printf("%s Response body: ", smallArrow)
		fmt.Println(entity.respBody)
	}
}
