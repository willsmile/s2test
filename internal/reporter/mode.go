package reporter

import "errors"

type printMode int

const (
	UnsupportedMode printMode = iota
	NormalMode
	FullMode
	ShortMode
)

var (
	// ErrUnsupportedPrintMode is returned when parameter value of print mode is not supported.
	ErrUnsupportedPrintMode = errors.New("print mode is unsupported")
)

func NewPrintMode(s string) (printMode, error) {
	switch s {
	case "normal":
		return NormalMode, nil
	case "full":
		return FullMode, nil
	case "short":
		return ShortMode, nil
	default:
		return UnsupportedMode, ErrUnsupportedPrintMode
	}
}
