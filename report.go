package main

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "==>"
const smallArrow = " ->"

// Report uses for providing a struct for a report of executed task
type Report struct {
	reqTarget  string
	reqCookies string
	respBody   string
	respStatus string
}

// Print uses for print target, body and status
func (report Report) Print() {
	printTarget(report.reqTarget)
	printResponse(report.respBody, report.respStatus)
}

// printTarget uses for representing label of target API
func printTarget(target string) {
	c := color.New(color.FgGreen, color.Bold)
	c.Printf("%s Target API: %s\n", arrow, target)
}

// printResponse uses for representing body and state of repsonse
func printResponse(body string, state string) {
	c := color.New(color.FgBlue)

	c.Printf("%s Response state: ", smallArrow)
	fmt.Println(state)
	c.Printf("%s Response body: ", smallArrow)
	fmt.Println(body)
}
