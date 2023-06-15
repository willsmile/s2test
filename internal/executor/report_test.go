package executor

import (
	"errors"
	"testing"
)

func TestReportPrint_ValidReport(t *testing.T) {
	report := Report{
		reportEntity{
			reqTarget:     "GET a sample post",
			reqAuthMethod: "",
			result:        "SENT",
			respBody:      "200 OK",
			respStatus:    "{}",
		},
		reportEntity{
			reqTarget:     "GET a sample todo",
			reqAuthMethod: "",
			result:        "SENT",
			respBody:      "200 OK",
			respStatus:    "{}",
		},
	}

	err := report.Print()
	if err != nil {
		t.Fatalf("report.Print(), expected none error, got %s", err)
	}
}

func TestReportPrint_EmptyReport(t *testing.T) {
	report := Report{}

	err := report.Print()
	if !errors.Is(err, ErrEmptyReport) {
		t.Fatalf("report.Print(), expected %s, got %s", ErrEmptyReport, err)
	}
}
