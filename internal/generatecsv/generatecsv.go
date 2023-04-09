package generatecsv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"suicmc23/internal/participants"
	"suicmc23/internal/volunteers"
)

func FinanceCSV(p participants.Participants) {
	file, err := os.Create("suicmc23-data/finance-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{"first name", "nick name", "race number", "email", "intended payment", "payment method", "paid?"}
	err = w.Write(row)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range p.Items {
		row := []string{
			v.FirstName,
			v.NickName,
			strconv.Itoa(v.RaceNumber),
			v.Email,
			strconv.Itoa(v.IntendedPayment),
			v.PaymentMethod,
			strconv.FormatBool(v.Paid),
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}

func PreEventCSV(p participants.Participants) {
	file, err := os.Create("suicmc23-data/preevent-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{"first name", "nick name", "race number", "email"}
	err = w.Write(row)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range p.Items {
		if v.PreEvent {
			row := []string{
				v.FirstName,
				v.NickName,
				strconv.Itoa(v.RaceNumber),
				v.Email,
			}
			if err := w.Write(row); err != nil {
				log.Fatalln("Couldn't write row to file", err)
			}
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}

func VolunteersCSV(v volunteers.Volunteers) {
	file, err := os.Create("suicmc23-data/volunteers.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{
		"first_name",
		"last_name",
		"pronouns",
		"email",
		"telephone",
		"friday evening",
		"saturday morning",
		"saturday afternoon",
		"saturday evening",
		"sunday morning",
		"sunday afternoon",
		"sunday evening",
		"monday morning",
		"monday afternoon",
		"comments",
	}
	err = w.Write(row)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range v.Items {
		row := []string{
			v.FirstName,
			v.LastName,
			v.Pronouns,
			v.Email,
			v.Telephone,
			strconv.FormatBool(v.VolunteerFriday),
			strconv.FormatBool(v.VolunteerSaturdayMorning),
			strconv.FormatBool(v.VolunteerSaturdayAfternoon),
			strconv.FormatBool(v.VolunteerSaturdayEvening),
			strconv.FormatBool(v.VolunteerSundayMorning),
			strconv.FormatBool(v.VolunteerSundayAfternoon),
			strconv.FormatBool(v.VolunteerSundayEvening),
			strconv.FormatBool(v.VolunteerMondayMorning),
			strconv.FormatBool(v.VolunteerMondayAfternoon),
			v.AdditionalComments,
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}
