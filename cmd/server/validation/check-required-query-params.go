package validation

import (
	"fmt"
	"net/http"
	"strings"
)

func CheckRequiredQueryParams(r *http.Request, params ...string) (map[string]any, error) {
	values := make(map[string]any)
	var missingParams []string

	query := r.URL.Query()

	for _, param := range params {
		paramValue := query[param]

		if len(paramValue) == 0 {
			missingParams = append(missingParams, param)
			continue
		}

		values[param] = paramValue
	}

	if len(missingParams) > 0 {
		return nil, fmt.Errorf(
			"missing required query params: (%s)", strings.Join(missingParams, ", "),
		)
	}

	fmt.Println(missingParams)

	return values, nil
}
