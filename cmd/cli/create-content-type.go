package cli

import (
	"errors"

	"github.com/Dobefu/csb/cmd/cli/utils"
	"github.com/Dobefu/csb/cmd/cs_sdk/api"
)

func CreateContentType(isDryRun bool, name string, id string) error {
	ctName, err := confirmName(name)

	if err != nil {
		return err
	}

	if ctName == "" {
		return errors.New("no name provided")
	}

	ctId, err := confirmMachineName(id)

	if err != nil {
		return err
	}

	if ctId == "" {
		return errors.New("no machine name provided")
	}

	if isDryRun {
		return nil
	}

	err = api.CreateContentType(ctName, ctId)

	if err != nil {
		return err
	}

	return nil
}

func confirmName(name string) (string, error) {
	ctName := name

	if ctName == "" {
		val, err := utils.ReadLine("Please enter the name of your new content type")

		if err != nil {
			return "", err
		}

		ctName = val
	}

	return ctName, nil
}

func confirmMachineName(name string) (string, error) {
	ctName := name

	if ctName == "" {
		val, err := utils.ReadLine("Please enter the machine name of your new content type")

		if err != nil {
			return "", err
		}

		ctName = val
	}

	return ctName, nil
}
