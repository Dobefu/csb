package routes

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/logger"
	jwt "github.com/golang-jwt/jwt/v5"
)

//go:embed templates/*
var content embed.FS

func DashboardEntriesTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("X-Content-Type-Options", "nosniff")

	jwtToken, err := validateToken(r)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if jwtToken.Claims == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := getData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
		return "", err
	}

	key, hasKey := data["signing-key"]

	if !hasKey {
		return "", errors.New("missing signing key")
	}

	return key.(string), nil
}

func getData() (map[string]interface{}, error) {
	data := make(map[string]interface{})

	entries, err := getEntries()

	if err != nil {
		return nil, err
	}

	nestedEntries := getNestedEntries(entries)

	data["Entries"] = nestedEntries

	return data, nil
}

func getNestedEntries(entries map[string]interface{}) map[string]interface{} {
	entryMap := make(map[string]map[string]interface{})

	for uid, entry := range entries {
		entry.(map[string]interface{})["children"] = []interface{}{}
		entryMap[uid] = entry.(map[string]interface{})
	}

	for _, entry := range entryMap {
		parentID := entry["parent"].(string)

		if parentID != "" {
			parentEntry, hasParentEntry := entryMap[parentID]

			if hasParentEntry {
				parentEntry["children"] = append(parentEntry["children"].([]interface{}), entry)
			}
		}
	}

	nestedEntries := make(map[string]interface{})

	for uid, entry := range entryMap {
		parentID := entry["parent"].(string)

		if parentID == "" {
			nestedEntries[uid] = entry
		}
	}

	return nestedEntries
}

func getEntries() (map[string]interface{}, error) {
	rows, err := queryQueryRows("routes", []string{"uid", "parent", "title"}, []structs.QueryWhere{
		{
			Name:  "locale",
			Value: "en",
		},
	})

	if err != nil {
		return nil, err
	}

	entries := make(map[string]interface{})

	for rows.Next() {
		var uid string
		var parent string
		var title string

		err := rows.Scan(
			&uid,
			&parent,
			&title,
		)

		if err != nil {
			continue
		}

		entries[uid] = map[string]interface{}{
			"uid":    uid,
			"title":  title,
			"parent": parent,
		}
	}

	return entries, nil
}
