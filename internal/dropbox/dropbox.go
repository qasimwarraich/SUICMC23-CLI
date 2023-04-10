package dropbox

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/file_properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func Upload() {
	var path string

	if os.Getenv("AWS_EXECUTION_ENV") != "" {
		path = "/tmp/suicmc23-data/"
	} else {
		path = "suicmc23-data/"
	}

	token := refreshOAuthToken()
	if token == "" {
		token = os.Getenv("DBX_TOKEN")
	}

	config := dropbox.Config{
		Token:    token,
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}

	filesClient := files.New(config)

	dataDir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, dataFiles := range dataDir {

		content, err := os.ReadFile(path + dataFiles.Name())
		if err != nil {
			log.Fatal(err)
		}

		arg := files.UploadArg{
			CommitInfo: files.CommitInfo{
				Path: "/" + dataFiles.Name(),
				Mode: &files.WriteMode{
					Tagged: dropbox.Tagged{
						Tag: "overwrite",
					},
					Update: "",
				},
				Autorename:     false,
				ClientModified: &time.Time{},
				Mute:           false,
				PropertyGroups: []*file_properties.PropertyGroup{},
				StrictConflict: false,
			},
			ContentHash: "",
		}

		_, err = filesClient.Upload(&arg, strings.NewReader(string(content)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func refreshOAuthToken() string {
	type AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	appKey := os.Getenv("DBX_APP_KEY")
	appSecret := os.Getenv("DBX_APP_SECRET")
	refreshToken := os.Getenv("DBX_REFRESH_TOKEN")

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest(
		"POST",
		"https://api.dropboxapi.com/oauth2/token",
		strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
	}

	req.SetBasicAuth(appKey, appSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var buf []byte
	buf, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(buf, &accessTokenResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return accessTokenResponse.AccessToken
}
