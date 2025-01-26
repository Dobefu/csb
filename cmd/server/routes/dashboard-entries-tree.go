package routes

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/Dobefu/csb/cmd/logger"
	jwt "github.com/golang-jwt/jwt/v5"
)

//go:embed templates/*
var content embed.FS

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

	payload, err := jwtParse(token, func(*jwt.Token) (interface{}, error) {
		publicKey, err := jwtParseRSAPublicKeyFromPEM([]byte(key))

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
	url := strings.Replace(csSdkGetUrl(true), "api", "app", 1)
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
		fmt.Println("TEST")
		return "", err
	}

	key, hasKey := data["signing-key"]

	if !hasKey {
		return "", errors.New("missing signing key")
	}

	return key.(string), nil
}
