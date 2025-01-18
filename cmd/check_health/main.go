package check_health

import (
	"errors"
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/database"
)

var databaseConnect = database.Connect
var csSdkRequest = cs_sdk.Request

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
	err := databaseConnect()

	if err != nil {
		return err
	}

	return nil
}

func checkCsSdk() error {
	var resp map[string]interface{}
	var err error

	// Create a temporary label in Contentstack, to test the management token.
	resp, err = csSdkRequest(
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

	label, hasLabel := resp["label"]

	if !hasLabel {
		return errors.New("response has no label")
	}

	// Delete the temporary label in Contentstack.
	_, err = csSdkRequest(
		fmt.Sprintf("labels/%s", label.(map[string]interface{})["uid"]),
		"DELETE",
		nil,
		true,
	)

	if err != nil {
		return err
	}

	return err
}
