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
			true,
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
				{
					"display_name": "OG Title",
					"uid":          "og_title",
					"data_type":    "text",
					"field_metadata": map[string]interface{}{
						"_default":    true,
						"instruction": "The page title that will be displayed when sharing the page to social media. For the best results, please keep the title between 55 to 60 characters in length.",
					},
					"unique":    false,
					"mandatory": false,
					"multiple":  false,
				},
				{
					"display_name": "OG Description",
					"uid":          "og_description",
					"data_type":    "text",
					"field_metadata": map[string]interface{}{
						"instruction": "A short description that will be displayed when sharing the page to social media. For the best results, please keep the description between 150 to 160 characters in length.",
						"multiline":   true,
					},
					"unique":    false,
					"mandatory": false,
					"multiple":  false,
				},
				{
					"display_name": "Exclude from the sitemap",
					"uid":          "exclude_sitemap",
					"data_type":    "boolean",
					"field_metadata": map[string]interface{}{
						"instruction":   "Whether or not to exclude this page from the sitemap. When checked, the page will not show up in the sitemap.",
						"default_value": false,
					},
					"unique":    false,
					"mandatory": false,
					"multiple":  false,
				},
			},
		},
	}
}
