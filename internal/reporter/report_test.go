package reporter

import (
	"errors"
	"testing"
)

func TestReportPrint_ValidReport(t *testing.T) {
	report := Report{
		ReportEntity{
			ReqTarget:     "GET a sample post",
			ReqAuthMethod: "",
			Result:        "SENT",
			RespBody:      "200 OK",
			RespStatus:    "{}",
		},
		ReportEntity{
			ReqTarget:     "GET a sample todo",
			ReqAuthMethod: "",
			Result:        "SENT",
			RespBody:      "200 OK",
			RespStatus:    "{}",
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
