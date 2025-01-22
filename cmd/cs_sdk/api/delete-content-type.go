package api

import (
	"errors"
	"fmt"

	"github.com/Dobefu/csb/cmd/logger"
)

func DeleteContentType(uid string, isForced bool) error {
	contentType := GetContentType(uid)

	if contentType == nil {
		return errors.New("the content type does not exist")
	}

	_, err := csSdkRequest(
		fmt.Sprintf("content_types/%s?force=%t", uid, isForced),
		"DELETE",
		nil,
		true,
	)

	if err != nil {
		return err
	}

	logger.Info("The content type has been deleted")

	return nil
}
