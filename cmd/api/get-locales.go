package api

func GetLocales() (map[string]interface{}, error) {
	locales, err := csSdkRequest(
		"locales",
		"GET",
		nil,
		true,
	)

	if err != nil {
		return nil, err
	}

	return locales, nil
}
