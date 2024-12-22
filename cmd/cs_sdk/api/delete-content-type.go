package api

import (
	"errors"
	"fmt"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/logger"
)

func DeleteContentType(uid string) error {
	contentType := GetContentType(uid)

	if contentType == nil {
		return errors.New("The content type does not exist")
	}

	_, err := cs_sdk.Request(
		fmt.Sprintf("content_types/%s", uid),
		"DELETE",
		nil,
	)

	if err != nil {
		return err
	}

	logger.Info("The content type has been deleted")

	return nil
}
