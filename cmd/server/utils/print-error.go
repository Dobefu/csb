package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func PrintError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"data": null, "error": "%s"}`, strings.ReplaceAll(err.Error(), `"`, `\"`))
}
