package main

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "==>"
const smallArrow = " ->"

// Report is a slice of report entities
type Report []reportEntity

// reportEntity is a report entity of each executed task
type reportEntity struct {
	reqTarget     string
	reqAuthMethod string
	respBody      string
	respStatus    string
}

// Print prints out each reportEntity in Report
func (report Report) Print() {
	for _, v := range report {
		printTarget(v.reqTarget)
		printResponse(v.respBody, v.respStatus)
	}
}

// printTarget prints a label of target API
func printTarget(target string) {
	c := color.New(color.FgGreen, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, target)
}

// printResponse prints body and state of repsonse
func printResponse(body string, state string) {
	c := color.New(color.FgBlue)

	c.Printf("%s Response state: ", smallArrow)
	fmt.Println(state)
	c.Printf("%s Response body: ", smallArrow)
	fmt.Println(body)
}
