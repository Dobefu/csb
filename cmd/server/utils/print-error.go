package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Dobefu/csb/cmd/logger"
)

func PrintError(w http.ResponseWriter, err error, isError bool) {
	statusCode := http.StatusBadRequest
	fmt.Fprintf(w, `{"data": null, "error": "%s"}`, strings.ReplaceAll(err.Error(), `"`, `\"`))

	if isError {
		statusCode = http.StatusInternalServerError
		logger.Error(err.Error())
	}

	w.WriteHeader(statusCode)
}
