package main

import (
	"flag"
	"fmt"

	"github.com/willsmile/s2test/lib"
)

func main() {
	var path string

	flag.StringVar(&path, "p", "", "please specify path of test plan to execute")
	flag.Parse()

	plan := lib.LoadPlan(path)
	store := lib.LoadAPIStore(plan.TargetPath)

	reports := lib.Execute(plan, store)

	// Debug
	fmt.Println("[Reports]")
	fmt.Println(reports)
}
