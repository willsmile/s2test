package lib

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "==>"
const smallArrow = " ->"

// Print uses for
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
