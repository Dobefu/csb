package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	api_utils "github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

type queryRows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

var queryQueryRows = func(table string, columns []string, where []db_structs.QueryWhere) (queryRows, error) {
	return query.QueryRows(table, columns, where)
}

var apiUtilsGetAltLocales = api_utils.GetAltLocales

func GetSitemapData(w http.ResponseWriter, r *http.Request) {
	sitemapData, err := getEntries()

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	output := utilsConstructOutput()
	output["data"] = sitemapData

	json, err := json.Marshal(output)

	if err != nil {
		utilsPrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}

func getEntries() (map[string]interface{}, error) {
	rows, err := queryQueryRows("routes",
		[]string{"uid", "locale", "url", "updated_at"},
		[]db_structs.QueryWhere{
			{
				Name:  "locale",
				Value: "en",
			},
			{
				Name:     "exclude_sitemap",
				Value:    true,
				Operator: db_structs.NOT_EQUALS,
			},
			{
				Name:     "url",
				Value:    "",
				Operator: db_structs.NOT_EQUALS,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	entries := map[string]interface{}{}

	for rows.Next() {
		var result structs.Route

		err = rows.Scan(
			&result.Uid,
			&result.Locale,
			&result.Url,
			&result.UpdatedAt,
		)

		if err != nil {
			return entries, err
		}

		altLocales, err := apiUtilsGetAltLocales(result, false)

		if err != nil {
			return entries, err
		}

		entries[result.Uid] = struct {
			Uid        string                  `json:"uid"`
			Locale     string                  `json:"locale"`
			Url        string                  `json:"url"`
			UpdatedAt  time.Time               `json:"updated_at"`
			AltLocales []api_structs.AltLocale `json:"alt_locales"`
		}{
			Uid:        result.Uid,
			Locale:     result.Locale,
			Url:        result.Url,
			UpdatedAt:  result.UpdatedAt,
			AltLocales: altLocales,
		}
	}

	return entries, nil
}
