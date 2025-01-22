package api

func GetGlobalFields() (map[string]interface{}, error) {
	data, err := csSdkRequest("global_fields", "GET", nil, false)

	if err != nil {
		return nil, err
	}

	return data, nil
}
