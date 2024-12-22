package api

import "github.com/Dobefu/csb/cmd/cs_sdk"

func CreateOrUpdateSeoGlobalField() error {
	body := getSeoGlobalFieldBody()
	err := CreateGlobalField("seo", body)

	if err != nil {
		_, err := cs_sdk.Request(
			"global_fields/seo",
			"PUT",
			body,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func getSeoGlobalFieldBody() map[string]interface{} {
	return map[string]interface{}{
		"global_field": map[string]interface{}{
			"title":       "SEO",
			"uid":         "seo",
			"description": "SEO metadata for a routable page",
			"schema": []map[string]interface{}{
				{
					"display_name": "Page Title",
					"uid":          "title",
					"data_type":    "text",
					"field_metadata": map[string]interface{}{
						"_default":    true,
						"instruction": "The page title that will be displayed in the browser tab.",
					},
					"unique":    false,
					"mandatory": false,
					"multiple":  false,
				},
				{
					"display_name": "Description",
					"uid":          "description",
					"data_type":    "text",
					"field_metadata": map[string]interface{}{
						"instruction": "A short description that search engines will see. For the best results, please keep the description between 150 to 160 characters in length.",
						"multiline":   true,
					},
					"unique":    false,
					"mandatory": false,
					"multiple":  false,
				},
			},
		},
	}
}
