package authentication

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func AuthenticateBackend() (string, error) {
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

func RefreshDBXOAuthToken() string {
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
