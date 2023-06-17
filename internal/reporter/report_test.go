package reporter

import (
	"errors"
	"testing"
)

func TestReportPrint_ValidReport(t *testing.T) {
	report1 := NewReport(
		ResultRequestSent,
		"GET a sample post",
		"",
		"200 OK",
		"{}",
	)
	report2 := NewReport(
		ResultRequestSent,
		"GET a sample todo",
		"",
		"200 OK",
		"{}",
	)
	reports := Reports{*report1, *report2}

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
