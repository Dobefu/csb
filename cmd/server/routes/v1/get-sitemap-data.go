package v1

import (
	"fmt"
	"net/http"
)

func GetSitemapData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}
