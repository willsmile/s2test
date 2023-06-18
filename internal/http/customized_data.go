package http

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CustomizedData map[string]string

func (ctd CustomizedData) Apply(raw json.RawMessage) string {
	result := string(raw)
	if len(ctd) == 0 {
		return result
	}
	replacer := ctd.newReplacer()

	return replacer.Replace(result)
}

func (ctd CustomizedData) newReplacer() *strings.Replacer {
	var oldnew []string
	for key, value := range ctd {
		oldnew = append(oldnew, wrap(key), value)
	}

	return strings.NewReplacer(oldnew...)
}

func wrap(k string) string {
	return fmt.Sprintf("#{%s}", k)
}
