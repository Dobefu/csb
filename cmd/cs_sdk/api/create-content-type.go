package api

import (
	"errors"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/logger"
)

func CreateContentType(name string, id string, withFields bool) error {
	contentType := GetContentType(id)

	if contentType != nil {
		return errors.New("The content type already exists")
	}

	body := map[string]interface{}{
		"content_type": map[string]interface{}{
			"title": name,
			"uid":   id,
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
				},
			},
			"options": map[string]interface{}{
				"title":       "title",
				"publishable": true,
				"singleton":   false,
			},
		},
	}

	if withFields {
		err := CreateOrUpdateSeoGlobalField()

		if err != nil {
			return err
		}

		body["content_type"].(map[string]interface{})["schema"] = []map[string]interface{}{
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
			{
				"display_name": "SEO",
				"uid":          "seo",
				"data_type":    "global_field",
				"reference_to": "seo",
				"field_metadata": map[string]interface{}{
					"_default":    true,
					"instruction": "Additional SEO (Search Engine Optimization) fields for the page.",
				},
				"unique":   false,
				"multiple": false,
			},
		}

		body["content_type"].(map[string]interface{})["options"] = map[string]interface{}{
			"title":       "title",
			"publishable": true,
			"is_page":     true,
			"singleton":   false,
			"sub_title":   []string{"url"},
			"url_pattern": "/:title",
			"url_prefix":  "/",
		}

	}

	_, err := cs_sdk.Request(
		"content_types",
		"POST",
		body,
		true,
	)

	if err != nil {
		return err
	}

	logger.Info("The content type has been created")

	return nil
}
