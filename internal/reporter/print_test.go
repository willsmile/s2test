package reporter

import (
	"errors"
	"testing"

	myhttp "github.com/willsmile/s2test/internal/http"
)

func TestReportPrint_ValidReport(t *testing.T) {
	report1 := NewReport(
		RequestSent,
		"GET a sample post",
		[]string{},
		&myhttp.Request{},
		&myhttp.Response{},
	)
	report2 := NewReport(
		RequestSent,
		"GET a sample todo",
		[]string{},
		&myhttp.Request{},
		&myhttp.Response{},
	)
	reports := Reports{*report1, *report2}

	mode, _ := NewPrintMode("normal")
	err := reports.Print(mode)
	if err != nil {
		t.Fatalf("reports.Print(), expected none error, got %s", err)
	}
}

func TestReportPrint_EmptyReport(t *testing.T) {
	reports := Reports{}

	mode, _ := NewPrintMode("normal")
	err := reports.Print(mode)
	if !errors.Is(err, ErrEmptyReport) {
		t.Fatalf("reports.Print(), expected %s, got %s", ErrEmptyReport, err)
	}
}
