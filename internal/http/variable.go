package http

import (
	"fmt"
	"strings"
)

type Variables map[string]string

func (vbs Variables) newReplacer() *strings.Replacer {
	var oldnew []string
	for key, value := range vbs {
		oldnew = append(oldnew, wrap(key), value)
	}

	return strings.NewReplacer(oldnew...)
}

func wrap(k string) string {
	return fmt.Sprintf("#{%s}", k)
}
