package api

import (
	"errors"
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/logger"
)

func DeleteGlobalField(uid string, isForced bool) error {
	globalField := GetGlobalField(uid)

	if globalField == nil {
		return errors.New("The global field does not exist")
	}

	_, err := cs_sdk.Request(
		fmt.Sprintf("global_fields/%s?force=%t", uid, isForced),
		"DELETE",
		nil,
	)

	if err != nil {
		return err
	}

	logger.Info("The global field has been deleted")

	return nil
}
