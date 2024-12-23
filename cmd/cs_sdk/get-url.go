package cs_sdk

import (
	"fmt"
	"os"
)

var VERSION = "v3"

func GetUrl(useManagementToken bool) string {
	region := os.Getenv("CS_REGION")
	endpoint := "cdn"
	extension := "com"

	region = fmt.Sprintf("%s-", region)

	if region == "us-" {
		region = ""
		extension = "io"
	}

	if useManagementToken {
		endpoint = "api"
	}

	return fmt.Sprintf("https://%s%s.contentstack.%s", region, endpoint, extension)
}
