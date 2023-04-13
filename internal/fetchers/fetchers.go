package fetchers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"suicmc23/internal/participants"
	"suicmc23/internal/volunteers"
)

func GetParticipants(token string) participants.Participants {
	res := makeRequest(token, "/api/collections/participants/records?perPage=420")

	var participants participants.Participants
	err := json.Unmarshal(res, &participants)
	if err != nil {
		log.Fatal(err)
	}

	return participants
}

func GetVolunteers(token string) volunteers.Volunteers {
	res := makeRequest(token, "/api/collections/volunteers/records?perPage=420")

	var volunteers volunteers.Volunteers
	err := json.Unmarshal(res, &volunteers)
	if err != nil {
		log.Fatal(err)
	}

	return volunteers
}

func makeRequest(token, endpoint string) []byte {
	url := os.Getenv("BACKEND_URL")

	req, err := http.NewRequest("GET", url+endpoint, nil)
	req.Header.Add("Authorization", token)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return buf
}
