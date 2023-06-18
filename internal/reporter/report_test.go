package reporter

import (
	"errors"
	"testing"

	myhttp "github.com/willsmile/s2test/internal/http"
)

func TestReportPrint_ValidReport(t *testing.T) {
	report1 := NewReport(
		ResultRequestSent,
		"GET a sample post",
		"",
		&myhttp.Request{},
		&myhttp.Response{},
	)
	report2 := NewReport(
		ResultRequestSent,
		"GET a sample todo",
		"",
		&myhttp.Request{},
		&myhttp.Response{},
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
