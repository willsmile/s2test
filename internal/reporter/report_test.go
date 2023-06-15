package reporter

import (
	"errors"
	"testing"
)

func TestReportPrint_ValidReport(t *testing.T) {
	reports := Reports{
		Report{
			ReqTarget:     "GET a sample post",
			ReqAuthMethod: "",
			Result:        "SENT",
			RespBody:      "200 OK",
			RespStatus:    "{}",
		},
		Report{
			ReqTarget:     "GET a sample todo",
			ReqAuthMethod: "",
			Result:        "SENT",
			RespBody:      "200 OK",
			RespStatus:    "{}",
		},
	}

	err := reports.Print()
	if err != nil {
		t.Fatalf("reports.Print(), expected none error, got %s", err)
	}
}

func TestReportPrint_EmptyReport(t *testing.T) {
	reports := Reports{}

	err := reports.Print()
	if !errors.Is(err, ErrEmptyReport) {
		t.Fatalf("reports.Print(), expected %s, got %s", ErrEmptyReport, err)
	}
}
