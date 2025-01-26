package routes

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"text/template"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/logger"
	jwt "github.com/golang-jwt/jwt/v5"
)

type FS interface {
	fs.FS
	ReadDir(string) ([]fs.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

//go:embed templates/*
var content embed.FS
var getFs = func() FS { return content }

var httpClient HttpClient = &http.Client{}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func DashboardEntriesTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("X-Content-Type-Options", "nosniff")

	_, err := validateToken(r)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	data := make(map[string]interface{})

	templates := []string{
		"templates/dashboard-entries-tree.html.tmpl",
	}

	tpl, err := (template.ParseFS(getFs(), templates...))

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	tpl.Option("missingkey=error")
	err = tpl.Execute(&buf, data)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	output := buf.Bytes()
	fmt.Fprint(w, string(output))
}

func validateToken(r *http.Request) (*jwt.Token, error) {
	token := r.URL.Query().Get("app-token")

	if token == "" {
		return nil, errors.New("missing app token")
	}

	key, err := getPublicKey()

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	payload, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))

		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if !payload.Valid {
		return nil, errors.New("invalid payload")
	}

	return payload, nil
}

func getPublicKey() (string, error) {
	url := strings.Replace(cs_sdk.GetUrl(true), "api", "app", 1)
	url = fmt.Sprintf("%s/.well-known/public-keys.json", url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var data map[string]interface{}

	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return "", err
	}

	key, hasKey := data["signing-key"]

	if !hasKey {
		return "", errors.New("missing signing key")
	}

	return key.(string), nil
}
