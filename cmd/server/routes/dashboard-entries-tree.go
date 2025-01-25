package routes

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"text/template"

	"github.com/Dobefu/csb/cmd/logger"
)

type FS interface {
	fs.FS
	ReadDir(string) ([]fs.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

//go:embed templates/*
var content embed.FS
var getFs = func() FS { return content }

func DashboardEntriesTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("X-Content-Type-Options", "nosniff")

	data := make(map[string]interface{})

	templates := []string{
		"templates/dashboard-entries-tree.html.tmpl",
	}

	tpl := template.Must(template.ParseFS(getFs(), templates...))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)

	if err != nil {
		logger.Error(err.Error())
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

	output := buf.Bytes()
	fmt.Fprint(w, string(output))
}
