package dropbox

import (
	"log"
	"os"
	"strings"
	"time"

	"suicmc23/internal/authentication"

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

	token := authentication.RefreshDBXOAuthToken()
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
