package api

import (
	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/logger"
)

func CreateOrUpdateTranslationsContentType() error {
	body := getTranslationsContentTypeBody()
	_ = CreateContentType("Translations", "translations", false)

	_, err := cs_sdk.Request(
		"content_types/translations",
		"PUT",
		body,
		true,
	)

	if err != nil {
		return err
	}

	logger.Info("The content type has been updated")

	return nil
}

func getTranslationsContentTypeBody() map[string]interface{} {
	return map[string]interface{}{
		"content_type": map[string]interface{}{
			"title":       "Translations",
			"uid":         "translations",
			"description": "Static translations for elements within the frontend application",
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
			},
			"options": map[string]interface{}{
				"title":       "title",
				"is_page":     false,
				"publishable": true,
				"singleton":   false,
			},
		},
	}
}
