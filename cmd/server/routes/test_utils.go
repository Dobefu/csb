package routes

import (
	"io/fs"
	"net/http"

	"github.com/Dobefu/csb/cmd/cs_sdk"
	"github.com/Dobefu/csb/cmd/server/utils"
	jwt "github.com/golang-jwt/jwt/v5"
)

type FS interface {
	fs.FS
	ReadDir(string) ([]fs.DirEntry, error)
	ReadFile(string) ([]byte, error)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var utilsConstructOutput = utils.ConstructOutput
var httpClient HttpClient = &http.Client{}
var csSdkGetUrl = cs_sdk.GetUrl
var jwtParseRSAPublicKeyFromPEM = jwt.ParseRSAPublicKeyFromPEM
var getFs = func() FS { return content }
