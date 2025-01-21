package remote_setup

import "github.com/Dobefu/csb/cmd/cs_sdk/api"

var apiCreateOrUpdateSeoGlobalField = api.CreateOrUpdateSeoGlobalField
var apiCreateOrUpdateTranslationsContentType = api.CreateOrUpdateTranslationsContentType

func Main() error {
	var err error

	err = apiCreateOrUpdateSeoGlobalField()

	if err != nil {
		return err
	}

	err = apiCreateOrUpdateTranslationsContentType()

	if err != nil {
		return err
	}

	return nil
}
