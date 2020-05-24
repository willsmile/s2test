package main

import (
	"flag"

	"github.com/willsmile/s2test/lib"
)

// Version of s2test
const version = "0.1.0"

func main() {
	var path string

	flag.StringVar(&path, "p", "", "please specify path of test plan to execute")
	flag.Parse()

	plan := lib.LoadPlan(path)
	store := lib.LoadAPIStore(plan.TargetPath)

	lib.Execute(plan, store)
}
