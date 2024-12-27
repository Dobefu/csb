package remote_setup

import "github.com/Dobefu/csb/cmd/cs_sdk/api"

func Main() error {
	var err error

	err = api.CreateOrUpdateSeoGlobalField()

	if err != nil {
		return err
	}

	err = api.CreateOrUpdateTranslationsContentType()

	if err != nil {
		return err
	}

	return nil
}
