package cs_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func RequestRaw(path string, method string, body map[string]interface{}, useManagementToken bool) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", GetUrl(method), VERSION, path)

	bodyJson, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))

	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"api_key": {os.Getenv("CS_API_KEY")},
		"branch":  {os.Getenv("CS_BRANCH")},
	}

	if useManagementToken {
		req.Header.Set("authorization", os.Getenv("CS_MANAGEMENT_TOKEN"))
	} else {
		req.Header.Set("access_token", os.Getenv("CS_DELIVERY_TOKEN"))
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Request(path string, method string, body map[string]interface{}, useManagementToken bool) (map[string]interface{}, error) {
	resp, err := RequestRaw(path, method, body, useManagementToken)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Could not connect to Contentstack: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(respBody, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
