package main

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "==>"
const smallArrow = " ->"

// reportEntity uses for providing a struct for a report of executed task
type reportEntity struct {
	reqTarget  string
	reqCookies string
	respBody   string
	respStatus string
}

// Report uses for providing a struct for slice of reportEntity
type Report []reportEntity

// Print uses for printing out each reportEntity in report
func (report Report) Print() {
	for _, v := range report {
		printTarget(v.reqTarget)
		printResponse(v.respBody, v.respStatus)
	}
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
