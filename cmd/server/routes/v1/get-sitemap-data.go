package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	api_structs "github.com/Dobefu/csb/cmd/api/structs"
	api_utils "github.com/Dobefu/csb/cmd/api/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	"github.com/Dobefu/csb/cmd/database/query"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/server/utils"
)

func GetSitemapData(w http.ResponseWriter, r *http.Request) {
	sitemapData, err := getEntries()

	if err != nil {
		utils.PrintError(w, err, false)
		return
	}

	output := utils.ConstructOutput()
	output["data"] = sitemapData

	json, err := json.Marshal(output)

	if err != nil {
		utils.PrintError(w, err, true)
		return
	}

	fmt.Fprint(w, string(json))
}

func getEntries() (map[string]interface{}, error) {
	rows, err := query.QueryRows("routes",
		[]string{"uid", "locale", "url"},
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
		)

		if err != nil {
			return entries, err
		}

		altLocales, err := api_utils.GetAltLocales(result, false)

		if err != nil {
			return entries, err
		}

		entries[result.Uid] = struct {
			Uid        string                  `json:"uid"`
			Locale     string                  `json:"locale"`
			Url        string                  `json:"url"`
			AltLocales []api_structs.AltLocale `json:"alt_locales"`
		}{
			Uid:        result.Uid,
			Locale:     result.Locale,
			Url:        result.Url,
			AltLocales: altLocales,
		}
	}

	return entries, nil
}
