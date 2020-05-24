package main

import (
	"flag"

	"github.com/willsmile/s2test/lib"
)

func main() {
	var path string

	flag.StringVar(&path, "p", "", "please specify path of test plan to execute")
	flag.Parse()

	plan := lib.LoadPlan(path)
	store := lib.LoadAPIStore(plan.TargetPath)

	lib.Execute(plan, store)
}
