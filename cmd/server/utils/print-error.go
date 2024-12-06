package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Dobefu/csb/cmd/logger"
)

func PrintError(w http.ResponseWriter, err error, logError bool) {
	fmt.Fprintf(w, `{"data": null, "error": "%s"}`, strings.ReplaceAll(err.Error(), `"`, `\"`))

	if logError {
		logger.Error(err.Error())
	}
}
