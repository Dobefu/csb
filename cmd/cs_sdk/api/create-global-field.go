package api

import (
	"errors"

	"github.com/Dobefu/csb/cmd/logger"
)

func CreateGlobalField(id string, data map[string]interface{}) error {
	globalField := GetGlobalField(id)

	if globalField != nil {
		return errors.New("The global field already exists")
	}

	_, err := csSdkRequest(
		"global_fields",
		"POST",
		data,
		true,
	)

	if err != nil {
		return err
	}

	logger.Info("The global field has been created")

	return nil
}
