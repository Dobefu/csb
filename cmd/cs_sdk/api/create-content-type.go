package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk"
)

func CreateContentType(name string, machineName string) error {
	_, err := cs_sdk.Request(
		"content_types",
		"POST",
		map[string]interface{}{
			"content_type": map[string]interface{}{
				"title": name,
				"uid":   machineName,
				"schema": []map[string]interface{}{
					{
						"display_name": "Title",
						"uid":          "title",
						"data_type":    "text",
						"field_metadata": map[string]interface{}{
							"_default": true,
						},
						"unique":    true,
						"mandatory": true,
						"multiple":  false,
					},
					{
						"display_name": "URL",
						"uid":          "url",
						"data_type":    "text",
						"field_metadata": map[string]interface{}{
							"_default": true,
						},
						"unique":   false,
						"multiple": false,
					},
				},
				"options": map[string]interface{}{
					"title":       "title",
					"publishable": true,
					"is_page":     true,
					"singleton":   false,
					"sub_title": []string{
						"url",
					},
					"url_pattern": "/:title",
					"url_prefix":  "/",
				},
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
