package cs_sdk

import (
	"fmt"
	"os"

	_ "github.com/Dobefu/csb/cmd/init"
)

var URL string
var VERSION string

func init() {
	URL = getUrl()
	VERSION = "v3"
}

func getUrl() string {
	region := os.Getenv("CS_REGION")
	extension := "com"

	region = fmt.Sprintf("%s-", region)

	if region == "us-" {
		region = ""
		extension = "io"
	}

	return fmt.Sprintf("https://%scdn.contentstack.%s", region, extension)
}
