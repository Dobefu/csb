package check_health

import (
	"fmt"
	"net/http"

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
	var resp *http.Response
	var err error

	resp, err = cs_sdk.RequestRaw("content_types", "GET")

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Could not connect to Contentstack: %s", resp.Status)
	}

	return err
}
