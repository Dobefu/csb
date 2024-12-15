package cli

import (
	"errors"
	"strings"

	"github.com/Dobefu/csb/cmd/cli/utils"
)

func CreateContentType(name string, machineName string) error {
	ctName, err := confirmName(name)

	if err != nil {
		return err
	}

	if ctName == "" {
		return errors.New("no name provided")
	}

	ctMachineName, err := confirmMachineName(machineName)

	if err != nil {
		return err
	}

	if ctMachineName == "" {
		return errors.New("no machine name provided")
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

		ctName = strings.ReplaceAll(val, "\n", "")
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

		ctName = strings.ReplaceAll(val, "\n", "")
	}

	return ctName, nil
}
