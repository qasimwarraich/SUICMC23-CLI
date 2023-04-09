package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"suicmc23/internal/authentication"
	"suicmc23/internal/dropbox"
	"suicmc23/internal/generatecsv"
	"suicmc23/internal/participants"
	"suicmc23/internal/printer"
	"suicmc23/internal/volunteers"
)

func main() {
	printer.Print("Welcome to SUICMC23 CLI", "theme")

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Couldn't load .env file", err)
	}

	printer.Print("Authenticating", "guide")
	token, err := authentication.Authenticate()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print("Authenticated 🥳!", "guide")

	if err := os.MkdirAll("suicmc23-data", os.ModePerm); err != nil {
		log.Fatalln("Could not create data directory", err)
	}

	fmt.Println()
	printer.Print("Getting participant list", "theme")
	participants := getParticipants(token)
	printer.Print("Getting volunteer list", "theme")
	volunteers := getVolunteers(token)

	fmt.Println()
	printer.Print("Generating finance csv file", "theme")
	generatecsv.FinanceCSV(participants)

	printer.Print("Generating pre-event csv file", "theme")
	generatecsv.PreEventCSV(participants)

	printer.Print("Generating volunteer csv file", "theme")
	generatecsv.VolunteersCSV(volunteers)
	fmt.Println()

	printer.Print("Uploading to Dropbox", "tip")
	dropbox.Upload()

	printer.Print("Done!", "guide")
	printer.Print("Ciao and Chistole 👋", "theme")
}

func getParticipants(token string) participants.Participants {
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

func getVolunteers(token string) volunteers.Volunteers {
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
