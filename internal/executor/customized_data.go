package executor

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

	for key, value := range ctd {
		wrappedKey := wrap(key)
		result = strings.Replace(result, wrappedKey, value, 1)
	}

	return result
}

func wrap(k string) string {
	return fmt.Sprintf("#{%s}", k)
}
