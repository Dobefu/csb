package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func CreateContentType(name string) error {
	ctName, err := confirmName(name)

	if err != nil {
		return err
	}

	if ctName == "" {
		return errors.New("no name provided")
	}

	machineName, err := confirmMachineName(name)

	if err != nil {
		return err
	}

	if machineName == "" {
		return errors.New("no machine name provided")
	}

	return nil
}

func confirmName(name string) (string, error) {
	ctName := name

	if ctName == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter the name of your new content type:")

		key, err := reader.ReadString('\n')

		if err != nil {
			return "", err
		}

		ctName = strings.ReplaceAll(key, "\n", "")
	}

	return ctName, nil
}

func confirmMachineName(name string) (string, error) {
	ctName := name

	if ctName == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter the machine name of your new content type:")

		key, err := reader.ReadString('\n')

		if err != nil {
			return "", err
		}

		ctName = strings.ReplaceAll(key, "\n", "")
	}

	return ctName, nil
}
