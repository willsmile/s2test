package lib

// Report uses for providing a struct for a report of executed task
type Report struct {
	reqTarget  string
	reqCookies string
	respBody   string
	respStatus string
}

// Execute uses for execute a test plan
func Execute(plan Plan, store APIStore) {

	for _, task := range plan.Tasks {
		target := store[task.TargetAPI]
		cookies := plan.PreparedCookies[task.UsedCookies]
		body, status := HTTPRequest(target.Method, target.URL, target.Headers, cookies)

		report := Report{
			reqTarget:  task.TargetAPI,
			reqCookies: task.UsedCookies,
			respBody:   body,
			respStatus: status,
		}
		report.Print()
	}
}
