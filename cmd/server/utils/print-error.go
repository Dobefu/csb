package utils

import (
	"fmt"
	"net/http"
)

func PrintError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, `{"data": null, "error": "%s"}`, err.Error())
}
