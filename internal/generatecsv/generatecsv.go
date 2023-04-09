package generatecsv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"suicmc23/internal/participants"
	"suicmc23/internal/volunteers"
)

var path string

func init() {
	if os.Getenv("AWS_EXECUTION_ENV") != "" {
		path = "/tmp/suicmc23-data/"
	} else {
		path = "suicmc23-data/"
	}
}

func ParticipantsCSV(p participants.Participants) {
	file, err := os.Create(path + "participants-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{
		"first_name",
		"last_name",
		"race_number",
		"nick_name",
		"pronouns",
		"email",
		"team",
		"city",
		"rank_selection",
		"tshirt_size",
		"cargo_race",
		"volunteering",
		"housing",
		"housing_friday",
		"housing_saturday",
		"housing_sunday",
		"pre_event",
		"payment_method",
		"intended_payment",
		"paid",
		"nabio",
		"additional_comments",
		"id",
		"created",
		"updated",
	}
	err = w.Write(row)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range p.Items {
		row := []string{
			v.FirstName,
			v.LastName,
			strconv.Itoa(v.RaceNumber),
			v.NickName,
			v.Pronouns,
			v.Email,
			v.Team,
			v.City,
			v.RankSelection,
			v.TshirtSize,
			strconv.FormatBool(v.CargoRace),
			strconv.FormatBool(v.Volunteering),
			v.Housing,
			strconv.FormatBool(v.HousingFriday),
			strconv.FormatBool(v.HousingSaturday),
			strconv.FormatBool(v.HousingSunday),
			strconv.FormatBool(v.PreEvent),
			v.PaymentMethod,
			strconv.Itoa(v.IntendedPayment),
			strconv.FormatBool(v.Paid),
			strconv.FormatBool(v.Nabio),
			v.AdditionalComments,
			v.ID,
			v.Created,
			v.Updated,
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}

func FinanceCSV(p participants.Participants) {
	file, err := os.Create(path + "finance-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{
		"first name",
		"nick name",
		"race number",
		"email",
		"intended payment",
		"payment method",
		"paid?",
	}
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
	file, err := os.Create(path + "preevent-suicmc23.csv")
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

func HousingCSV(p participants.Participants) {
	file, err := os.Create(path + "housing-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{
		"first name",
		"nick name",
		"race number",
		"email",
		"housing_friday",
		"housing_saturday",
		"housing_sunday",
	}
	err = w.Write(row)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range p.Items {
		if v.HousingFriday || v.HousingSaturday || v.HousingSunday {
			row := []string{
				v.FirstName,
				v.NickName,
				strconv.Itoa(v.RaceNumber),
				v.Email,
				strconv.FormatBool(v.HousingFriday),
				strconv.FormatBool(v.HousingSaturday),
				strconv.FormatBool(v.HousingSunday),
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
	file, err := os.Create(path + "volunteers-suicmc23.csv")
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
