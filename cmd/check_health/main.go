package check_health

import (
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/database"
)

func Main() error {
	var err error

	err = checkDatabase()

	if err != nil {
		return err
	}

	err = checkCsSdk()

	if err != nil {
		return err
	}

	return nil
}

func checkDatabase() error {
	return database.Connect()
}

func checkCsSdk() error {
	var resp map[string]interface{}
	var err error

	// Create a temporary label in Contentstack, to test the management token.
	resp, err = cs_sdk.Request(
		"labels",
		"POST",
		map[string]interface{}{
			"label": map[string]interface{}{
				"name": "__csb_healthcheck",
			},
		},
		true,
	)

	if err != nil {
		return err
	}

	// Delete the temporary label in Contentstack.
	_, err = cs_sdk.Request(
		fmt.Sprintf("labels/%s", resp["label"].(map[string]interface{})["uid"]),
		"DELETE",
		nil,
		true,
	)

	if err != nil {
		return err
	}

	return err
}
