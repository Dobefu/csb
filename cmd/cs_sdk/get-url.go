package cs_sdk

import (
	"fmt"
	"os"
)

var VERSION = "v3"

func GetUrl() string {
	region := os.Getenv("CS_REGION")
	extension := "com"

	region = fmt.Sprintf("%s-", region)

	if region == "us-" {
		region = ""
		extension = "io"
	}

	return fmt.Sprintf("https://%scdn.contentstack.%s", region, extension)
}
