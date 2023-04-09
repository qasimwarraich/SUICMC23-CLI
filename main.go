package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"

	"suicmc23/internal/authentication"
	"suicmc23/internal/dropbox"
	"suicmc23/internal/generatecsv"
	"suicmc23/internal/participants"
	"suicmc23/internal/printer"
	"suicmc23/internal/volunteers"
)

var path string

func main() {
	if os.Getenv("AWS_EXECUTION_ENV") != "" {
		path = "/tmp/suicmc23-data"
		lambda.Start(app)
	} else {
		path = "suicmc23-data"
		app()
	}
}

func app() (events.APIGatewayProxyResponse, error) {
	printer.Print("Welcome to SUICMC23 CLI", "theme")

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Couldn't load .env file", err)
	}

	printer.Print("Authenticating...", "guide")
	token, err := authentication.Authenticate()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print("Authenticated 🥳!", "guide")

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalln("Could not create data directory", err)
	}

	fmt.Println()
	printer.Print("Fetching Data", "guide")
	printer.Print("Getting participant list", "theme")
	participants := getParticipants(token)
	printer.Print("Getting volunteer list", "theme")
	volunteers := getVolunteers(token)

	fmt.Println()
	printer.Print("Generating CSV Files", "guide")
	printer.Print("Generating participants csv file", "theme")
	generatecsv.ParticipantsCSV(participants)

	printer.Print("Generating volunteers csv file", "theme")
	generatecsv.VolunteersCSV(volunteers)

	printer.Print("Generating finance csv file", "theme")
	generatecsv.FinanceCSV(participants)

	printer.Print("Generating pre-event csv file", "theme")
	generatecsv.PreEventCSV(participants)

	printer.Print("Generating housing csv file", "theme")
	generatecsv.HousingCSV(participants)

	fmt.Println()

	printer.Print("Uploading to Dropbox", "tip")
	dropbox.Upload()
	printer.Print("Done!", "guide")

	fmt.Println()
	printer.Print("Ciao and Chistole 👋", "theme")

	response := events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "text/html"},
		MultiValueHeaders: map[string][]string{},
		Body:              "Spamming of Data Successful 🤠, Have a nice day.",
		IsBase64Encoded:   false,
	}
	return response, nil
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
