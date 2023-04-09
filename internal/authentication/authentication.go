package authentication

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Authenticate() (string, error) {
	type AuthReq struct {
		Identity string
		Password string
	}

	type AuthRes struct {
		Admin map[string]interface{} `json:"admin"`
		Token string                 `json:"token"`
	}

	url := os.Getenv("BACKEND_URL")
	user := os.Getenv("IDENTITY")
	password := os.Getenv("PW")

	authBody := AuthReq{
		Identity: user,
		Password: password,
	}

	AuthBody, err := json.Marshal(authBody)
	requestBody := bytes.NewBuffer(AuthBody)
	if err != nil {
		return "", err
	}

	res, err := http.Post(url+"/api/admins/auth-with-password", "application/json", requestBody)
	if err != nil {
		return "", err
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resBody AuthRes
	err = json.Unmarshal(buf, &resBody)
	if err != nil {
		return "", err
	}

	return resBody.Token, nil
}
