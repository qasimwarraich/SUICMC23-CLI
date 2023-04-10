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
	url := os.Getenv("BACKEND_URL")

	req, err := http.NewRequest("GET", url+"/api/collections/participants/records?perPage=420", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var participants participants.Participants
	err = json.Unmarshal(buf, &participants)
	if err != nil {
		log.Fatal(err)
	}

	return participants
}

func GetVolunteers(token string) volunteers.Volunteers {
	url := os.Getenv("BACKEND_URL")

	req, err := http.NewRequest("GET", url+"/api/collections/volunteers/records?perPage=420", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var volunteers volunteers.Volunteers
	err = json.Unmarshal(buf, &volunteers)
	if err != nil {
		log.Fatal(err)
	}

	return volunteers
}
