package api

func GetContentTypes() (map[string]interface{}, error) {
	data, err := csSdkRequest("content_types", "GET", nil, false)

	if err != nil {
		return nil, err
	}

	return data, nil
}
