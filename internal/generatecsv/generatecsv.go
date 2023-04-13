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

func StatisticsCSV(p participants.Participants) {
	file, err := os.Create(path + "statistics-suicmc23.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	totalRegistered := 0
	totalIntendedPayments := 0
	totalTWINT := 0
	totalBankTransfer := 0
	totalCash := 0

	totalHousing := 0
	totalVolunteers := 0
	totalPreEvent := 0
	totalCargoRace := 0
	totalNabio := 0

	totalSmallTshirts := 0
	totalMediumTshirts := 0
	totalLargeTshirts := 0
	totalXlargeTshirts := 0

	for _, participant := range p.Items {

		totalRegistered += 1
		totalIntendedPayments += participant.IntendedPayment

		if participant.PreEvent {
			totalPreEvent += 1
		}

		if participant.CargoRace {
			totalCargoRace += 1
		}

		if participant.Nabio {
			totalNabio += 1
		}

		if participant.Housing == "true" {
			totalHousing += 1
		}

		if participant.Volunteering {
			totalVolunteers += 1
		}

		switch participant.PaymentMethod {
		case "TWINT":
			totalTWINT += 1

		case "Bank Transfer":
			totalBankTransfer += 1

		case "Cash":
			totalCash += 1
		}

		switch participant.TshirtSize {
		case "s":
			totalSmallTshirts += 1

		case "m":
			totalMediumTshirts += 1

		case "l":
			totalLargeTshirts += 1

		case "xl":
			totalXlargeTshirts += 1
		}
	}

	rows := [][]string{
		{"Category", "Totals"},
		{"Total Registered Participants", strconv.Itoa(totalRegistered)},
		{"", ""},
		{"FINANCE", ""},
		{"Total Intended Payments", strconv.Itoa(totalIntendedPayments)},
		{"Average Intended Payment", strconv.Itoa(totalIntendedPayments / totalRegistered)},
		{"Total Twint", strconv.Itoa(totalTWINT)},
		{"Total Bank Transfer", strconv.Itoa(totalBankTransfer)},
		{"Total Cash", strconv.Itoa(totalCash)},
		{"", ""},
		{"MISC", ""},
		{"Total Housing", strconv.Itoa(totalHousing)},
		{"Total Volunteers", strconv.Itoa(totalVolunteers)},
		{"Total PreEvent", strconv.Itoa(totalPreEvent)},
		{"Total Cargo Race", strconv.Itoa(totalCargoRace)},
		{"Total Safety", strconv.Itoa(totalNabio)},
		{"", ""},
		{"T-SHIRTS", ""},
		{"Small", strconv.Itoa(totalSmallTshirts)},
		{"Medium", strconv.Itoa(totalMediumTshirts)},
		{"Large", strconv.Itoa(totalLargeTshirts)},
		{"X-Large", strconv.Itoa(totalXlargeTshirts)},
		{"", ""},
	}

	var data [][]string
	for _, row := range rows {
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}
	}

	if err := w.WriteAll(data); err != nil {
		log.Fatalln("Couldn't write rows to file", err)
	}
}
